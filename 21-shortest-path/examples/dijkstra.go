package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Dijkstra 알고리즘 - 우선순위 큐를 이용한 단일 출발점 최단 경로
// 시간 복잡도: O((V + E) log V)
// 공간 복잡도: O(V + E)

// Edge는 인접 정점과 가중치를 나타낸다
type Edge struct {
	to, weight int
}

// Item은 우선순위 큐에 저장되는 (거리, 정점) 쌍이다
type Item struct {
	dist, node int
}

// PQ는 최소 힙 기반 우선순위 큐이다
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

// dijkstra 함수는 출발 정점에서 모든 정점까지의 최단 거리를 반환한다
func dijkstra(graph [][]Edge, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0

	// 우선순위 큐에 출발 정점 삽입
	pq := &PQ{{0, start}}
	heap.Init(pq)

	for pq.Len() > 0 {
		// 거리가 가장 짧은 정점을 꺼낸다
		cur := heap.Pop(pq).(Item)
		u := cur.node

		// 이미 더 짧은 경로가 확정되었으면 건너뛴다
		if cur.dist > dist[u] {
			continue
		}

		// 인접 정점에 대해 완화(relaxation) 수행
		for _, e := range graph[u] {
			newDist := dist[u] + e.weight
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				heap.Push(pq, Item{newDist, e.to})
			}
		}
	}

	return dist
}

func main() {
	// 그래프 생성 (정점 5개, 0~4)
	// 0 --1--> 1 --3--> 3
	// 0 --4--> 2 --1--> 3
	// 1 --2--> 2
	// 3 --2--> 4
	graph := make([][]Edge, 5)
	graph[0] = []Edge{{1, 1}, {2, 4}}
	graph[1] = []Edge{{2, 2}, {3, 3}}
	graph[2] = []Edge{{3, 1}}
	graph[3] = []Edge{{4, 2}}

	// 정점 0에서 출발하는 최단 거리 계산
	dist := dijkstra(graph, 0)

	fmt.Println("=== Dijkstra 최단 경로 ===")
	for i, d := range dist {
		if d == math.MaxInt64 {
			fmt.Printf("0 → %d: 도달 불가\n", i)
		} else {
			fmt.Printf("0 → %d: %d\n", i, d)
		}
	}
}
