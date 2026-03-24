package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// IntMinHeap은 정수 최소 힙을 구현한다.
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// cookies는 모든 쿠키의 달콤함이 K 이상이 되기 위한 최소 연산 횟수를 반환한다.
//
// [매개변수]
//   - k: 목표 달콤함
//   - arr: 쿠키의 달콤함 배열
//
// [반환값]
//   - int: 최소 연산 횟수 (불가능하면 -1)
func cookies(k int, arr []int) int {
	// 여기에 코드를 작성하세요
	_ = heap.Init
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 쿠키 개수 N, 목표 달콤함 K
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 쿠키 달콤함 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := cookies(k, arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
