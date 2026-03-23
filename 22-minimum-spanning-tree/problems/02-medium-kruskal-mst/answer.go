package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Edge는 간선 정보를 저장하는 구조체이다
type Edge struct {
	u, v, w int
}

var parent []int
var rank_ []int

// initialize는 유니온 파인드를 초기화한다
func initialize(n int) {
	parent = make([]int, n+1)
	rank_ = make([]int, n+1)
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
	rootX := find(x)
	rootY := find(y)
	if rootX == rootY {
		return false
	}
	if rank_[rootX] < rank_[rootY] {
		parent[rootX] = rootY
	} else if rank_[rootX] > rank_[rootY] {
		parent[rootY] = rootX
	} else {
		parent[rootY] = rootX
		rank_[rootX]++
	}
	return true
}

// kruskalMST는 크루스칼 알고리즘으로 MST의 총 가중치와 선택된 간선 목록을 구한다.
//
// [매개변수]
//   - n: 정점의 수
//   - edges: 정렬된 간선 목록 (u, v, w)
//
// [반환값]
//   - int: MST의 총 가중치
//   - []Edge: MST에 포함된 간선 목록 (선택 순서대로)
//
// [알고리즘 힌트]
//
//	크루스칼 알고리즘: 가중치가 작은 간선부터 순서대로 선택하되,
//	유니온 파인드로 사이클 여부를 확인하여 사이클이 생기지 않는 간선만 추가한다.
//	N-1개의 간선을 선택하면 MST가 완성된다.
//	선택된 간선을 별도 목록에 저장하여 반환한다.
func kruskalMST(n int, edges []Edge) (int, []Edge) {
	initialize(n)

	totalWeight := 0
	mstEdges := []Edge{}

	for _, e := range edges {
		if union(e.u, e.v) {
			totalWeight += e.w
			mstEdges = append(mstEdges, e)
			if len(mstEdges) == n-1 {
				break
			}
		}
	}
	return totalWeight, mstEdges
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

	// 간선 정렬: 가중치 오름차순, 같으면 u 오름차순, 같으면 v 오름차순
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].w != edges[j].w {
			return edges[i].w < edges[j].w
		}
		if edges[i].u != edges[j].u {
			return edges[i].u < edges[j].u
		}
		return edges[i].v < edges[j].v
	})

	// 핵심 함수 호출
	totalWeight, mstEdges := kruskalMST(n, edges)

	// 총 가중치 출력
	fmt.Fprintln(writer, totalWeight)

	// MST 간선 출력
	for _, e := range mstEdges {
		fmt.Fprintf(writer, "%d %d %d\n", e.u, e.v, e.w)
	}
}
