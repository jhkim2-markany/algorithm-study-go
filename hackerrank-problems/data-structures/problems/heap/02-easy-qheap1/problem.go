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
func processQueries(queries [][]int) []int {
	// 여기에 코드를 작성하세요
	_ = heap.Init
	return nil
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
