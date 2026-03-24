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
//
// [알고리즘 힌트]
//
//	각 상자에서 가장 많은 색상의 공을 남기고 나머지를 이동한다.
//	총 공 수에서 각 상자의 최대 색상 공 수를 빼면 최소 이동 횟수이다.
//	힙을 사용하여 각 색상별로 가장 많은 공을 가진 상자를 우선 처리할 수 있다.
func minOperations(boxes [][]int) int64 {
	// 각 색상별 힙을 구성하여 최적 배정을 찾는다
	// 단순 그리디: 각 상자에서 가장 많은 색상의 공을 남긴다
	var totalOps int64

	// 힙을 사용하여 각 상자의 최대 색상 공 수를 효율적으로 추적
	h := &ColorHeap{}
	heap.Init(h)

	for i, box := range boxes {
		total := box[0] + box[1] + box[2]
		if total == 0 {
			continue
		}
		// 각 상자에서 가장 많은 색상의 공 수를 찾는다
		maxColor := box[0]
		if box[1] > maxColor {
			maxColor = box[1]
		}
		if box[2] > maxColor {
			maxColor = box[2]
		}
		heap.Push(h, ColorCount{count: maxColor, box: i})
		// 이동 횟수 = 총 공 수 - 최대 색상 공 수
		totalOps += int64(total - maxColor)
	}

	return totalOps
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
