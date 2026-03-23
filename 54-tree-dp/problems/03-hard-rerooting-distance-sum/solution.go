package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		// 자식 서브트리의 거리 합 + 자식 서브트리 노드 수 (각각 간선 1개 추가)
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
		// 나머지 (n - sz[u])개 노드들은 1씩 멀어짐 → +(n - sz[u])
		ans[u] = ans[v] - sz[u] + (n - sz[u])
		dfs2(u, v)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수
	fmt.Fscan(reader, &n)

	adj = make([][]int, n+1)
	sz = make([]int, n+1)
	dp = make([]int, n+1)
	ans = make([]int, n+1)

	// 입력: 간선 정보
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 1단계: Bottom-Up DFS (루트 = 1)
	dfs1(1, 0)
	ans[1] = dp[1]

	// 2단계: Top-Down 리루팅
	dfs2(1, 0)

	// 출력: 각 노드의 거리 합
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, ans[i])
	}
	fmt.Fprintln(writer)
}
