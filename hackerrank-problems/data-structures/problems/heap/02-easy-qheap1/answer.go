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

// processQueries는 힙 쿼리를 처리한다.
//
// [매개변수]
//   - queries: 쿼리 목록 (각 쿼리는 [타입, 값] 또는 [타입])
//
// [반환값]
//   - []int: 타입 3 쿼리에 대한 결과 목록
//
// [알고리즘 힌트]
//
//	최소 힙과 lazy deletion을 조합한다.
//	삭제할 원소를 맵에 기록해두고, 최솟값 조회 시 삭제된 원소를 건너뛴다.
func processQueries(queries [][]int) []int {
	h := &IntMinHeap{}
	heap.Init(h)

	// 삭제 표시용 맵
	deleted := make(map[int]int)
	var results []int

	for _, query := range queries {
		switch query[0] {
		case 1:
			// 원소 추가
			heap.Push(h, query[1])
		case 2:
			// lazy deletion: 삭제 표시
			deleted[query[1]]++
		case 3:
			// 삭제된 원소를 건너뛰고 최솟값 조회
			for h.Len() > 0 && deleted[(*h)[0]] > 0 {
				val := heap.Pop(h).(int)
				deleted[val]--
				if deleted[val] == 0 {
					delete(deleted, val)
				}
			}
			if h.Len() > 0 {
				results = append(results, (*h)[0])
			}
		}
	}

	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 쿼리 개수 입력
	var q int
	fmt.Fscan(reader, &q)

	// 쿼리 입력
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 3 {
			queries[i] = []int{t}
		} else {
			var v int
			fmt.Fscan(reader, &v)
			queries[i] = []int{t, v}
		}
	}

	// 핵심 함수 호출
	results := processQueries(queries)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
