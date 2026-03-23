package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathMaxEdge는 HLD와 세그먼트 트리(최댓값)를 이용하여 트리에서
// 간선 가중치 갱신과 경로 최대 간선 가중치 질의를 처리한다.
// 간선 가중치는 자식 노드에 매핑하여 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - edgeList: 간선 목록 (a, b, w)
//   - ops: 연산 목록 (op=1: 간선 갱신 [1,idx,w], op=2: 경로 질의 [2,u,v])
//
// [반환값]
//   - []int: 경로 최대 간선 가중치 질의(op=2)의 결과 배열
//
// [알고리즘 힌트]
//   1. HLD로 트리를 체인으로 분해하고 DFS 번호를 부여한다
//   2. 간선 (a,b)에서 깊이가 더 깊은 쪽을 자식 노드로 하여 간선 가중치를 매핑한다
//   3. 세그먼트 트리로 점 갱신과 구간 최댓값 질의를 지원한다
//   4. 경로 질의 시 LCA에서는 간선을 제외한다 (pos[u]+1부터 질의)
func pathMaxEdge(n int, edgeList [][3]int, ops [][]int) []int {
	type adjEdge struct {
		to, idx int
	}
	adj := make([][]adjEdge, n+1)
	for i, e := range edgeList {
		adj[e[0]] = append(adj[e[0]], adjEdge{e[1], i})
		adj[e[1]] = append(adj[e[1]], adjEdge{e[0], i})
	}

	sz := make([]int, n+1)
	dep := make([]int, n+1)
	par := make([]int, n+1)
	heavy := make([]int, n+1)
	top := make([]int, n+1)
	pos := make([]int, n+1)
	seg := make([]int, 4*n+4)
	edgeNode := make([]int, n)
	curPos := 0

	var dfs1 func(v, p, d int)
	dfs1 = func(v, p, d int) {
		par[v] = p
		dep[v] = d
		sz[v] = 1
		heavy[v] = -1
		maxSz := 0
		for _, e := range adj[v] {
			if e.to == p {
				continue
			}
			dfs1(e.to, v, d+1)
			sz[v] += sz[e.to]
			if sz[e.to] > maxSz {
				maxSz = sz[e.to]
				heavy[v] = e.to
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
		for _, e := range adj[v] {
			if e.to == par[v] || e.to == heavy[v] {
				continue
			}
			dfs2(e.to, e.to)
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
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
		seg[node] = max(seg[node*2], seg[node*2+1])
	}

	var segUpdate func(node, s, e, idx, val int)
	segUpdate = func(node, s, e, idx, val int) {
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

	var segQuery func(node, s, e, l, r int) int
	segQuery = func(node, s, e, l, r int) int {
		if r < s || e < l {
			return 0
		}
		if l <= s && e <= r {
			return seg[node]
		}
		mid := (s + e) / 2
		return max(segQuery(node*2, s, mid, l, r), segQuery(node*2+1, mid+1, e, l, r))
	}

	dfs1(1, 0, 0)
	dfs2(1, 1)

	flat := make([]int, n)
	for i, e := range edgeList {
		child := e[0]
		if dep[e[1]] > dep[e[0]] {
			child = e[1]
		}
		edgeNode[i] = child
		flat[pos[child]] = e[2]
	}
	build(1, 0, n-1, flat)

	pathMax := func(u, v int) int {
		res := 0
		for top[u] != top[v] {
			if dep[top[u]] < dep[top[v]] {
				u, v = v, u
			}
			res = max(res, segQuery(1, 0, n-1, pos[top[u]], pos[u]))
			u = par[top[u]]
		}
		if dep[u] > dep[v] {
			u, v = v, u
		}
		if pos[u]+1 <= pos[v] {
			res = max(res, segQuery(1, 0, n-1, pos[u]+1, pos[v]))
		}
		return res
	}

	var results []int
	for _, op := range ops {
		if op[0] == 1 {
			idx := op[1] - 1
			segUpdate(1, 0, n-1, pos[edgeNode[idx]], op[2])
		} else {
			results = append(results, pathMax(op[1], op[2]))
		}
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	edgeList := make([][3]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edgeList[i][0], &edgeList[i][1], &edgeList[i][2])
	}

	var q int
	fmt.Fscan(reader, &q)

	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)
		if op == 1 {
			var idx, w int
			fmt.Fscan(reader, &idx, &w)
			ops[i] = []int{op, idx, w}
		} else {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			ops[i] = []int{op, u, v}
		}
	}

	results := pathMaxEdge(n, edgeList, ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
