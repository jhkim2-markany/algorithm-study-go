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

// Union은 x와 y가 속한 집합을 합친다 (랭크 기반).
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

// countComponents는 유니온 파인드를 사용하여 연결 요소 개수를 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 간선은 [2]int{u, v})
//
// [반환값]
//   - int: 연결 요소의 개수
func countComponents(n int, edges [][2]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	// 핵심 함수 호출
	result := countComponents(n, edges)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
