package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// MaxH는 최대 힙을 구현한다.
type MaxH []int

func (h MaxH) Len() int           { return len(h) }
func (h MaxH) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxH) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinH는 최소 힙을 구현한다.
type MinH []int

func (h MinH) Len() int           { return len(h) }
func (h MinH) Less(i, j int) bool { return h[i] < h[j] }
func (h MinH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinH) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// medianUpdates는 각 연산 후의 중앙값을 반환한다.
//
// [매개변수]
//   - ops: 연산 목록 (각 연산은 [타입, 값])
//
// [반환값]
//   - []string: 각 연산 후의 중앙값 문자열
func medianUpdates(ops [][]interface{}) []string {
	// 여기에 코드를 작성하세요
	_ = heap.Init
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 연산 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 연산 입력
	ops := make([][]interface{}, n)
	for i := 0; i < n; i++ {
		var op string
		var x int
		fmt.Fscan(reader, &op, &x)
		ops[i] = []interface{}{op, x}
	}

	// 핵심 함수 호출
	results := medianUpdates(ops)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
