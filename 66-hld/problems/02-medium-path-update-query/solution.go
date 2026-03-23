package main

import (
	"bufio"
	"fmt"
	"os"
)

// HLD + Lazy Propagation 세그먼트 트리를 이용한 경로 갱신/질의
// 시간 복잡도: 전처리 O(N), 연산 O(Q log²N)

var (
	adj    [][]int
	sz     []int
	dep    []int
	par    []int
	heavy  []int
	top    []int
	pos    []int
	seg    []int64 // 세그먼트 트리 (합)
	lazy   []int64 // 지연 전파 배열
	curPos int
	n, q   int
)

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

func dfs2(v, chainTop int) {
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

// --- Lazy Propagation 세그먼트 트리 ---

// 지연 값을 자식에게 전파한다
func pushDown(node, s, e int) {
	if lazy[node] != 0 {
		mid := (s + e) / 2
		seg[node*2] += lazy[node] * int64(mid-s+1)
		lazy[node*2] += lazy[node]
		seg[node*2+1] += lazy[node] * int64(e-mid)
		lazy[node*2+1] += lazy[node]
		lazy[node] = 0
	}
}

// 구간 [l, r]에 val을 더한다
func update(node, s, e, l, r int, val int64) {
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

// 구간 [l, r]의 합을 반환한다
func query(node, s, e, l, r int) int64 {
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

// 경로 갱신: 노드 u에서 v까지의 경로에 val을 더한다
func pathUpdate(u, v int, val int64) {
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

// 경로 질의: 노드 u에서 v까지의 경로 합
func pathQuery(u, v int) int64 {
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &q)

	adj = make([][]int, n+1)
	sz = make([]int, n+1)
	dep = make([]int, n+1)
	par = make([]int, n+1)
	heavy = make([]int, n+1)
	top = make([]int, n+1)
	pos = make([]int, n+1)
	seg = make([]int64, 4*n+4)
	lazy = make([]int64, 4*n+4)

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

	// 연산 처리
	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)
		if op == 1 {
			// 경로 갱신
			var u, v, w int
			fmt.Fscan(reader, &u, &v, &w)
			pathUpdate(u, v, int64(w))
		} else {
			// 경로 질의
			var u, v int
			fmt.Fscan(reader, &u, &v)
			fmt.Fprintln(writer, pathQuery(u, v))
		}
	}
}
