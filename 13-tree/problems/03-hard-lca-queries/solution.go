package main

import (
	"bufio"
	"fmt"
	"os"
)

// LCA (최소 공통 조상) - 희소 테이블(Sparse Table)을 이용한 풀이
// 시간 복잡도: 전처리 O(N log N), 쿼리 O(log N)
// 공간 복잡도: O(N log N)

const MAXLOG = 17 // log2(100000) ≈ 16.6

var (
	adj   [][]int
	depth []int
	up    [MAXLOG][]int // up[k][v] = v의 2^k번째 조상
)

// dfs 함수는 루트에서 시작하여 깊이와 부모 정보를 구축한다
func dfs(cur, par, d int) {
	depth[cur] = d
	up[0][cur] = par
	for _, next := range adj[cur] {
		if next != par {
			dfs(next, cur, d+1)
		}
	}
}

// lca 함수는 두 노드의 최소 공통 조상을 반환한다
func lca(u, v int) int {
	// 깊이가 더 깊은 노드를 u로 설정
	if depth[u] < depth[v] {
		u, v = v, u
	}

	// u의 깊이를 v와 맞춘다 (이진 리프팅)
	diff := depth[u] - depth[v]
	for k := 0; k < MAXLOG; k++ {
		if (diff>>k)&1 == 1 {
			u = up[k][u]
		}
	}

	// 같은 노드이면 바로 반환
	if u == v {
		return u
	}

	// 두 노드를 동시에 올려서 LCA 직전까지 이동
	for k := MAXLOG - 1; k >= 0; k-- {
		if up[k][u] != up[k][v] {
			u = up[k][u]
			v = up[k][v]
		}
	}

	// 한 칸 더 올리면 LCA
	return up[0][u]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 인접 리스트 초기화
	adj = make([][]int, n+1)
	depth = make([]int, n+1)
	for k := 0; k < MAXLOG; k++ {
		up[k] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// DFS로 깊이와 직접 부모 정보 구축
	dfs(1, 0, 0)

	// 희소 테이블 구축: up[k][v] = v의 2^k번째 조상
	for k := 1; k < MAXLOG; k++ {
		for v := 1; v <= n; v++ {
			up[k][v] = up[k-1][up[k-1][v]]
		}
	}

	// 쿼리 처리
	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintln(writer, lca(a, b))
	}
}
