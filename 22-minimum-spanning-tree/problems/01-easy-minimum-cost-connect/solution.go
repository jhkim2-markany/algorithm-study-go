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

	// 간선을 가중치 오름차순으로 정렬 (크루스칼 알고리즘)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// 유니온 파인드 초기화
	initialize(n)

	// 최소 비용 계산
	totalCost := 0
	edgeCount := 0
	for _, e := range edges {
		if union(e.u, e.v) {
			totalCost += e.w
			edgeCount++
			// N-1개의 간선을 선택하면 종료
			if edgeCount == n-1 {
				break
			}
		}
	}

	fmt.Fprintln(writer, totalCost)
}
