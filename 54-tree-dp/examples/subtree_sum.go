package main

import "fmt"

// 트리 DP - 서브트리 크기 구하기
// 각 노드를 루트로 하는 서브트리의 크기를 DFS 후위 순회로 계산한다.
// 시간 복잡도: O(N)
// 공간 복잡도: O(N)

var (
	adj [][]int // 인접 리스트
	sz  []int   // 서브트리 크기
)

// dfs는 후위 순회로 서브트리 크기를 계산한다
func dfs(v, parent int) {
	sz[v] = 1 // 자기 자신
	for _, u := range adj[v] {
		if u == parent {
			continue // 부모 방향으로 역행 방지
		}
		dfs(u, v)
		sz[v] += sz[u] // 자식 서브트리 크기 합산
	}
}

func main() {
	// 예시 트리:
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6
	n := 6
	adj = make([][]int, n+1)
	sz = make([]int, n+1)

	// 간선 추가 (양방향)
	edges := [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 루트를 1로 잡고 DFS 수행
	dfs(1, 0)

	// 결과 출력
	fmt.Println("서브트리 크기:")
	for i := 1; i <= n; i++ {
		fmt.Printf("  노드 %d: %d\n", i, sz[i])
	}
	// 출력:
	// 서브트리 크기:
	//   노드 1: 6
	//   노드 2: 3
	//   노드 3: 2
	//   노드 4: 1
	//   노드 5: 1
	//   노드 6: 1
}
