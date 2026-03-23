package main

import (
	"bufio"
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
func findMedians(nums []int) []int {
	// 여기에 코드를 작성하세요
	return nil
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
