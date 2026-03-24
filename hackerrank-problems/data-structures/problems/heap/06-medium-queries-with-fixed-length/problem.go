package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Item은 힙에 저장되는 원소로, 값과 인덱스를 가진다.
type Item struct {
	value int
	index int
}

// MaxItemHeap은 값 기준 최대 힙을 구현한다.
type MaxItemHeap []Item

func (h MaxItemHeap) Len() int            { return len(h) }
func (h MaxItemHeap) Less(i, j int) bool  { return h[i].value > h[j].value }
func (h MaxItemHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxItemHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MaxItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// solve는 각 쿼리에 대해 길이 d인 슬라이딩 윈도우 최댓값의 최솟값을 반환한다.
//
// [매개변수]
//   - arr: 정수 수열
//   - queries: 각 쿼리의 d 값
//
// [반환값]
//   - []int: 각 쿼리에 대한 답
func solve(arr []int, queries []int) []int {
	// 여기에 코드를 작성하세요
	_ = heap.Init
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 수열 길이, 쿼리 개수 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 수열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 쿼리 입력
	queries := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i])
	}

	// 핵심 함수 호출
	results := solve(arr, queries)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
