package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Edge는 그래프의 간선을 나타낸다.
type Edge struct {
	U, V, W int
}

// kruskalMST는 크루스칼 알고리즘으로 MST의 가중치 합을 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록
//
// [반환값]
//   - int: MST의 가중치 합
//
// [알고리즘 힌트]
//
//	간선을 가중치 오름차순으로 정렬한 뒤,
//	Union-Find로 사이클을 확인하며 MST를 구성한다.
func kruskalMST(n int, edges []Edge) int {
	// 간선을 가중치 오름차순 정렬 (같은 가중치면 노드 합 오름차순)
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].W != edges[j].W {
			return edges[i].W < edges[j].W
		}
		return edges[i].U+edges[i].V < edges[j].U+edges[j].V
	})

	// Union-Find 초기화
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}

	// Find 함수 (경로 압축)
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// Union 함수 (랭크 기반)
	union := func(a, b int) bool {
		ra, rb := find(a), find(b)
		if ra == rb {
			return false
		}
		if rank[ra] < rank[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		if rank[ra] == rank[rb] {
			rank[ra]++
		}
		return true
	}

	// 크루스칼 알고리즘 수행
	totalWeight := 0
	edgeCount := 0
	for _, e := range edges {
		if union(e.U, e.V) {
			totalWeight += e.W
			edgeCount++
			if edgeCount == n-1 {
				break
			}
		}
	}

	return totalWeight
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i].U, &edges[i].V, &edges[i].W)
	}

	var s int
	fmt.Fscan(reader, &s)
	_ = s

	result := kruskalMST(n, edges)
	fmt.Fprintln(writer, result)
}
