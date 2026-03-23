package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// 차선 최소 신장 트리(Second-best MST) 풀이
// 1) 크루스칼로 MST를 구한다
// 2) MST 위에서 두 정점 사이 경로의 최대 가중치를 전처리한다
// 3) MST에 포함되지 않은 각 간선에 대해, 해당 간선을 추가하고
//    경로상 최대 가중치 간선을 제거했을 때의 비용 변화를 계산한다
// 4) 비용 증가가 최소인 경우가 차선 MST이다

// Edge 구조체는 간선 정보를 저장한다
type Edge struct {
	u, v, w int
}

var parent []int
var rankArr []int

// 유니온 파인드 초기화
func initialize(n int) {
	parent = make([]int, n+1)
	rankArr = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
}

// 경로 압축을 적용한 Find 연산
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// 랭크 기반 Union 연산
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

// AdjEdge 구조체는 인접 리스트의 간선 정보를 저장한다
type AdjEdge struct {
	to, weight int
}

// maxOnPath 함수는 MST에서 u와 v 사이 경로의 최대 가중치를 구한다
// DFS를 사용하여 경로를 탐색한다
func maxOnPath(adj [][]AdjEdge, u, v, n int) int {
	visited := make([]bool, n+1)
	// DFS로 u에서 v까지의 경로를 찾으며 최대 가중치를 추적
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

	// 크루스칼 알고리즘으로 MST 구하기
	initialize(n)
	mstWeight := 0
	inMST := make([]bool, m) // 각 간선이 MST에 포함되는지 여부

	// MST 인접 리스트 구성
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

	// MST에 포함되지 않은 각 간선에 대해 차선 MST 비용 계산
	secondBest := math.MaxInt64

	for i, e := range edges {
		if inMST[i] {
			continue
		}
		// 간선 (u, v, w)를 추가하면 사이클이 생긴다
		// 사이클에서 최대 가중치 간선을 제거하면 새로운 신장 트리가 된다
		maxW := maxOnPath(adj, e.u, e.v, n)

		// 비용 변화: 새 간선 가중치 - 제거할 간선 가중치
		diff := e.w - maxW
		if diff >= 0 {
			candidate := mstWeight + diff
			if candidate < secondBest {
				secondBest = candidate
			}
		}
	}

	fmt.Fprintln(writer, secondBest)
}
