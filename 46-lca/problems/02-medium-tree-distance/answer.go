package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxNode = 100001
const maxLog = 17

type edge struct {
	to, weight int
}

var adj [maxNode][]edge
var parent [maxNode][maxLog]int
var depth [maxNode]int
var dist [maxNode]int64

// buildTree는 DFS로 깊이, 거리, 희소 배열을 구성한다.
func buildTree(v, par, d int, w int64) {
	depth[v] = d
	dist[v] = w
	parent[v][0] = par
	for k := 1; k < maxLog; k++ {
		parent[v][k] = parent[parent[v][k-1]][k-1]
	}
	for _, e := range adj[v] {
		if e.to != par {
			buildTree(e.to, v, d+1, w+int64(e.weight))
		}
	}
}

// lca는 Binary Lifting으로 두 노드의 최소 공통 조상을 구한다.
func lca(u, v int) int {
	if depth[u] < depth[v] {
		u, v = v, u
	}
	diff := depth[u] - depth[v]
	for k := 0; k < maxLog; k++ {
		if (diff>>k)&1 == 1 {
			u = parent[u][k]
		}
	}
	if u == v {
		return u
	}
	for k := maxLog - 1; k >= 0; k-- {
		if parent[u][k] != parent[v][k] {
			u = parent[u][k]
			v = parent[v][k]
		}
	}
	return parent[u][0]
}

// treeDist는 가중치 트리에서 두 노드 사이의 거리를 구한다.
//
// [매개변수]
//   - u: 첫 번째 노드 번호
//   - v: 두 번째 노드 번호
//
// [반환값]
//   - int64: u와 v 사이의 거리 (간선 가중치 합)
//
// [알고리즘 힌트]
//   1. LCA를 구한다: l = lca(u, v).
//   2. dist(u, v) = dist[u] + dist[v] - 2 * dist[l]로 계산한다.
//      (dist[x]는 루트에서 x까지의 거리)
func treeDist(u, v int) int64 {
	l := lca(u, v)
	return dist[u] + dist[v] - 2*dist[l]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n-1; i++ {
		var a, b, w int
		fmt.Fscan(reader, &a, &b, &w)
		adj[a] = append(adj[a], edge{b, w})
		adj[b] = append(adj[b], edge{a, w})
	}

	buildTree(1, 0, 0, 0)

	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		fmt.Fprintln(writer, treeDist(u, v))
	}
}
