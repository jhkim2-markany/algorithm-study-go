package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxN = 100001
const logN = 17

// 트리 인접 리스트
var tree [maxN][]int

// Binary Lifting 배열
var up [maxN][logN]int
var dep [maxN]int

// DFS로 깊이와 희소 배열을 구성한다
func dfs(v, par, d int) {
	dep[v] = d
	up[v][0] = par
	for k := 1; k < logN; k++ {
		up[v][k] = up[up[v][k-1]][k-1]
	}
	for _, u := range tree[v] {
		if u != par {
			dfs(u, v, d+1)
		}
	}
}

// Binary Lifting으로 LCA를 구한다
func lca(u, v int) int {
	// u가 더 깊도록 보장한다
	if dep[u] < dep[v] {
		u, v = v, u
	}
	// 깊이 차이만큼 u를 올린다
	diff := dep[u] - dep[v]
	for k := 0; k < logN; k++ {
		if (diff>>k)&1 == 1 {
			u = up[u][k]
		}
	}
	// 같으면 LCA이다
	if u == v {
		return u
	}
	// 큰 점프부터 시도하며 LCA 바로 아래까지 올린다
	for k := logN - 1; k >= 0; k-- {
		if up[u][k] != up[v][k] {
			u = up[u][k]
			v = up[v][k]
		}
	}
	return up[u][0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 간선 입력
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}

	// 전처리: DFS로 깊이와 희소 배열 구성
	dfs(1, 0, 0)

	// 쿼리 처리
	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		// LCA를 구하여 출력한다
		fmt.Fprintln(writer, lca(u, v))
	}
}
