package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// 최대 힙: 중앙값 이하의 값들을 저장한다 (왼쪽 절반)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 최소 힙: 중앙값 초과의 값들을 저장한다 (오른쪽 절반)
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
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

	var n int
	fmt.Fscan(reader, &n)

	// 최대 힙(왼쪽)과 최소 힙(오른쪽)을 사용하여 중앙값을 유지한다.
	// 규칙: maxH의 크기 >= minH의 크기 (차이는 최대 1)
	// 중앙값은 항상 maxH의 루트(최댓값)이다.
	maxH := &MaxHeap{} // 작은 쪽 절반
	minH := &MinHeap{} // 큰 쪽 절반
	heap.Init(maxH)
	heap.Init(minH)

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)

		// 새 값을 적절한 힙에 삽입
		if maxH.Len() == 0 || x <= (*maxH)[0] {
			heap.Push(maxH, x)
		} else {
			heap.Push(minH, x)
		}

		// 크기 균형 조정: maxH의 크기가 minH보다 1 크거나 같도록 유지
		if maxH.Len() > minH.Len()+1 {
			// maxH에서 minH로 이동
			heap.Push(minH, heap.Pop(maxH))
		} else if minH.Len() > maxH.Len() {
			// minH에서 maxH로 이동
			heap.Push(maxH, heap.Pop(minH))
		}

		// 중앙값 출력: maxH의 루트가 중앙값
		fmt.Fprintln(writer, (*maxH)[0])
	}
}
