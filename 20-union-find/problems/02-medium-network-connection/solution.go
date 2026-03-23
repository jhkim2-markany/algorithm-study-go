package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 크루스칼 알고리즘으로 최소 신장 트리를 구하는 풀이
// 유니온 파인드를 활용하여 사이클 여부를 판별한다

type Edge struct {
	u, v, w int
}

var parent []int
var rank_ []int

// find 함수는 x의 루트를 반환한다 (경로 압축 적용)
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// union 함수는 x와 y가 속한 집합을 합친다 (랭크 기반)
// 이미 같은 집합이면 false를 반환한다
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

	// 간선 입력
	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i].u, &edges[i].v, &edges[i].w)
	}

	// 크루스칼 알고리즘: 간선을 비용 기준 오름차순 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// 유니온 파인드 초기화
	parent = make([]int, n+1)
	rank_ = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	// 비용이 작은 간선부터 선택하며 MST 구성
	totalCost := 0
	edgeCount := 0
	for _, e := range edges {
		// 두 노드가 다른 집합에 속하면 간선을 선택
		if union(e.u, e.v) {
			totalCost += e.w
			edgeCount++
			// N-1개의 간선을 선택하면 MST 완성
			if edgeCount == n-1 {
				break
			}
		}
	}

	// 모든 노드를 연결할 수 없는 경우
	if edgeCount < n-1 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, totalCost)
	}
}
