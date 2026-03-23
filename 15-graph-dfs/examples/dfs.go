package main

import "fmt"

// 그래프 DFS 기본 구현 - 재귀 DFS와 스택 DFS
// 시간 복잡도: O(V + E) (V: 정점 수, E: 간선 수)
// 공간 복잡도: O(V)

// 인접 리스트로 그래프를 표현한다
var adj [][]int
var visited []bool

// dfsRecursive 함수는 재귀를 사용한 DFS를 수행한다
func dfsRecursive(v int, order *[]int) {
	// 현재 정점을 방문 처리
	visited[v] = true
	*order = append(*order, v)

	// 인접 정점을 순회하며 미방문 정점을 재귀 탐색
	for _, next := range adj[v] {
		if !visited[next] {
			dfsRecursive(next, order)
		}
	}
}

// dfsStack 함수는 명시적 스택을 사용한 DFS를 수행한다
func dfsStack(start int) []int {
	order := []int{}
	stack := []int{start}

	for len(stack) > 0 {
		// 스택에서 정점을 꺼낸다
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 이미 방문한 정점이면 건너뛴다
		if visited[v] {
			continue
		}

		// 방문 처리
		visited[v] = true
		order = append(order, v)

		// 인접 정점을 역순으로 스택에 넣어 번호가 작은 정점부터 방문
		for i := len(adj[v]) - 1; i >= 0; i-- {
			next := adj[v][i]
			if !visited[next] {
				stack = append(stack, next)
			}
		}
	}
	return order
}

// countComponents 함수는 연결 요소의 개수를 구한다
func countComponents(n int) int {
	visited = make([]bool, n+1)
	count := 0

	for i := 1; i <= n; i++ {
		if !visited[i] {
			// 미방문 정점에서 DFS를 시작하면 새로운 연결 요소
			order := []int{}
			dfsRecursive(i, &order)
			count++
			fmt.Printf("  연결 요소 %d: %v\n", count, order)
		}
	}
	return count
}

func main() {
	// 6개 정점, 2개의 연결 요소로 구성된 그래프
	// 연결 요소 1: 1-2-3-4
	// 연결 요소 2: 5-6
	n := 6
	adj = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 추가 (양방향)
	edges := [][2]int{{1, 2}, {1, 3}, {2, 4}, {5, 6}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 재귀 DFS 실행
	fmt.Println("=== 재귀 DFS ===")
	visited = make([]bool, n+1)
	order := []int{}
	dfsRecursive(1, &order)
	fmt.Printf("1번 정점에서 시작: %v\n", order)

	// 스택 DFS 실행
	fmt.Println("\n=== 스택 DFS ===")
	visited = make([]bool, n+1)
	stackOrder := dfsStack(1)
	fmt.Printf("1번 정점에서 시작: %v\n", stackOrder)

	// 연결 요소 개수 구하기
	fmt.Println("\n=== 연결 요소 탐색 ===")
	count := countComponents(n)
	fmt.Printf("연결 요소의 개수: %d\n", count)
}
