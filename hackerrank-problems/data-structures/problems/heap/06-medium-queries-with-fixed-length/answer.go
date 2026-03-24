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
//
// [알고리즘 힌트]
//
//	최대 힙을 사용하여 슬라이딩 윈도우의 최댓값을 추적한다.
//	힙의 top이 윈도우 범위 밖이면 지연 삭제(lazy deletion)한다.
//	각 윈도우의 최댓값 중 최솟값을 추적한다.
func solve(arr []int, queries []int) []int {
	n := len(arr)
	results := make([]int, len(queries))

	for qi, d := range queries {
		h := &MaxItemHeap{}
		heap.Init(h)

		minOfMax := int(^uint(0) >> 1) // MaxInt

		// 첫 번째 윈도우의 원소를 힙에 삽입
		for i := 0; i < d; i++ {
			heap.Push(h, Item{value: arr[i], index: i})
		}

		// 첫 번째 윈도우의 최댓값
		curMax := (*h)[0].value
		if curMax < minOfMax {
			minOfMax = curMax
		}

		// 슬라이딩 윈도우 이동
		for i := d; i < n; i++ {
			// 새 원소 삽입
			heap.Push(h, Item{value: arr[i], index: i})

			// 윈도우 범위 밖의 원소를 지연 삭제
			for h.Len() > 0 && (*h)[0].index <= i-d {
				heap.Pop(h)
			}

			// 현재 윈도우의 최댓값
			curMax = (*h)[0].value
			if curMax < minOfMax {
				minOfMax = curMax
			}
		}

		results[qi] = minOfMax
	}

	return results
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
