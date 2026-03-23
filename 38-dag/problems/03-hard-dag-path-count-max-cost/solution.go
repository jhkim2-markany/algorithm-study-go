package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MOD     = 1000000007
	NEG_INF = -int(1e18)
)

// 간선 구조체
type Edge struct {
	to, weight int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 정점 수, 간선 수, 시작 정점, 도착 정점
	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	// 인접 리스트 구성 및 진입 차수 계산
	graph := make([][]Edge, n)
	inDegree := make([]int, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		graph[u] = append(graph[u], Edge{v, w})
		inDegree[v]++
	}

	// 위상 정렬 (Kahn's Algorithm)
	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := []int{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, e := range graph[u] {
			inDegree[e.to]--
			if inDegree[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	// DP 배열 초기화
	// totalPaths: S에서 각 정점까지의 전체 경로 수
	totalPaths := make([]int, n)
	totalPaths[s] = 1

	// maxDist: S에서 각 정점까지의 최장 거리
	maxDist := make([]int, n)
	for i := range maxDist {
		maxDist[i] = NEG_INF
	}
	maxDist[s] = 0

	// maxPaths: 최장 경로를 달성하는 경로 수
	maxPaths := make([]int, n)
	maxPaths[s] = 1

	// 위상 정렬 순서대로 DP 수행
	for _, u := range order {
		for _, e := range graph[u] {
			v, w := e.to, e.weight

			// 전체 경로 수 갱신 (모듈러 연산)
			if totalPaths[u] > 0 {
				totalPaths[v] = (totalPaths[v] + totalPaths[u]) % MOD
			}

			// 최장 거리 및 최장 경로 수 갱신
			if maxDist[u] != NEG_INF {
				newDist := maxDist[u] + w
				if newDist > maxDist[v] {
					// 더 긴 경로 발견: 거리와 경로 수 모두 갱신
					maxDist[v] = newDist
					maxPaths[v] = maxPaths[u]
				} else if newDist == maxDist[v] {
					// 같은 길이의 경로 발견: 경로 수만 추가
					maxPaths[v] = (maxPaths[v] + maxPaths[u]) % MOD
				}
			}
		}
	}

	// 결과 출력
	if totalPaths[t] == 0 {
		// 도달 불가능한 경우
		fmt.Fprintln(writer, 0)
		fmt.Fprintln(writer, 0)
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, totalPaths[t])
		fmt.Fprintln(writer, maxDist[t])
		fmt.Fprintln(writer, maxPaths[t])
	}
}
