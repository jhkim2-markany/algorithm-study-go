package main

import (
	"bufio"
	"fmt"
	"os"
)

// HLD를 이용한 경로 합 질의
// 시간 복잡도: 전처리 O(N), 질의 O(Q log²N)

var (
	adj    [][]int
	sz     []int
	dep    []int
	par    []int
	heavy  []int
	top    []int
	pos    []int
	seg    []int
	val    []int
	curPos int
	n, q   int
)

// dfs1: 서브트리 크기, 깊이, 부모, Heavy Child 계산
func dfs1(v, p, d int) {
	par[v] = p
	dep[v] = d
	sz[v] = 1
	heavy[v] = -1
	maxSz := 0
	for _, u := range adj[v] {
		if u == p {
			continue
		}
		dfs1(u, v, d+1)
		sz[v] += sz[u]
		if sz[u] > maxSz {
			maxSz = sz[u]
			heavy[v] = u
		}
	}
}

// dfs2: DFS 번호 부여, 체인 구성
func dfs2(v, chainTop int) {
	top[v] = chainTop
	pos[v] = curPos
	curPos++
	if heavy[v] != -1 {
		dfs2(heavy[v], chainTop) // Heavy Child 우선 방문
	}
	for _, u := range adj[v] {
		if u == par[v] || u == heavy[v] {
			continue
		}
		dfs2(u, u) // Light Child → 새 체인
	}
}

// 세그먼트 트리 구축
func build(node, s, e int, arr []int) {
	if s == e {
		seg[node] = arr[s]
		return
	}
	mid := (s + e) / 2
	build(node*2, s, mid, arr)
	build(node*2+1, mid+1, e, arr)
	seg[node] = seg[node*2] + seg[node*2+1]
}

// 구간 합 질의
func query(node, s, e, l, r int) int {
	if r < s || e < l {
		return 0
	}
	if l <= s && e <= r {
		return seg[node]
	}
	mid := (s + e) / 2
	return query(node*2, s, mid, l, r) + query(node*2+1, mid+1, e, l, r)
}

// 경로 합 질의: 노드 u에서 v까지의 경로 합
func pathQuery(u, v int) int {
	res := 0
	for top[u] != top[v] {
		if dep[top[u]] < dep[top[v]] {
			u, v = v, u
		}
		res += query(1, 0, n-1, pos[top[u]], pos[u])
		u = par[top[u]]
	}
	if dep[u] > dep[v] {
		u, v = v, u
	}
	res += query(1, 0, n-1, pos[u], pos[v])
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &q)

	val = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &val[i])
	}

	adj = make([][]int, n+1)
	sz = make([]int, n+1)
	dep = make([]int, n+1)
	par = make([]int, n+1)
	heavy = make([]int, n+1)
	top = make([]int, n+1)
	pos = make([]int, n+1)
	seg = make([]int, 4*n+4)

	// 간선 입력
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// HLD 전처리
	dfs1(1, 0, 0)
	curPos = 0
	dfs2(1, 1)

	// DFS 번호 순서대로 값 배열 구성
	flat := make([]int, n)
	for i := 1; i <= n; i++ {
		flat[pos[i]] = val[i]
	}
	build(1, 0, n-1, flat)

	// 질의 처리
	for i := 0; i < q; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		fmt.Fprintln(writer, pathQuery(u, v))
	}
}
