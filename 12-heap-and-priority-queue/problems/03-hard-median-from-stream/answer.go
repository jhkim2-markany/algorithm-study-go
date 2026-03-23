package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// MaxHeap은 정수 최대 힙이다 (중앙값 이하의 값들을 저장).
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

// MinHeap은 정수 최소 힙이다 (중앙값 초과의 값들을 저장).
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

// findMedians는 수열이 하나씩 주어질 때마다 현재까지의 중앙값을 반환한다.
//
// [매개변수]
//   - nums: N개의 정수가 순서대로 주어지는 배열
//
// [반환값]
//   - []int: 각 수가 추가될 때마다의 중앙값 배열 (길이 N)
//
// [알고리즘 힌트]
//
//	최대 힙(왼쪽 절반)과 최소 힙(오른쪽 절반)을 사용한다.
//	새 값이 최대 힙의 루트 이하이면 최대 힙에, 아니면 최소 힙에 삽입한다.
//	두 힙의 크기 차이가 1 이하가 되도록 균형을 맞추면,
//	최대 힙의 루트가 항상 중앙값이 된다.
func findMedians(nums []int) []int {
	maxH := &MaxHeap{}
	minH := &MinHeap{}
	heap.Init(maxH)
	heap.Init(minH)

	result := make([]int, len(nums))
	for i, x := range nums {
		if maxH.Len() == 0 || x <= (*maxH)[0] {
			heap.Push(maxH, x)
		} else {
			heap.Push(minH, x)
		}

		if maxH.Len() > minH.Len()+1 {
			heap.Push(minH, heap.Pop(maxH))
		} else if minH.Len() > maxH.Len() {
			heap.Push(maxH, heap.Pop(minH))
		}

		result[i] = (*maxH)[0]
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	// 핵심 함수 호출
	medians := findMedians(nums)

	// 결과 출력
	for _, v := range medians {
		fmt.Fprintln(writer, v)
	}
}
