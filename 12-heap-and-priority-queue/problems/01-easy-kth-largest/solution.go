package main

import (
	"bufio"
	"fmt"
	"os"
)

// MinHeap은 정수 최소 힙이다.
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

// findKthLargest는 N개의 수에서 K번째로 큰 수를 반환한다.
//
// [매개변수]
//   - nums: N개의 정수 배열
//   - k: 찾고자 하는 순위 (K번째로 큰 수)
//
// [반환값]
//   - int: K번째로 큰 수
func findKthLargest(nums []int, k int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N과 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 수열 입력
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	// 핵심 함수 호출
	result := findKthLargest(nums, k)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
