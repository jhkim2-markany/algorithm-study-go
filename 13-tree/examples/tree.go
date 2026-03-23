package main

import "fmt"

// 트리 기본 구현 - 트리 표현, 순회, 높이/깊이 계산
// 시간 복잡도: O(N) (N: 노드 수)
// 공간 복잡도: O(N)

// 인접 리스트로 트리를 표현한다
var adj [][]int

// 부모 배열: parent[v] = v의 부모 노드
var parent []int

// 깊이 배열: depth[v] = 루트에서 v까지의 거리
var depth []int

// buildTree 함수는 루트에서 DFS로 부모-자식 관계와 깊이를 구축한다
func buildTree(cur, par, d int) {
	parent[cur] = par
	depth[cur] = d
	for _, next := range adj[cur] {
		// 부모 노드로 되돌아가지 않도록 확인
		if next != par {
			buildTree(next, cur, d+1)
		}
	}
}

// treeHeight 함수는 특정 노드를 루트로 하는 서브트리의 높이를 반환한다
func treeHeight(cur, par int) int {
	// 리프 노드의 높이는 0
	maxH := 0
	for _, next := range adj[cur] {
		if next != par {
			// 자식 서브트리의 높이 중 최댓값을 구한다
			h := treeHeight(next, cur) + 1
			if h > maxH {
				maxH = h
			}
		}
	}
	return maxH
}

// preorder 함수는 전위 순회 결과를 반환한다 (현재 → 자식들)
func preorder(cur, par int, result *[]int) {
	// 현재 노드를 먼저 방문
	*result = append(*result, cur)
	for _, next := range adj[cur] {
		if next != par {
			preorder(next, cur, result)
		}
	}
}

// postorder 함수는 후위 순회 결과를 반환한다 (자식들 → 현재)
func postorder(cur, par int, result *[]int) {
	for _, next := range adj[cur] {
		if next != par {
			postorder(next, cur, result)
		}
	}
	// 자식을 모두 방문한 뒤 현재 노드를 방문
	*result = append(*result, cur)
}

// levelOrder 함수는 레벨 순회(BFS) 결과를 반환한다
func levelOrder(root int) []int {
	result := []int{}
	// 큐를 사용하여 레벨 순서대로 방문
	queue := []int{root}
	visited := make([]bool, len(adj))
	visited[root] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)

		for _, next := range adj[cur] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return result
}

func main() {
	// 7개 노드의 트리 구성
	// 구조:
	//        1
	//       / \
	//      2   3
	//     / \   \
	//    4   5   6
	//   /
	//  7
	n := 7
	adj = make([][]int, n+1)
	parent = make([]int, n+1)
	depth = make([]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 추가 (양방향)
	edges := [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {4, 7}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 루트를 1로 설정하고 트리 구축
	root := 1
	buildTree(root, 0, 0)

	fmt.Println("=== 트리 기본 정보 ===")
	for i := 1; i <= n; i++ {
		fmt.Printf("노드 %d: 부모=%d, 깊이=%d\n", i, parent[i], depth[i])
	}

	// 트리 높이 계산
	h := treeHeight(root, 0)
	fmt.Printf("\n트리의 높이: %d\n", h)

	// 전위 순회
	pre := []int{}
	preorder(root, 0, &pre)
	fmt.Printf("\n전위 순회: %v\n", pre)

	// 후위 순회
	post := []int{}
	postorder(root, 0, &post)
	fmt.Printf("후위 순회: %v\n", post)

	// 레벨 순회
	level := levelOrder(root)
	fmt.Printf("레벨 순회: %v\n", level)
}
