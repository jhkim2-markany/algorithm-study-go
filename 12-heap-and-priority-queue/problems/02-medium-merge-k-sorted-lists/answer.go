package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Item은 힙에 저장되는 원소로, 값과 출처 리스트 정보를 포함한다.
type Item struct {
	value   int
	listIdx int
	elemIdx int
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

// mergeKSortedLists는 K개의 정렬된 리스트를 하나의 정렬된 배열로 병합한다.
//
// [매개변수]
//   - lists: K개의 정렬된 정수 배열
//
// [반환값]
//   - []int: 병합된 하나의 정렬된 배열
//
// [알고리즘 힌트]
//
//	최소 힙에 각 리스트의 첫 번째 원소를 삽입한다.
//	힙에서 최솟값을 꺼내 결과에 추가하고,
//	해당 리스트의 다음 원소가 있으면 힙에 삽입한다.
//	힙이 빌 때까지 반복하면 전체가 정렬된 순서로 병합된다.
func mergeKSortedLists(lists [][]int) []int {
	h := &ItemHeap{}
	heap.Init(h)

	for i := 0; i < len(lists); i++ {
		if len(lists[i]) > 0 {
			heap.Push(h, Item{lists[i][0], i, 0})
		}
	}

	var result []int
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		result = append(result, item.value)

		nextIdx := item.elemIdx + 1
		if nextIdx < len(lists[item.listIdx]) {
			heap.Push(h, Item{
				value:   lists[item.listIdx][nextIdx],
				listIdx: item.listIdx,
				elemIdx: nextIdx,
			})
		}
	}

	return result
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

	// 핵심 함수 호출
	result := mergeKSortedLists(lists)

	// 결과 출력
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
