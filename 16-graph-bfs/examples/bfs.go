package main

import "fmt"

// 그래프 BFS 기본 구현 - 큐를 이용한 너비 우선 탐색과 최단 거리 계산
// 시간 복잡도: O(V + E) (V: 정점 수, E: 간선 수)
// 공간 복잡도: O(V)

// bfs 함수는 시작 정점에서 BFS를 수행하고 방문 순서를 반환한다
func bfs(adj [][]int, start int, n int) []int {
	visited := make([]bool, n+1)
	queue := []int{start}
	visited[start] = true
	order := []int{}

	for len(queue) > 0 {
		// 큐의 앞에서 정점을 꺼낸다
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)

		// 인접 정점을 순회하며 미방문 정점을 큐에 넣는다
		for _, next := range adj[v] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return order
}

// bfsDistance 함수는 시작 정점에서 모든 정점까지의 최단 거리를 구한다
func bfsDistance(adj [][]int, start int, n int) []int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = -1 // -1은 도달 불가능을 의미
	}
	dist[start] = 0
	queue := []int{start}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, next := range adj[v] {
			if dist[next] == -1 {
				// 최단 거리는 현재 정점의 거리 + 1
				dist[next] = dist[v] + 1
				queue = append(queue, next)
			}
		}
	}
	return dist
}

func main() {
	// 7개 정점으로 구성된 그래프
	//   1 - 2 - 4
	//   |   |
	//   3 - 5 - 6
	//           |
	//           7
	n := 7
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 추가 (양방향)
	edges := [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 5}, {5, 6}, {6, 7}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// BFS 방문 순서
	fmt.Println("=== BFS 방문 순서 ===")
	order := bfs(adj, 1, n)
	fmt.Printf("1번 정점에서 시작: %v\n", order)

	// 최단 거리 계산
	fmt.Println("\n=== 최단 거리 (1번 정점 기준) ===")
	dist := bfsDistance(adj, 1, n)
	for i := 1; i <= n; i++ {
		fmt.Printf("  1 → %d: 거리 %d\n", i, dist[i])
	}
}
