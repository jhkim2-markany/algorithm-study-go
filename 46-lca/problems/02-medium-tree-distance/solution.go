package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxNode = 100001
const maxLog = 17

// 간선 정보 (인접 노드, 가중치)
type edge struct {
	to, weight int
}

var adj [maxNode][]edge
var parent [maxNode][maxLog]int
var depth [maxNode]int
var dist [maxNode]int64 // 루트에서 각 노드까지의 거리

// DFS로 깊이, 거리, 희소 배열을 구성한다
func dfs(v, par, d int, w int64) {
	depth[v] = d
	dist[v] = w
	parent[v][0] = par
	for k := 1; k < maxLog; k++ {
		parent[v][k] = parent[parent[v][k-1]][k-1]
	}
	for _, e := range adj[v] {
		if e.to != par {
			dfs(e.to, v, d+1, w+int64(e.weight))
		}
	}
}

// Binary Lifting으로 LCA를 구한다
func lca(u, v int) int {
	if depth[u] < depth[v] {
		u, v = v, u
	}
	// 깊이를 맞춘다
	diff := depth[u] - depth[v]
	for k := 0; k < maxLog; k++ {
		if (diff>>k)&1 == 1 {
			u = parent[u][k]
		}
	}
	if u == v {
		return u
	}
	// LCA 바로 아래까지 올린다
	for k := maxLog - 1; k >= 0; k-- {
		if parent[u][k] != parent[v][k] {
			u = parent[u][k]
			v = parent[v][k]
		}
	}
	return parent[u][0]
}

// 두 노드 사이의 거리를 구한다
// dist(u, v) = dist[u] + dist[v] - 2 * dist[LCA(u, v)]
func treeDist(u, v int) int64 {
	l := lca(u, v)
	return dist[u] + dist[v] - 2*dist[l]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 간선 입력 (가중치 포함)
	for i := 0; i < n-1; i++ {
		var a, b, w int
		fmt.Fscan(reader, &a, &b, &w)
		adj[a] = append(adj[a], edge{b, w})
		adj[b] = append(adj[b], edge{a, w})
	}

	// 전처리: DFS로 깊이, 거리, 희소 배열 구성
	dfs(1, 0, 0, 0)

	// 쿼리 처리
	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		// 두 노드 사이의 거리를 출력한다
		fmt.Fprintln(writer, treeDist(u, v))
	}
}
