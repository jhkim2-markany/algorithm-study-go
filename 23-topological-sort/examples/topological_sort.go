package main

import "fmt"

// 위상 정렬 기본 구현 - Kahn 알고리즘(BFS)과 DFS 기반
// 시간 복잡도: O(V + E)
// 공간 복잡도: O(V + E)

// kahnTopologicalSort 함수는 진입 차수 기반 BFS로 위상 정렬을 수행한다
func kahnTopologicalSort(n int, edges [][2]int) ([]int, bool) {
	// 인접 리스트와 진입 차수 배열 초기화
	adj := make([][]int, n)
	inDegree := make([]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}

	// 그래프 구성
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		inDegree[v]++
	}

	// 진입 차수가 0인 정점을 큐에 추가
	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		// 큐에서 정점을 꺼내 결과에 추가
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)

		// 인접 정점의 진입 차수를 감소
		for _, next := range adj[cur] {
			inDegree[next]--
			// 진입 차수가 0이 되면 큐에 추가
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 모든 정점이 결과에 포함되었는지 확인 (사이클 판별)
	if len(result) != n {
		return nil, false
	}
	return result, true
}

// dfsTopologicalSort 함수는 DFS 후위 순회 역순으로 위상 정렬을 수행한다
func dfsTopologicalSort(n int, edges [][2]int) ([]int, bool) {
	// 인접 리스트 초기화
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}

	// 방문 상태: 0=미방문, 1=방문 중, 2=방문 완료
	state := make([]int, n)
	stack := []int{}
	hasCycle := false

	// DFS 함수 정의
	var dfs func(u int)
	dfs = func(u int) {
		if hasCycle {
			return
		}
		state[u] = 1 // 방문 중 표시

		for _, v := range adj[u] {
			if state[v] == 1 {
				// 방문 중인 정점을 다시 만나면 사이클 존재
				hasCycle = true
				return
			}
			if state[v] == 0 {
				dfs(v)
			}
		}

		state[u] = 2 // 방문 완료 표시
		// 후위 순회: 모든 인접 정점 방문 후 스택에 추가
		stack = append(stack, u)
	}

	// 모든 정점에 대해 DFS 수행
	for i := 0; i < n; i++ {
		if state[i] == 0 {
			dfs(i)
		}
	}

	if hasCycle {
		return nil, false
	}

	// 스택을 뒤집으면 위상 정렬 결과
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = stack[n-1-i]
	}
	return result, true
}

func main() {
	// 예제 그래프: 6개 정점, 6개 간선
	// 0 → 1, 0 → 2, 1 → 3, 2 → 3, 3 → 4, 2 → 5
	n := 6
	edges := [][2]int{
		{0, 1}, {0, 2}, {1, 3}, {2, 3}, {3, 4}, {2, 5},
	}

	// Kahn 알고리즘 (BFS 기반) 위상 정렬
	fmt.Println("=== Kahn 알고리즘 (BFS 기반) ===")
	if order, ok := kahnTopologicalSort(n, edges); ok {
		fmt.Println("위상 정렬 결과:", order)
	} else {
		fmt.Println("사이클이 존재하여 위상 정렬 불가")
	}

	// DFS 기반 위상 정렬
	fmt.Println("\n=== DFS 기반 위상 정렬 ===")
	if order, ok := dfsTopologicalSort(n, edges); ok {
		fmt.Println("위상 정렬 결과:", order)
	} else {
		fmt.Println("사이클이 존재하여 위상 정렬 불가")
	}

	// 사이클이 있는 그래프 테스트
	fmt.Println("\n=== 사이클 판별 테스트 ===")
	cycleEdges := [][2]int{
		{0, 1}, {1, 2}, {2, 0},
	}
	if _, ok := kahnTopologicalSort(3, cycleEdges); !ok {
		fmt.Println("Kahn: 사이클 감지됨")
	}
	if _, ok := dfsTopologicalSort(3, cycleEdges); !ok {
		fmt.Println("DFS: 사이클 감지됨")
	}
}
