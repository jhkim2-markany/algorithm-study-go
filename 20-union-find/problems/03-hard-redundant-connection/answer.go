package main

import (
	"bufio"
	"fmt"
	"os"
)

// UnionFind는 유니온 파인드 자료구조이다.
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind는 n개의 노드를 가진 유니온 파인드를 생성한다.
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n+1),
		rank:   make([]int, n+1),
	}
	for i := 1; i <= n; i++ {
		uf.parent[i] = i
	}
	return uf
}

// Find는 x의 루트를 반환한다 (경로 압축 적용).
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union은 x와 y가 속한 집합을 합친다.
// 이미 같은 집합이면 false를 반환한다 (사이클 탐지).
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return false
	}
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

// findRedundantConnection은 트리에 하나의 간선이 추가된 그래프에서 여분의 간선을 찾는다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 간선은 [2]int{u, v}, 순서대로 주어짐)
//
// [반환값]
//   - [2]int: 사이클을 형성하는 여분의 간선 {u, v}
//
// [알고리즘 힌트]
//
//	간선을 순서대로 추가하며 유니온 파인드로 사이클을 탐지한다.
//	Union이 실패하면(이미 같은 집합) 해당 간선이 사이클을 형성하는 여분의 간선이다.
//	여러 개가 있을 수 있으므로 가장 마지막에 발견된 간선을 반환한다.
func findRedundantConnection(n int, edges [][2]int) [2]int {
	uf := NewUnionFind(n)

	var result [2]int
	for _, e := range edges {
		if !uf.Union(e[0], e[1]) {
			result = e
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	edges := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	// 핵심 함수 호출
	result := findRedundantConnection(n, edges)

	// 결과 출력
	fmt.Fprintf(writer, "%d %d\n", result[0], result[1])
}
