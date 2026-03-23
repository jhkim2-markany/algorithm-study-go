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

// AdjEdge는 인접 리스트의 간선 정보를 저장하는 구조체이다
type AdjEdge struct {
	to, weight int
}

var parent []int
var rankArr []int

// initialize는 유니온 파인드를 초기화한다
func initialize(n int) {
	parent = make([]int, n+1)
	rankArr = make([]int, n+1)
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

// secondBestMST는 차선 최소 신장 트리의 비용을 구한다.
//
// [매개변수]
//   - n: 정점의 수
//   - m: 간선의 수
//   - edges: 가중치 오름차순으로 정렬된 간선 목록
//
// [반환값]
//   - int: 차선 최소 신장 트리의 총 비용
func secondBestMST(n, m int, edges []Edge) int {
	// 여기에 코드를 작성하세요
	return 0
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

	// 핵심 함수 호출
	result := secondBestMST(n, m, edges)

	fmt.Fprintln(writer, result)
}
