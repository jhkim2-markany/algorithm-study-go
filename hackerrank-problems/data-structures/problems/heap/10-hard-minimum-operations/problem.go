package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// ColorCount는 특정 상자의 특정 색상 공 개수를 나타낸다.
type ColorCount struct {
	count int
	box   int
}

// ColorHeap은 공 개수 기준 최대 힙을 구현한다.
type ColorHeap []ColorCount

func (h ColorHeap) Len() int            { return len(h) }
func (h ColorHeap) Less(i, j int) bool  { return h[i].count > h[j].count }
func (h ColorHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ColorHeap) Push(x interface{}) { *h = append(*h, x.(ColorCount)) }
func (h *ColorHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// minOperations는 공을 색상별로 분리하는 데 필요한 최소 연산 횟수를 반환한다.
//
// [매개변수]
//   - boxes: 각 상자의 [빨강, 초록, 파랑] 공 개수
//
// [반환값]
//   - int64: 최소 연산 횟수
func minOperations(boxes [][]int) int64 {
	// 여기에 코드를 작성하세요
	_ = heap.Init
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 상자 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 각 상자의 공 개수 입력
	boxes := make([][]int, n)
	for i := 0; i < n; i++ {
		var r, g, b int
		fmt.Fscan(reader, &r, &g, &b)
		boxes[i] = []int{r, g, b}
	}

	// 핵심 함수 호출
	result := minOperations(boxes)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
