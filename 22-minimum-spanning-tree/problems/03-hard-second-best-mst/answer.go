package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Edge는 간선 정보를 저장하는 구조체이다
type Edge struct {
	u, v, w int
}

// AdjEdge는 인접 리스트의 간선 정보를 저장하는 구조체이다
type AdjEdge struct {
	to, weight int
}

var parent []int
var rankArr []int

// initialize는 유니온 파인드를 초기화한다
func initialize(n int) {
	parent = make([]int, n+1)
	rankArr = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
}

// find는 경로 압축을 적용한 Find 연산이다
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// union은 랭크 기반 Union 연산이다
func union(x, y int) bool {
	rx, ry := find(x), find(y)
	if rx == ry {
		return false
	}
	if rankArr[rx] < rankArr[ry] {
		parent[rx] = ry
	} else if rankArr[rx] > rankArr[ry] {
		parent[ry] = rx
	} else {
		parent[ry] = rx
		rankArr[rx]++
	}
	return true
}

// maxOnPath는 MST에서 u와 v 사이 경로의 최대 가중치를 구한다
func maxOnPath(adj [][]AdjEdge, u, v, n int) int {
	visited := make([]bool, n+1)
	var dfs func(cur, target, curMax int) (int, bool)
	dfs = func(cur, target, curMax int) (int, bool) {
		if cur == target {
			return curMax, true
		}
		visited[cur] = true
		for _, e := range adj[cur] {
			if !visited[e.to] {
				newMax := curMax
				if e.weight > newMax {
					newMax = e.weight
				}
				if result, found := dfs(e.to, target, newMax); found {
					return result, true
				}
			}
		}
		return 0, false
	}
	result, _ := dfs(u, v, 0)
	return result
}

// secondBestMST는 차선 최소 신장 트리의 비용을 구한다.
//
// [매개변수]
//   - n: 정점의 수
//   - m: 간선의 수
//   - edges: 가중치 오름차순으로 정렬된 간선 목록
//
// [반환값]
//   - int: 차선 최소 신장 트리의 총 비용
//
// [알고리즘 힌트]
//
//  1. 크루스칼로 MST를 구한다.
//  2. MST 위에서 두 정점 사이 경로의 최대 가중치를 DFS로 구한다.
//  3. MST에 포함되지 않은 각 간선에 대해, 해당 간선을 추가하고
//     경로상 최대 가중치 간선을 제거했을 때의 비용 변화를 계산한다.
//  4. 비용 증가가 최소인 경우가 차선 MST이다.
func secondBestMST(n, m int, edges []Edge) int {
	initialize(n)
	mstWeight := 0
	inMST := make([]bool, m)

	adj := make([][]AdjEdge, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []AdjEdge{}
	}

	for i, e := range edges {
		if union(e.u, e.v) {
			mstWeight += e.w
			inMST[i] = true
			adj[e.u] = append(adj[e.u], AdjEdge{e.v, e.w})
			adj[e.v] = append(adj[e.v], AdjEdge{e.u, e.w})
		}
	}

	secondBest := math.MaxInt64

	for i, e := range edges {
		if inMST[i] {
			continue
		}
		maxW := maxOnPath(adj, e.u, e.v, n)
		diff := e.w - maxW
		if diff >= 0 {
			candidate := mstWeight + diff
			if candidate < secondBest {
				secondBest = candidate
			}
		}
	}

	return secondBest
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i].u, &edges[i].v, &edges[i].w)
	}

	// 간선을 가중치 오름차순으로 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// 핵심 함수 호출
	result := secondBestMST(n, m, edges)

	fmt.Fprintln(writer, result)
}
