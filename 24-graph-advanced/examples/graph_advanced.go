package main

import "fmt"

// 고급 그래프 알고리즘 - 이분 그래프 판별 및 강한 연결 요소(SCC)
// 이분 그래프 판별: O(V + E)
// Kosaraju SCC: O(V + E)

// isBipartite 함수는 인접 리스트로 표현된 그래프가 이분 그래프인지 판별한다
func isBipartite(n int, adj [][]int) bool {
	color := make([]int, n)
	for i := range color {
		color[i] = -1 // 미색칠 상태
	}

	// BFS로 각 연결 요소를 색칠한다
	for start := 0; start < n; start++ {
		if color[start] != -1 {
			continue
		}
		// 시작 정점에 색 0을 칠한다
		color[start] = 0
		queue := []int{start}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for _, next := range adj[cur] {
				if color[next] == -1 {
					// 인접 정점에 반대 색을 칠한다
					color[next] = 1 - color[cur]
					queue = append(queue, next)
				} else if color[next] == color[cur] {
					// 같은 색이면 이분 그래프가 아니다
					return false
				}
			}
		}
	}
	return true
}

// kosarajuSCC 함수는 Kosaraju 알고리즘으로 강한 연결 요소를 구한다
func kosarajuSCC(n int, adj [][]int) [][]int {
	// 1단계: 원본 그래프에서 DFS 수행, 완료 순서를 스택에 기록
	visited := make([]bool, n)
	order := []int{}

	var dfs1 func(u int)
	dfs1 = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs1(v)
			}
		}
		order = append(order, u) // 완료 시점에 스택에 추가
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs1(i)
		}
	}

	// 2단계: 역방향 그래프 생성
	radj := make([][]int, n)
	for i := 0; i < n; i++ {
		radj[i] = []int{}
	}
	for u := 0; u < n; u++ {
		for _, v := range adj[u] {
			radj[v] = append(radj[v], u)
		}
	}

	// 3단계: 스택 역순으로 역방향 그래프에서 DFS 수행
	for i := range visited {
		visited[i] = false
	}
	sccs := [][]int{}
	var component []int

	var dfs2 func(u int)
	dfs2 = func(u int) {
		visited[u] = true
		component = append(component, u)
		for _, v := range radj[u] {
			if !visited[v] {
				dfs2(v)
			}
		}
	}

	// 완료 순서의 역순으로 처리
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		if !visited[u] {
			component = []int{}
			dfs2(u)
			sccs = append(sccs, component)
		}
	}

	return sccs
}

func main() {
	// === 이분 그래프 판별 예제 ===
	fmt.Println("=== 이분 그래프 판별 ===")

	// 이분 그래프 예시: 0-1, 0-3, 1-2, 2-3 (4개 정점, 사이클 길이 4)
	adj1 := [][]int{
		{1, 3}, // 정점 0의 인접 정점
		{0, 2}, // 정점 1의 인접 정점
		{1, 3}, // 정점 2의 인접 정점
		{0, 2}, // 정점 3의 인접 정점
	}
	fmt.Printf("그래프 1 (0-1-2-3-0): 이분 그래프 = %v\n", isBipartite(4, adj1))

	// 이분 그래프가 아닌 예시: 0-1, 1-2, 0-2 (삼각형)
	adj2 := [][]int{
		{1, 2}, // 정점 0의 인접 정점
		{0, 2}, // 정점 1의 인접 정점
		{0, 1}, // 정점 2의 인접 정점
	}
	fmt.Printf("그래프 2 (삼각형 0-1-2): 이분 그래프 = %v\n", isBipartite(3, adj2))

	// === 강한 연결 요소 (Kosaraju) 예제 ===
	fmt.Println("\n=== 강한 연결 요소 (Kosaraju) ===")

	// 방향 그래프: 0→1, 1→2, 2→0, 1→3, 3→4, 4→5, 5→3
	// SCC: {0,1,2}, {3,4,5}
	adj3 := [][]int{
		{1},    // 정점 0 → 1
		{2, 3}, // 정점 1 → 2, 3
		{0},    // 정점 2 → 0
		{4},    // 정점 3 → 4
		{5},    // 정점 4 → 5
		{3},    // 정점 5 → 3
	}
	sccs := kosarajuSCC(6, adj3)
	fmt.Printf("SCC 개수: %d\n", len(sccs))
	for i, scc := range sccs {
		fmt.Printf("SCC %d: %v\n", i+1, scc)
	}
}
