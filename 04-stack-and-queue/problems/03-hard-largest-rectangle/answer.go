package main

import (
	"bufio"
	"fmt"
	"os"
)

// largestRectangle은 히스토그램에서 가장 큰 직사각형의 넓이를 반환한다.
//
// [매개변수]
//   - heights: N개 막대의 높이 배열
//
// [반환값]
//   - int64: 히스토그램 내부에 그릴 수 있는 가장 큰 직사각형의 넓이
//
// [알고리즘 힌트]
//
//	스택에 막대의 인덱스를 저장하며, 높이가 오름차순을 유지하도록 한다.
//	현재 막대의 높이가 스택 맨 위 막대보다 낮으면,
//	스택에서 Pop하여 해당 막대를 높이로 하는 직사각형의 넓이를 계산한다.
//	너비는 스택이 비어 있으면 왼쪽 끝(0)부터 현재 위치(i)까지,
//	아니면 (i - 스택맨위 - 1)이다.
//	마지막에 높이 0을 추가하여 스택에 남은 모든 막대를 처리한다.
//
//	시간복잡도: O(N), 공간복잡도: O(N)
func largestRectangle(heights []int64) int64 {
	n := len(heights)
	stack := []int{}
	var maxArea int64

	for i := 0; i <= n; i++ {
		var curHeight int64
		if i < n {
			curHeight = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > curHeight {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[top]

			var width int64
			if len(stack) == 0 {
				width = int64(i)
			} else {
				width = int64(i - stack[len(stack)-1] - 1)
			}

			area := h * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 막대 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 막대 높이 입력
	heights := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &heights[i])
	}

	// 핵심 함수 호출
	result := largestRectangle(heights)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
