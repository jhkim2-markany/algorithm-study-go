package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// MinHeap은 정수 최소 힙이다.
// K번째 큰 수를 구하기 위해 크기 K의 최소 힙을 유지한다.
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

	// N과 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 크기 K의 최소 힙을 유지하여 K번째 큰 수를 구한다.
	// 힙의 루트가 항상 K번째로 큰 수가 된다.
	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)

		if h.Len() < k {
			// 힙 크기가 K 미만이면 무조건 삽입
			heap.Push(h, x)
		} else if x > (*h)[0] {
			// 현재 값이 힙의 최솟값보다 크면 교체
			heap.Pop(h)
			heap.Push(h, x)
		}
	}

	// 힙의 루트가 K번째로 큰 수
	fmt.Fprintln(writer, (*h)[0])
}
