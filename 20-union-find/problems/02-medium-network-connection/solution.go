package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 가중치가 있는 간선을 나타낸다.
type Edge struct {
	u, v, w int
}

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
// 이미 같은 집합이면 false를 반환한다.
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

// minimumSpanningTreeCost는 크루스칼 알고리즘으로 MST의 총 비용을 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 가중치 간선 목록
//
// [반환값]
//   - int: MST의 총 비용 (모든 노드를 연결할 수 없으면 -1)
func minimumSpanningTreeCost(n int, edges []Edge) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 간선 입력
	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i].u, &edges[i].v, &edges[i].w)
	}

	// 핵심 함수 호출
	result := minimumSpanningTreeCost(n, edges)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
