package main

import (
	"bufio"
	"container/heap"
	"fmt"
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
func dijkstra(n int, edges [][3]int, s int) []int {
	// 여기에 코드를 작성하세요
	return nil
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

	// heap 패키지 사용을 위한 임포트 유지
	_ = heap.Init
}
