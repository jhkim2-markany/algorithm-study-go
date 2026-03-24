package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// RaptureItem은 우선순위 큐의 원소이다.
type RaptureItem struct {
	node, maxWeight int
}

// RapturePQ는 최소 힙 기반 우선순위 큐이다.
type RapturePQ []RaptureItem

func (pq RapturePQ) Len() int            { return len(pq) }
func (pq RapturePQ) Less(i, j int) bool  { return pq[i].maxWeight < pq[j].maxWeight }
func (pq RapturePQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *RapturePQ) Push(x interface{}) { *pq = append(*pq, x.(RaptureItem)) }
func (pq *RapturePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// jackGoesToRapture는 1번에서 N번까지 경로 상 최대 요금의 최솟값을 반환한다.
//
// [매개변수]
//   - n: 역 수
//   - edges: 노선 목록 (각 원소는 [3]int{u, v, w})
//
// [반환값]
//   - int: 최대 요금의 최솟값 (-1이면 도달 불가)
func jackGoesToRapture(n int, edges [][3]int) int {
	// 여기에 코드를 작성하세요
	return -1
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

	result := jackGoesToRapture(n, edges)
	if result == -1 {
		fmt.Fprintln(writer, "NO PATH EXISTS")
	} else {
		fmt.Fprintln(writer, result)
	}

	// heap 패키지 사용을 위한 임포트 유지
	_ = heap.Init
}
