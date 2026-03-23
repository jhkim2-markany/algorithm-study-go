package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Item은 힙에 저장되는 원소로, 값과 출처 리스트 정보를 포함한다.
type Item struct {
	value   int // 원소 값
	listIdx int // 출처 리스트 번호
	elemIdx int // 해당 리스트 내 인덱스
}

// ItemHeap은 Item의 최소 힙이다.
type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 리스트 개수 입력
	var k int
	fmt.Fscan(reader, &k)

	// K개의 리스트 입력
	lists := make([][]int, k)
	for i := 0; i < k; i++ {
		var m int
		fmt.Fscan(reader, &m)
		lists[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &lists[i][j])
		}
	}

	// 각 리스트의 첫 번째 원소를 힙에 삽입
	h := &ItemHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		if len(lists[i]) > 0 {
			heap.Push(h, Item{lists[i][0], i, 0})
		}
	}

	// 힙에서 최솟값을 꺼내고, 해당 리스트의 다음 원소를 삽입
	first := true
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)

		if !first {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, item.value)
		first = false

		// 같은 리스트의 다음 원소가 있으면 힙에 추가
		nextIdx := item.elemIdx + 1
		if nextIdx < len(lists[item.listIdx]) {
			heap.Push(h, Item{
				value:   lists[item.listIdx][nextIdx],
				listIdx: item.listIdx,
				elemIdx: nextIdx,
			})
		}
	}
	fmt.Fprintln(writer)
}
