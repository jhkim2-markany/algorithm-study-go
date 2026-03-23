package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 간선 정보를 저장하는 구조체
type Edge struct {
	u, v, w int
}

var parent []int
var rank_ []int

// 유니온 파인드 초기화
func initialize(n int) {
	parent = make([]int, n+1)
	rank_ = make([]int, n+1)
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 간선 입력 받기
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

	// 유니온 파인드 초기화
	initialize(n)

	// 크루스칼 알고리즘으로 MST 구하기
	totalWeight := 0
	mstEdges := []Edge{}

	for _, e := range edges {
		// 사이클이 생기지 않는 간선만 선택
		if union(e.u, e.v) {
			totalWeight += e.w
			mstEdges = append(mstEdges, e)
			if len(mstEdges) == n-1 {
				break
			}
		}
	}

	// 총 가중치 출력
	fmt.Fprintln(writer, totalWeight)

	// MST 간선 출력 (선택 순서대로)
	for _, e := range mstEdges {
		fmt.Fprintf(writer, "%d %d %d\n", e.u, e.v, e.w)
	}
}
