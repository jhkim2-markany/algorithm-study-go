package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathUpdateQuery는 HLD와 Lazy Propagation 세그먼트 트리를 이용하여
// 트리에서 경로 갱신(구간 덧셈)과 경로 합 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (u, v 쌍)
//   - ops: 연산 목록 (op=1: 경로 갱신 [1,u,v,w], op=2: 경로 질의 [2,u,v])
//
// [반환값]
//   - []int64: 경로 합 질의(op=2)의 결과 배열
//
// [알고리즘 힌트]
//   1. HLD로 트리를 체인으로 분해하고 DFS 번호를 부여한다
//   2. Lazy Propagation 세그먼트 트리로 구간 덧셈과 구간 합을 지원한다
//   3. 경로 갱신: 두 노드가 같은 체인에 올 때까지 깊은 쪽 체인의 구간에 값을 더한다
//   4. 경로 질의: 같은 방식으로 체인을 올라가며 구간 합을 누적한다
func pathUpdateQuery(n int, edges [][2]int, ops [][]int) []int64 {
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
	seg := make([]int64, 4*n+4)
	lazy := make([]int64, 4*n+4)
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

	pushDown := func(node, s, e int) {
		if lazy[node] != 0 {
			mid := (s + e) / 2
			seg[node*2] += lazy[node] * int64(mid-s+1)
			lazy[node*2] += lazy[node]
			seg[node*2+1] += lazy[node] * int64(e-mid)
			lazy[node*2+1] += lazy[node]
			lazy[node] = 0
		}
	}

	var update func(node, s, e, l, r int, val int64)
	update = func(node, s, e, l, r int, val int64) {
		if r < s || e < l {
			return
		}
		if l <= s && e <= r {
			seg[node] += val * int64(e-s+1)
			lazy[node] += val
			return
		}
		pushDown(node, s, e)
		mid := (s + e) / 2
		update(node*2, s, mid, l, r, val)
		update(node*2+1, mid+1, e, l, r, val)
		seg[node] = seg[node*2] + seg[node*2+1]
	}

	var query func(node, s, e, l, r int) int64
	query = func(node, s, e, l, r int) int64 {
		if r < s || e < l {
			return 0
		}
		if l <= s && e <= r {
			return seg[node]
		}
		pushDown(node, s, e)
		mid := (s + e) / 2
		return query(node*2, s, mid, l, r) + query(node*2+1, mid+1, e, l, r)
	}

	dfs1(1, 0, 0)
	dfs2(1, 1)

	pathUpdate := func(u, v int, val int64) {
		for top[u] != top[v] {
			if dep[top[u]] < dep[top[v]] {
				u, v = v, u
			}
			update(1, 0, n-1, pos[top[u]], pos[u], val)
			u = par[top[u]]
		}
		if dep[u] > dep[v] {
			u, v = v, u
		}
		update(1, 0, n-1, pos[u], pos[v], val)
	}

	pathQuery := func(u, v int) int64 {
		var res int64
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

	var results []int64
	for _, op := range ops {
		if op[0] == 1 {
			pathUpdate(op[1], op[2], int64(op[3]))
		} else {
			results = append(results, pathQuery(op[1], op[2]))
		}
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)
		if op == 1 {
			var u, v, w int
			fmt.Fscan(reader, &u, &v, &w)
			ops[i] = []int{op, u, v, w}
		} else {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			ops[i] = []int{op, u, v}
		}
	}

	results := pathUpdateQuery(n, edges, ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
