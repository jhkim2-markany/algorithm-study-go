package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edmonds-Karp 알고리즘 (BFS 기반 Ford-Fulkerson)으로 최대 유량을 구한다

const INF = 1<<31 - 1

var (
	n, m, s, t int
	capacity   [][]int // 잔여 용량 행렬
	adjList    [][]int // 인접 리스트 (역방향 포함)
)

// bfs 함수는 소스에서 싱크까지의 증가 경로를 찾고 경로상 최소 잔여 용량을 반환한다
func bfs() (int, []int) {
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = -1
	}
	parent[s] = s

	queue := []int{s}

	// BFS로 증가 경로 탐색
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, next := range adjList[cur] {
			// 잔여 용량이 있고 미방문인 정점으로 이동
			if parent[next] == -1 && capacity[cur][next] > 0 {
				parent[next] = cur
				if next == t {
					// 싱크에 도달하면 경로상 최소 잔여 용량(bottleneck) 계산
					bottleneck := INF
					v := t
					for v != s {
						u := parent[v]
						if capacity[u][v] < bottleneck {
							bottleneck = capacity[u][v]
						}
						v = u
					}
					return bottleneck, parent
				}
				queue = append(queue, next)
			}
		}
	}

	return 0, parent
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &m, &s, &t)

	// 잔여 용량 행렬 초기화
	capacity = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		capacity[i] = make([]int, n+1)
	}

	// 인접 리스트 초기화
	adjList = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adjList[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		capacity[u][v] += c // 중복 간선 처리: 용량을 합산

		// 인접 리스트에 양방향 추가 (역방향 간선도 필요)
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	// Edmonds-Karp: BFS로 증가 경로를 반복적으로 찾아 유량을 늘린다
	maxFlow := 0
	for {
		bottleneck, parent := bfs()
		if bottleneck == 0 {
			break // 증가 경로가 없으면 종료
		}

		// 경로상 잔여 용량 갱신
		v := t
		for v != s {
			u := parent[v]
			capacity[u][v] -= bottleneck // 순방향 용량 감소
			capacity[v][u] += bottleneck // 역방향 용량 증가
			v = u
		}

		maxFlow += bottleneck
	}

	fmt.Fprintln(writer, maxFlow)
}
