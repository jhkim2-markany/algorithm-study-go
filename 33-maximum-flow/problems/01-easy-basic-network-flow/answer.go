package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// Edge는 유량 네트워크의 간선을 나타낸다.
type Edge struct {
	to, cap, flow, rev int
}

var graph [][]Edge

// addEdge는 용량 cap인 간선 u→v를 추가한다.
func addEdge(u, v, cap int) {
	graph[u] = append(graph[u], Edge{to: v, cap: cap, flow: 0, rev: len(graph[v])})
	graph[v] = append(graph[v], Edge{to: u, cap: 0, flow: 0, rev: len(graph[u]) - 1})
}

// maxFlow는 소스에서 싱크까지의 최대 유량을 구한다.
//
// [매개변수]
//   - n: 정점의 수 (1-indexed)
//   - edges: 간선 목록 (각 원소는 [u, v, cap])
//   - source: 소스 정점 번호
//   - sink: 싱크 정점 번호
//
// [반환값]
//   - int: 소스에서 싱크까지의 최대 유량
//
// [알고리즘 힌트]
//
//	에드몬드-카프 알고리즘 (BFS 기반 Ford-Fulkerson)을 사용한다.
//	BFS로 증가 경로를 찾고 병목 용량만큼 유량을 갱신한다.
//	시간복잡도: O(V * E^2)
func maxFlow(n int, edges [][3]int, source, sink int) int {
	graph = make([][]Edge, n+1)
	for i := range graph {
		graph[i] = []Edge{}
	}

	for _, e := range edges {
		addEdge(e[0], e[1], e[2])
	}

	totalFlow := 0
	total := n + 1

	for {
		parent := make([]int, total)
		parentEdge := make([]int, total)
		for i := range parent {
			parent[i] = -1
		}
		parent[source] = source

		queue := []int{source}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for idx, e := range graph[cur] {
				if e.cap-e.flow > 0 && parent[e.to] == -1 {
					parent[e.to] = cur
					parentEdge[e.to] = idx
					queue = append(queue, e.to)
				}
			}
		}

		if parent[sink] == -1 {
			break
		}

		bottleneck := INF
		v := sink
		for v != source {
			u := parent[v]
			e := graph[u][parentEdge[v]]
			if e.cap-e.flow < bottleneck {
				bottleneck = e.cap - e.flow
			}
			v = u
		}

		v = sink
		for v != source {
			u := parent[v]
			idx := parentEdge[v]
			graph[u][idx].flow += bottleneck
			graph[v][graph[u][idx].rev].flow -= bottleneck
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

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	fmt.Fprintln(writer, maxFlow(n, edges, 1, n))
}
