package main

import "fmt"

// DAG DP - 위상 정렬 기반 최단/최장 경로 및 경로 수 계산
// 시간 복잡도: O(V + E)
// 공간 복잡도: O(V + E)

const INF = int(1e18)

// 간선 구조체
type Edge struct {
	to, weight int
}

// 위상 정렬 (Kahn's Algorithm, BFS 기반)
func topologicalSort(n int, graph [][]Edge) []int {
	// 진입 차수 계산
	inDegree := make([]int, n)
	for u := 0; u < n; u++ {
		for _, e := range graph[u] {
			inDegree[e.to]++
		}
	}

	// 진입 차수가 0인 정점을 큐에 추가
	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS로 위상 정렬 수행
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

	return order
}

// DAG 최단 경로: 시작 정점에서 모든 정점까지의 최단 거리
func dagShortestPath(n int, graph [][]Edge, start int) []int {
	order := topologicalSort(n, graph)

	// 거리 배열 초기화
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0

	// 위상 정렬 순서대로 간선 완화
	for _, u := range order {
		if dist[u] == INF {
			continue
		}
		for _, e := range graph[u] {
			// 최단 거리 갱신
			if dist[u]+e.weight < dist[e.to] {
				dist[e.to] = dist[u] + e.weight
			}
		}
	}

	return dist
}

// DAG 최장 경로: 시작 정점에서 모든 정점까지의 최장 거리
func dagLongestPath(n int, graph [][]Edge, start int) []int {
	order := topologicalSort(n, graph)

	// 거리 배열 초기화 (음의 무한대)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -INF
	}
	dist[start] = 0

	// 위상 정렬 순서대로 간선 완화
	for _, u := range order {
		if dist[u] == -INF {
			continue
		}
		for _, e := range graph[u] {
			// 최장 거리 갱신
			if dist[u]+e.weight > dist[e.to] {
				dist[e.to] = dist[u] + e.weight
			}
		}
	}

	return dist
}

// DAG 경로 수: 시작 정점에서 각 정점까지의 경로 수
func dagCountPaths(n int, graph [][]Edge, start int) []int {
	order := topologicalSort(n, graph)

	// 경로 수 배열 초기화
	count := make([]int, n)
	count[start] = 1

	// 위상 정렬 순서대로 경로 수 전파
	for _, u := range order {
		if count[u] == 0 {
			continue
		}
		for _, e := range graph[u] {
			count[e.to] += count[u]
		}
	}

	return count
}

func main() {
	// 예제 DAG: 6개 정점, 7개 간선
	// 0 → 1 (가중치 5)
	// 0 → 2 (가중치 3)
	// 1 → 3 (가중치 6)
	// 1 → 2 (가중치 2)
	// 2 → 4 (가중치 4)
	// 2 → 3 (가중치 7)
	// 3 → 4 (가중치 1)
	n := 5
	graph := make([][]Edge, n)
	for i := range graph {
		graph[i] = []Edge{}
	}

	// 간선 추가
	graph[0] = append(graph[0], Edge{1, 5})
	graph[0] = append(graph[0], Edge{2, 3})
	graph[1] = append(graph[1], Edge{3, 6})
	graph[1] = append(graph[1], Edge{2, 2})
	graph[2] = append(graph[2], Edge{4, 4})
	graph[2] = append(graph[2], Edge{3, 7})
	graph[3] = append(graph[3], Edge{4, 1})

	start := 0

	// 위상 정렬 결과 출력
	order := topologicalSort(n, graph)
	fmt.Println("위상 정렬 순서:", order)

	// 최단 경로 계산
	shortest := dagShortestPath(n, graph, start)
	fmt.Println("\n최단 경로 (정점 0에서 출발):")
	for i := 0; i < n; i++ {
		if shortest[i] == INF {
			fmt.Printf("  정점 %d: 도달 불가\n", i)
		} else {
			fmt.Printf("  정점 %d: %d\n", i, shortest[i])
		}
	}

	// 최장 경로 계산
	longest := dagLongestPath(n, graph, start)
	fmt.Println("\n최장 경로 (정점 0에서 출발):")
	for i := 0; i < n; i++ {
		if longest[i] == -INF {
			fmt.Printf("  정점 %d: 도달 불가\n", i)
		} else {
			fmt.Printf("  정점 %d: %d\n", i, longest[i])
		}
	}

	// 경로 수 계산
	paths := dagCountPaths(n, graph, start)
	fmt.Println("\n경로 수 (정점 0에서 출발):")
	for i := 0; i < n; i++ {
		fmt.Printf("  정점 %d: %d개\n", i, paths[i])
	}
}
