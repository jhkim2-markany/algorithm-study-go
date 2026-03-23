package main

import (
	"bufio"
	"fmt"
	"os"
)

// HLD를 이용한 경로 최대 간선 가중치 질의 (간선 갱신 포함)
// 간선 가중치를 자식 노드에 매핑하여 처리한다.
// 시간 복잡도: 전처리 O(N), 연산 O(Q log²N)

var (
	adj      [][]struct{ to, idx int } // 인접 리스트 (간선 번호 포함)
	sz       []int
	dep      []int
	par      []int
	heavy    []int
	top      []int
	pos      []int
	seg      []int
	edgeNode []int // edgeNode[i] = i번째 간선의 자식 노드
	curPos   int
	nn       int // 노드 수
)

func dfs1(v, p, d int) {
	par[v] = p
	dep[v] = d
	sz[v] = 1
	heavy[v] = -1
	maxSz := 0
	for _, e := range adj[v] {
		u := e.to
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
	for _, e := range adj[v] {
		u := e.to
		if u == par[v] || u == heavy[v] {
			continue
		}
		dfs2(u, u)
	}
}

// --- 세그먼트 트리 (최댓값) ---

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func build(node, s, e int, arr []int) {
	if s == e {
		seg[node] = arr[s]
		return
	}
	mid := (s + e) / 2
	build(node*2, s, mid, arr)
	build(node*2+1, mid+1, e, arr)
	seg[node] = max(seg[node*2], seg[node*2+1])
}

func segUpdate(node, s, e, idx, val int) {
	if s == e {
		seg[node] = val
		return
	}
	mid := (s + e) / 2
	if idx <= mid {
		segUpdate(node*2, s, mid, idx, val)
	} else {
		segUpdate(node*2+1, mid+1, e, idx, val)
	}
	seg[node] = max(seg[node*2], seg[node*2+1])
}

func segQuery(node, s, e, l, r int) int {
	if r < s || e < l {
		return 0 // 최댓값의 항등원 (가중치가 양수이므로 0)
	}
	if l <= s && e <= r {
		return seg[node]
	}
	mid := (s + e) / 2
	return max(segQuery(node*2, s, mid, l, r), segQuery(node*2+1, mid+1, e, l, r))
}

// 경로 최대 간선 가중치 질의
// 간선 가중치는 자식 노드에 매핑되어 있으므로, LCA에서는 제외한다
func pathMax(u, v int) int {
	res := 0
	for top[u] != top[v] {
		if dep[top[u]] < dep[top[v]] {
			u, v = v, u
		}
		res = max(res, segQuery(1, 0, nn-1, pos[top[u]], pos[u]))
		u = par[top[u]]
	}
	if dep[u] > dep[v] {
		u, v = v, u
	}
	// 같은 체인 내: pos[u]+1부터 질의 (u는 LCA, 간선은 자식 노드에 매핑)
	if pos[u]+1 <= pos[v] {
		res = max(res, segQuery(1, 0, nn-1, pos[u]+1, pos[v]))
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &nn)

	adj = make([][]struct{ to, idx int }, nn+1)
	sz = make([]int, nn+1)
	dep = make([]int, nn+1)
	par = make([]int, nn+1)
	heavy = make([]int, nn+1)
	top = make([]int, nn+1)
	pos = make([]int, nn+1)
	seg = make([]int, 4*nn+4)
	edgeNode = make([]int, nn) // 간선 번호 → 자식 노드

	type edge struct{ a, b, w int }
	edges := make([]edge, nn-1)

	// 간선 입력
	for i := 0; i < nn-1; i++ {
		fmt.Fscan(reader, &edges[i].a, &edges[i].b, &edges[i].w)
		adj[edges[i].a] = append(adj[edges[i].a], struct{ to, idx int }{edges[i].b, i})
		adj[edges[i].b] = append(adj[edges[i].b], struct{ to, idx int }{edges[i].a, i})
	}

	// HLD 전처리
	dfs1(1, 0, 0)
	curPos = 0
	dfs2(1, 1)

	// 간선 가중치를 자식 노드에 매핑
	// 간선 (a, b)에서 깊이가 더 깊은 쪽이 자식 노드
	flat := make([]int, nn)
	for i := 0; i < nn-1; i++ {
		child := edges[i].a
		if dep[edges[i].b] > dep[edges[i].a] {
			child = edges[i].b
		}
		edgeNode[i] = child
		flat[pos[child]] = edges[i].w
	}
	build(1, 0, nn-1, flat)

	// 연산 처리
	var q int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)
		if op == 1 {
			// 간선 가중치 변경
			var idx, w int
			fmt.Fscan(reader, &idx, &w)
			idx-- // 0-indexed
			segUpdate(1, 0, nn-1, pos[edgeNode[idx]], w)
		} else {
			// 경로 최대 간선 가중치 질의
			var u, v int
			fmt.Fscan(reader, &u, &v)
			fmt.Fprintln(writer, pathMax(u, v))
		}
	}
}
