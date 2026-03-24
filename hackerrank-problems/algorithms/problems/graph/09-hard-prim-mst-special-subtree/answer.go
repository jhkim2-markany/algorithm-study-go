package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// PrimItem은 프림 알고리즘의 우선순위 큐 원소이다.
type PrimItem struct {
	node, weight, sumNodes int
}

// PrimPQ는 최소 힙 기반 우선순위 큐이다.
type PrimPQ []PrimItem

func (pq PrimPQ) Len() int { return len(pq) }
func (pq PrimPQ) Less(i, j int) bool {
	if pq[i].weight != pq[j].weight {
		return pq[i].weight < pq[j].weight
	}
	return pq[i].sumNodes < pq[j].sumNodes
}
func (pq PrimPQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PrimPQ) Push(x interface{}) { *pq = append(*pq, x.(PrimItem)) }
func (pq *PrimPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// primMST는 프림 알고리즘으로 MST의 가중치 합을 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 원소는 [3]int{u, v, w})
//   - s: 시작 노드 (1-indexed)
//
// [반환값]
//   - int: MST의 가중치 합
//
// [알고리즘 힌트]
//
//	시작 노드에서 출발하여 우선순위 큐로 최소 가중치 간선을 선택한다.
//	이미 MST에 포함된 노드는 건너뛴다.
func primMST(n int, edges [][3]int, s int) int {
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

	// MST 포함 여부
	inMST := make([]bool, n+1)
	totalWeight := 0

	// 시작 노드의 인접 간선을 큐에 추가
	pq := &PrimPQ{}
	heap.Init(pq)
	inMST[s] = true
	for _, e := range adj[s] {
		heap.Push(pq, PrimItem{e.to, e.w, s + e.to})
	}

	// 프림 알고리즘 수행
	for pq.Len() > 0 {
		item := heap.Pop(pq).(PrimItem)
		// 이미 MST에 포함된 노드는 건너뛰기
		if inMST[item.node] {
			continue
		}
		// MST에 추가
		inMST[item.node] = true
		totalWeight += item.weight
		// 새 노드의 인접 간선을 큐에 추가
		for _, e := range adj[item.node] {
			if !inMST[e.to] {
				heap.Push(pq, PrimItem{e.to, e.w, item.node + e.to})
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

	edges := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	var s int
	fmt.Fscan(reader, &s)

	result := primMST(n, edges, s)
	fmt.Fprintln(writer, result)
}
