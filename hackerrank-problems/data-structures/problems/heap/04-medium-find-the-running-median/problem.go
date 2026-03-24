package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// MaxHeap은 최대 힙을 구현한다.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeap은 최소 힙을 구현한다.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// runningMedian은 각 원소 삽입 후의 중앙값을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - []float64: 각 단계의 중앙값
func runningMedian(arr []int) []float64 {
	// 여기에 코드를 작성하세요
	_ = heap.Init
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정수 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 정수 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	results := runningMedian(arr)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintf(writer, "%.1f\n", r)
	}
}
