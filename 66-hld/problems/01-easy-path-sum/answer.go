package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathSum은 HLD(Heavy-Light Decomposition)와 세그먼트 트리를 이용하여
// 트리에서 두 노드 사이의 경로 합 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - val: 각 노드의 값 (1-indexed)
//   - edges: 간선 목록 (u, v 쌍)
//   - queries: 질의 목록 (u, v 쌍)
//
// [반환값]
//   - []int: 각 질의에 대한 경로 합 결과
//
// [알고리즘 힌트]
//   1. DFS로 서브트리 크기, 깊이, 부모, Heavy Child를 계산한다
//   2. Heavy Child 우선 DFS로 체인을 구성하고 DFS 번호를 부여한다
//   3. 세그먼트 트리로 DFS 번호 순서의 값 배열에 대해 구간 합을 지원한다
//   4. 경로 질의: 두 노드가 같은 체인에 올 때까지 깊은 쪽 체인의 구간 합을 누적한다
func pathSum(n int, val []int, edges [][2]int, queries [][2]int) []int {
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	sz := make([]int, n+1)
	dep := make([]int, n+1)
	par := make([]int, n+1)
	heavy := make([]int, n+1)
	top := make([]int, n+1)
	pos := make([]int, n+1)
	seg := make([]int, 4*n+4)
	curPos := 0

	var dfs1 func(v, p, d int)
	dfs1 = func(v, p, d int) {
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

	var dfs2 func(v, chainTop int)
	dfs2 = func(v, chainTop int) {
		top[v] = chainTop
		pos[v] = curPos
		curPos++
		if heavy[v] != -1 {
			dfs2(heavy[v], chainTop)
		}
		for _, u := range adj[v] {
			if u == par[v] || u == heavy[v] {
				continue
			}
			dfs2(u, u)
		}
	}

	var build func(node, s, e int, arr []int)
	build = func(node, s, e int, arr []int) {
		if s == e {
			seg[node] = arr[s]
			return
		}
		mid := (s + e) / 2
		build(node*2, s, mid, arr)
		build(node*2+1, mid+1, e, arr)
		seg[node] = seg[node*2] + seg[node*2+1]
	}

	var query func(node, s, e, l, r int) int
	query = func(node, s, e, l, r int) int {
		if r < s || e < l {
			return 0
		}
		if l <= s && e <= r {
			return seg[node]
		}
		mid := (s + e) / 2
		return query(node*2, s, mid, l, r) + query(node*2+1, mid+1, e, l, r)
	}

	dfs1(1, 0, 0)
	dfs2(1, 1)

	flat := make([]int, n)
	for i := 1; i <= n; i++ {
		flat[pos[i]] = val[i]
	}
	build(1, 0, n-1, flat)

	pathQuery := func(u, v int) int {
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

	results := make([]int, len(queries))
	for i, qr := range queries {
		results[i] = pathQuery(qr[0], qr[1])
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	val := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &val[i])
	}

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	queries := make([][2]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	results := pathSum(n, val, edges, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
