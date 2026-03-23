package main

import (
	"fmt"
	"sort"
)

// 크루스칼(Kruskal) 알고리즘 - 최소 신장 트리 구하기
// Union-Find를 사용하여 사이클을 판별하며 간선을 선택한다
// 시간 복잡도: O(E log E) (간선 정렬이 지배적)
// 공간 복잡도: O(V + E)

// Edge 구조체는 간선 정보를 저장한다
type Edge struct {
	u, v, weight int
}

// parent 배열과 rank 배열 (Union-Find용)
var parent []int
var rankArr []int

// initialize 함수는 Union-Find를 초기화한다
func initialize(n int) {
	parent = make([]int, n)
	rankArr = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 자기 자신이 루트
	}
}

// find 함수는 경로 압축을 적용하여 루트를 찾는다
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// union 함수는 두 집합을 랭크 기반으로 합친다
func union(x, y int) bool {
	rootX := find(x)
	rootY := find(y)
	if rootX == rootY {
		return false // 이미 같은 집합 (사이클 발생)
	}
	if rankArr[rootX] < rankArr[rootY] {
		parent[rootX] = rootY
	} else if rankArr[rootX] > rankArr[rootY] {
		parent[rootY] = rootX
	} else {
		parent[rootY] = rootX
		rankArr[rootX]++
	}
	return true
}

// kruskal 함수는 크루스칼 알고리즘으로 MST를 구한다
func kruskal(n int, edges []Edge) (int, []Edge) {
	// 간선을 가중치 오름차순으로 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	initialize(n)

	totalWeight := 0
	mstEdges := []Edge{}

	for _, e := range edges {
		// 두 정점이 다른 집합에 속하면 간선을 MST에 추가
		if union(e.u, e.v) {
			totalWeight += e.weight
			mstEdges = append(mstEdges, e)
			// V-1개의 간선을 선택하면 종료
			if len(mstEdges) == n-1 {
				break
			}
		}
	}

	return totalWeight, mstEdges
}

func main() {
	// 예제 그래프: 5개 정점, 7개 간선
	//     1
	//   0---1
	//   |\ /|
	//  4| X |3
	//   |/ \|
	//   2---3
	//     2
	//   3---4 (가중치 5)
	//   2---4 (가중치 6)
	n := 5
	edges := []Edge{
		{0, 1, 1},
		{0, 2, 4},
		{1, 2, 2},
		{1, 3, 3},
		{2, 3, 2},
		{3, 4, 5},
		{2, 4, 6},
	}

	fmt.Println("=== 크루스칼 알고리즘 ===")
	fmt.Printf("정점 수: %d, 간선 수: %d\n\n", n, len(edges))

	totalWeight, mstEdges := kruskal(n, edges)

	fmt.Println("MST에 포함된 간선:")
	for _, e := range mstEdges {
		fmt.Printf("  %d -- %d (가중치: %d)\n", e.u, e.v, e.weight)
	}
	fmt.Printf("\nMST 총 가중치: %d\n", totalWeight)
}
