package main

import "fmt"

// 이분 매칭 (Bipartite Matching) - 헝가리안 알고리즘 (Kuhn's Algorithm)
// DFS를 이용하여 이분 그래프에서 최대 매칭을 구한다.
// 시간 복잡도: O(V × E)
// 공간 복잡도: O(V + E)

// 왼쪽 정점 수, 오른쪽 정점 수
var n, m int

// adj[i]: 왼쪽 정점 i와 연결된 오른쪽 정점 목록
var adj [][]int

// matchL[i]: 왼쪽 정점 i가 매칭된 오른쪽 정점 (-1이면 미매칭)
// matchR[j]: 오른쪽 정점 j가 매칭된 왼쪽 정점 (-1이면 미매칭)
var matchL, matchR []int

// visited: DFS 중 방문 여부
var visited []bool

// dfs: 왼쪽 정점 u에서 시작하여 증가 경로를 찾는다
func dfs(u int) bool {
	for _, v := range adj[u] {
		if visited[v] {
			continue
		}
		visited[v] = true

		// v가 미매칭이거나, v의 현재 매칭 상대에서 다른 경로를 찾을 수 있으면 매칭한다
		if matchR[v] == -1 || dfs(matchR[v]) {
			matchL[u] = v
			matchR[v] = u
			return true
		}
	}
	return false
}

// bipartiteMatching: 최대 매칭 수를 반환한다
func bipartiteMatching() int {
	matchL = make([]int, n)
	matchR = make([]int, m)
	for i := range matchL {
		matchL[i] = -1
	}
	for i := range matchR {
		matchR[i] = -1
	}

	result := 0
	// 왼쪽 그룹의 각 정점에 대해 증가 경로를 찾는다
	for u := 0; u < n; u++ {
		visited = make([]bool, m)
		if dfs(u) {
			result++
		}
	}
	return result
}

func main() {
	// 예시: 왼쪽 4명의 직원, 오른쪽 4개의 업무
	// 각 직원이 수행 가능한 업무가 주어질 때 최대 매칭을 구한다
	n = 4 // 왼쪽 정점 수 (직원)
	m = 4 // 오른쪽 정점 수 (업무)

	adj = make([][]int, n)
	// 직원 0: 업무 0, 1 가능
	adj[0] = []int{0, 1}
	// 직원 1: 업무 0 가능
	adj[1] = []int{0}
	// 직원 2: 업무 1, 2 가능
	adj[2] = []int{1, 2}
	// 직원 3: 업무 2, 3 가능
	adj[3] = []int{2, 3}

	// 최대 매칭을 구한다
	maxMatch := bipartiteMatching()

	fmt.Printf("왼쪽 정점 수: %d, 오른쪽 정점 수: %d\n", n, m)
	fmt.Printf("최대 매칭 수: %d\n", maxMatch)

	// 매칭 결과를 출력한다
	fmt.Println("\n매칭 결과:")
	for i := 0; i < n; i++ {
		if matchL[i] != -1 {
			fmt.Printf("  직원 %d → 업무 %d\n", i, matchL[i])
		} else {
			fmt.Printf("  직원 %d → 미배정\n", i)
		}
	}
}
