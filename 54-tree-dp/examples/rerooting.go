package main

import "fmt"

// 트리 DP - 리루팅 (Rerooting)
// 모든 노드를 루트로 삼았을 때, 다른 모든 노드까지의 거리 합을 O(N)에 구한다.
// 시간 복잡도: O(N)
// 공간 복잡도: O(N)

var (
	adj [][]int // 인접 리스트
	sz  []int   // 서브트리 크기
	dp  []int   // 서브트리 내 거리 합 (Bottom-Up)
	ans []int   // 전체 트리 기준 거리 합 (리루팅 결과)
	n   int
)

// dfs1: Bottom-Up으로 서브트리 크기와 서브트리 내 거리 합을 계산한다
func dfs1(v, parent int) {
	sz[v] = 1
	dp[v] = 0
	for _, u := range adj[v] {
		if u == parent {
			continue
		}
		dfs1(u, v)
		sz[v] += sz[u]
		// 자식 u의 서브트리 내 거리 합 + u의 서브트리 노드 수 (각각 1만큼 더 멀어짐)
		dp[v] += dp[u] + sz[u]
	}
}

// dfs2: Top-Down 리루팅으로 모든 노드의 전체 거리 합을 계산한다
func dfs2(v, parent int) {
	for _, u := range adj[v] {
		if u == parent {
			continue
		}
		// 루트를 v에서 u로 옮길 때:
		// u의 서브트리 노드들은 1씩 가까워짐 → -sz[u]
		// 나머지 노드들은 1씩 멀어짐 → +(n - sz[u])
		ans[u] = ans[v] - sz[u] + (n - sz[u])
		dfs2(u, v)
	}
}

func main() {
	// 예시 트리:
	//   1 --- 2 --- 3
	//         |
	//         4
	n = 4
	adj = make([][]int, n+1)
	sz = make([]int, n+1)
	dp = make([]int, n+1)
	ans = make([]int, n+1)

	// 간선 추가 (양방향)
	edges := [][2]int{{1, 2}, {2, 3}, {2, 4}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 1단계: Bottom-Up DFS (루트 = 1)
	dfs1(1, 0)
	ans[1] = dp[1]

	// 2단계: Top-Down 리루팅
	dfs2(1, 0)

	// 결과 출력
	fmt.Println("각 노드에서 다른 모든 노드까지의 거리 합:")
	for i := 1; i <= n; i++ {
		fmt.Printf("  노드 %d: %d\n", i, ans[i])
	}
	// 출력:
	// 각 노드에서 다른 모든 노드까지의 거리 합:
	//   노드 1: 5
	//   노드 2: 3
	//   노드 3: 5
	//   노드 4: 5
}
