package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	// 스택 기반으로 히스토그램에서 가장 큰 직사각형 넓이를 구한다
	// 스택에는 막대의 인덱스를 저장하며, 높이가 오름차순을 유지하도록 한다
	stack := []int{}
	var maxArea int64

	for i := 0; i <= n; i++ {
		// 현재 높이: 마지막에 0을 추가하여 스택에 남은 모든 막대를 처리한다
		var curHeight int64
		if i < n {
			curHeight = heights[i]
		}

		// 스택 맨 위 막대의 높이가 현재 높이보다 크면
		// 해당 막대를 높이로 하는 직사각형의 넓이를 계산한다
		for len(stack) > 0 && heights[stack[len(stack)-1]] > curHeight {
			// 스택에서 Pop하여 직사각형의 높이를 결정
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[top]

			// 너비 계산: 스택이 비어 있으면 왼쪽 끝(0)부터 현재 위치(i)까지
			var width int64
			if len(stack) == 0 {
				width = int64(i)
			} else {
				width = int64(i - stack[len(stack)-1] - 1)
			}

			// 최대 넓이 갱신
			area := h * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	// 결과 출력
	fmt.Fprintln(writer, maxArea)
}
