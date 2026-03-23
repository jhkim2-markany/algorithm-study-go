package main

import (
	"bufio"
	"fmt"
	"os"
)

// coloredPathQuery는 센트로이드 분할로 거리별 같은 색상 쌍의 수를 전처리하고,
// 각 쿼리에 대해 거리 k에서의 같은 색상 쌍 수를 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - color: 각 노드의 색상 (1-indexed)
//   - edges: 간선 목록 (u, v 쌍)
//   - queries: 거리 쿼리 목록
//
// [반환값]
//   - []int64: 각 쿼리에 대한 결과 (거리 k에서 같은 색상인 쌍의 수)
//
// [알고리즘 힌트]
//   1. 센트로이드를 찾아 제거하고, 센트로이드를 지나는 경로를 처리한다
//   2. 각 서브트리에서 (거리, 색상) 정보를 수집한다
//   3. 같은 색상 내에서 모든 쌍의 거리 합을 구하여 ans 배열에 누적한다
//   4. 같은 서브트리 내의 쌍은 포함-배제로 빼준다
//   5. 쿼리마다 ans[k]를 O(1)로 반환한다
func coloredPathQuery(n int, color []int, edges [][2]int, queries []int) []int64 {
	const MAXN = 50001

	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	sz := make([]int, n+1)
	removed := make([]bool, n+1)
	ans := make([]int64, MAXN)

	type NodeInfo struct {
		dist, color int
	}

	var calcSize func(v, p int) int
	calcSize = func(v, p int) int {
		sz[v] = 1
		for _, u := range adj[v] {
			if u == p || removed[u] {
				continue
			}
			sz[v] += calcSize(u, v)
		}
		return sz[v]
	}

	var findCentroid func(v, p, treeSize int) int
	findCentroid = func(v, p, treeSize int) int {
		for _, u := range adj[v] {
			if u == p || removed[u] {
				continue
			}
			if sz[u] > treeSize/2 {
				return findCentroid(u, v, treeSize)
			}
		}
		return v
	}

	var getNodes func(v, p, dist int, nodes *[]NodeInfo)
	getNodes = func(v, p, dist int, nodes *[]NodeInfo) {
		*nodes = append(*nodes, NodeInfo{dist, color[v]})
		for _, u := range adj[v] {
			if u == p || removed[u] {
				continue
			}
			getNodes(u, v, dist+1, nodes)
		}
	}

	countColorPairs := func(nodes []NodeInfo, sign int64) {
		colorDist := make(map[int][]int)
		for _, nd := range nodes {
			colorDist[nd.color] = append(colorDist[nd.color], nd.dist)
		}
		distCount := make(map[int]int64)
		for _, dists := range colorDist {
			if len(dists) < 2 {
				continue
			}
			for i := 0; i < len(dists); i++ {
				for j := i + 1; j < len(dists); j++ {
					distCount[dists[i]+dists[j]]++
				}
			}
		}
		for d, cnt := range distCount {
			if d < MAXN {
				ans[d] += sign * cnt
			}
		}
	}

	var decompose func(v int)
	decompose = func(v int) {
		treeSize := calcSize(v, -1)
		centroid := findCentroid(v, -1, treeSize)
		removed[centroid] = true

		allNodes := []NodeInfo{{0, color[centroid]}}
		for _, u := range adj[centroid] {
			if removed[u] {
				continue
			}
			subNodes := []NodeInfo{}
			getNodes(u, centroid, 1, &subNodes)
			countColorPairs(subNodes, -1)
			allNodes = append(allNodes, subNodes...)
		}
		countColorPairs(allNodes, 1)

		for _, u := range adj[centroid] {
			if removed[u] {
				continue
			}
			decompose(u)
		}
	}

	decompose(1)

	results := make([]int64, len(queries))
	for i, k := range queries {
		results[i] = ans[k]
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	color := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &color[i])
	}

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	queries := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i])
	}

	results := coloredPathQuery(n, color, edges, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
