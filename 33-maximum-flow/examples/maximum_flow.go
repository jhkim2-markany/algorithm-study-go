package main

import "fmt"

// 최대 유량 (Maximum Flow) - Edmonds-Karp 알고리즘
// BFS를 이용하여 소스에서 싱크로의 최대 유량을 구한다.
// 시간 복잡도: O(V × E²)
// 공간 복잡도: O(V + E)

const INF = 1<<60 - 1

// 간선 구조체: 도착 정점, 용량, 현재 유량, 역방향 간선 인덱스
type Edge struct {
	to, cap, flow, rev int
}

// 그래프: 인접 리스트로 표현
var graph [][]Edge

// initGraph: 정점 수 n으로 그래프를 초기화한다
func initGraph(n int) {
	graph = make([][]Edge, n)
	for i := range graph {
		graph[i] = []Edge{}
	}
}

// addEdge: 용량 cap인 간선 u→v를 추가한다 (역방향 간선도 함께 추가)
func addEdge(u, v, cap int) {
	// 순방향 간선과 역방향 간선을 서로 참조하도록 인덱스를 저장한다
	graph[u] = append(graph[u], Edge{to: v, cap: cap, flow: 0, rev: len(graph[v])})
	graph[v] = append(graph[v], Edge{to: u, cap: 0, flow: 0, rev: len(graph[u]) - 1})
}

// edmondsKarp: 소스 s에서 싱크 t로의 최대 유량을 구한다
func edmondsKarp(s, t int) int {
	totalFlow := 0
	n := len(graph)

	for {
		// BFS로 소스에서 싱크까지의 증가 경로를 찾는다
		parent := make([]int, n)
		parentEdge := make([]int, n)
		for i := range parent {
			parent[i] = -1
		}
		parent[s] = s

		queue := []int{s}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for idx, e := range graph[cur] {
				// 잔여 용량이 있고 아직 방문하지 않은 정점이면 탐색한다
				if e.cap-e.flow > 0 && parent[e.to] == -1 {
					parent[e.to] = cur
					parentEdge[e.to] = idx
					queue = append(queue, e.to)
				}
			}
		}

		// 싱크에 도달하지 못하면 종료한다
		if parent[t] == -1 {
			break
		}

		// 증가 경로의 병목 용량을 구한다
		bottleneck := INF
		v := t
		for v != s {
			u := parent[v]
			e := graph[u][parentEdge[v]]
			residual := e.cap - e.flow
			if residual < bottleneck {
				bottleneck = residual
			}
			v = u
		}

		// 경로를 따라 유량을 갱신한다
		v = t
		for v != s {
			u := parent[v]
			idx := parentEdge[v]
			graph[u][idx].flow += bottleneck               // 순방향 유량 증가
			graph[v][graph[u][idx].rev].flow -= bottleneck // 역방향 유량 감소
			v = u
		}

		totalFlow += bottleneck
	}

	return totalFlow
}

func main() {
	// 예시: 6개 정점의 네트워크 유량 문제
	// 소스: 0, 싱크: 5
	n := 6
	initGraph(n)

	// 간선 추가 (시작, 끝, 용량)
	addEdge(0, 1, 16)
	addEdge(0, 2, 13)
	addEdge(1, 2, 4)
	addEdge(1, 3, 12)
	addEdge(2, 1, 10)
	addEdge(2, 4, 14)
	addEdge(3, 2, 9)
	addEdge(3, 5, 20)
	addEdge(4, 3, 7)
	addEdge(4, 5, 4)

	source := 0
	sink := 5

	// Edmonds-Karp 알고리즘으로 최대 유량을 구한다
	maxFlow := edmondsKarp(source, sink)

	fmt.Printf("정점 수: %d\n", n)
	fmt.Printf("소스: %d, 싱크: %d\n", source, sink)
	fmt.Printf("최대 유량: %d\n", maxFlow)

	// 각 간선의 유량 상태를 출력한다
	fmt.Println("\n간선별 유량 상태:")
	for u := 0; u < n; u++ {
		for _, e := range graph[u] {
			// 순방향 간선만 출력한다 (용량이 0보다 큰 간선)
			if e.cap > 0 {
				fmt.Printf("  %d → %d: 유량 %d / 용량 %d\n", u, e.to, e.flow, e.cap)
			}
		}
	}
}
