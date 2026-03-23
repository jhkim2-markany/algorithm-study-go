package main

import (
	"container/heap"
	"fmt"
)

// 프림(Prim) 알고리즘 - 최소 신장 트리 구하기
// 우선순위 큐를 사용하여 현재 트리에서 가장 가까운 정점을 확장한다
// 시간 복잡도: O((V + E) log V)
// 공간 복잡도: O(V + E)

// Item 구조체는 우선순위 큐의 원소를 나타낸다
type Item struct {
	node, weight, from int
}

// PriorityQueue는 최소 힙 기반 우선순위 큐이다
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].weight < pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Edge 구조체는 인접 리스트의 간선 정보를 저장한다
type Edge struct {
	to, weight int
}

// prim 함수는 프림 알고리즘으로 MST를 구한다
func prim(n int, adj [][]Edge) (int, [][3]int) {
	visited := make([]bool, n)
	totalWeight := 0
	mstEdges := [][3]int{} // {from, to, weight}

	// 우선순위 큐 초기화: 시작 정점 0에서 출발
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: 0, weight: 0, from: -1})

	for pq.Len() > 0 {
		// 가중치가 가장 작은 간선을 꺼낸다
		item := heap.Pop(pq).(*Item)
		u := item.node

		// 이미 방문한 정점이면 건너뛴다
		if visited[u] {
			continue
		}

		// 정점을 MST에 추가
		visited[u] = true
		totalWeight += item.weight

		if item.from != -1 {
			mstEdges = append(mstEdges, [3]int{item.from, u, item.weight})
		}

		// 인접 정점들을 큐에 추가
		for _, e := range adj[u] {
			if !visited[e.to] {
				heap.Push(pq, &Item{node: e.to, weight: e.weight, from: u})
			}
		}
	}

	return totalWeight, mstEdges
}

func main() {
	// 예제 그래프: 5개 정점, 7개 간선 (무방향)
	n := 5
	adj := make([][]Edge, n)
	for i := range adj {
		adj[i] = []Edge{}
	}

	// 간선 추가 (무방향이므로 양쪽 모두 추가)
	addEdge := func(u, v, w int) {
		adj[u] = append(adj[u], Edge{v, w})
		adj[v] = append(adj[v], Edge{u, w})
	}

	addEdge(0, 1, 1)
	addEdge(0, 2, 4)
	addEdge(1, 2, 2)
	addEdge(1, 3, 3)
	addEdge(2, 3, 2)
	addEdge(3, 4, 5)
	addEdge(2, 4, 6)

	fmt.Println("=== 프림 알고리즘 ===")
	fmt.Printf("정점 수: %d\n\n", n)

	totalWeight, mstEdges := prim(n, adj)

	fmt.Println("MST에 포함된 간선:")
	for _, e := range mstEdges {
		fmt.Printf("  %d -- %d (가중치: %d)\n", e[0], e[1], e[2])
	}
	fmt.Printf("\nMST 총 가중치: %d\n", totalWeight)
}
