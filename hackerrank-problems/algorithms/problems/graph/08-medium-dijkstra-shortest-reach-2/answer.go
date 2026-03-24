package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

// Item은 우선순위 큐의 원소를 나타낸다.
type Item struct {
	node, dist int
}

// PQ는 최소 힙 기반 우선순위 큐이다.
type PQ []Item

func (pq PQ) Len() int            { return len(pq) }
func (pq PQ) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// dijkstra는 시작 노드에서 모든 노드까지의 최단 거리를 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 원소는 [3]int{u, v, w})
//   - s: 시작 노드 (1-indexed)
//
// [반환값]
//   - []int: 시작 노드를 제외한 각 노드까지의 최단 거리 (-1은 도달 불가)
//
// [알고리즘 힌트]
//
//	인접 리스트를 구성하고 우선순위 큐(최소 힙)로 다익스트라를 수행한다.
//	중복 간선은 최소 가중치만 사용한다.
func dijkstra(n int, edges [][3]int, s int) []int {
	// 인접 리스트 구성
	type Edge struct {
		to, w int
	}
	adj := make([][]Edge, n+1)
	for i := range adj {
		adj[i] = []Edge{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], Edge{e[0], e[2]})
	}

	// 거리 배열 초기화
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[s] = 0

	// 우선순위 큐로 다익스트라 수행
	pq := &PQ{{s, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		// 이미 더 짧은 경로가 확정된 경우 건너뛰기
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range adj[cur.node] {
			newDist := dist[cur.node] + e.w
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				heap.Push(pq, Item{e.to, newDist})
			}
		}
	}

	// 시작 노드를 제외한 결과 구성
	result := make([]int, 0, n-1)
	for i := 1; i <= n; i++ {
		if i == s {
			continue
		}
		if dist[i] == math.MaxInt64 {
			result = append(result, -1)
		} else {
			result = append(result, dist[i])
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		edges := make([][3]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
		}

		var s int
		fmt.Fscan(reader, &s)

		result := dijkstra(n, edges, s)
		for i, v := range result {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
