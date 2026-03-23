package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1

// maxFlow는 Edmonds-Karp 알고리즘으로 소스에서 싱크까지의 최대 유량을 구한다.
//
// [매개변수]
//   - capacity: 잔여 용량 행렬 (n+1 × n+1)
//   - adjList: 인접 리스트 (역방향 포함)
//   - n: 정점의 수
//   - s: 소스 정점
//   - t: 싱크 정점
//
// [반환값]
//   - int: 최대 유량
//
// [알고리즘 힌트]
//
//	Edmonds-Karp 알고리즘 (BFS 기반 Ford-Fulkerson)을 사용한다.
//	BFS로 소스에서 싱크까지의 증가 경로를 찾고,
//	경로상 최소 잔여 용량(bottleneck)만큼 유량을 흘린다.
//	순방향 용량을 감소시키고 역방향 용량을 증가시킨다.
//	증가 경로가 없을 때까지 반복한다.
func maxFlow(capacity [][]int, adjList [][]int, n, s, t int) int {
	// BFS로 증가 경로를 찾는 내부 함수
	bfs := func() (int, []int) {
		parent := make([]int, n+1)
		for i := range parent {
			parent[i] = -1
		}
		parent[s] = s

		queue := []int{s}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for _, next := range adjList[cur] {
				if parent[next] == -1 && capacity[cur][next] > 0 {
					parent[next] = cur
					if next == t {
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

	totalFlow := 0
	for {
		bottleneck, parent := bfs()
		if bottleneck == 0 {
			break
		}

		v := t
		for v != s {
			u := parent[v]
			capacity[u][v] -= bottleneck
			capacity[v][u] += bottleneck
			v = u
		}

		totalFlow += bottleneck
	}

	return totalFlow
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	// 잔여 용량 행렬 초기화
	capacity := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		capacity[i] = make([]int, n+1)
	}

	// 인접 리스트 초기화
	adjList := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adjList[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		capacity[u][v] += c
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	// 핵심 함수 호출
	result := maxFlow(capacity, adjList, n, s, t)

	fmt.Fprintln(writer, result)
}
