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
func primMST(n int, edges [][3]int, s int) int {
	// 여기에 코드를 작성하세요
	return 0
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

	// heap 패키지 사용을 위한 임포트 유지
	_ = heap.Init
}
