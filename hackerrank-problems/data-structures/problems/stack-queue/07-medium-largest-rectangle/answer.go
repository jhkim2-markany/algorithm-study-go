package main

import (
	"bufio"
	"fmt"
	"os"
)

// largestRectangle는 히스토그램에서 가장 큰 직사각형의 넓이를 반환한다.
//
// [매개변수]
//   - h: 건물 높이 배열
//
// [반환값]
//   - int64: 가장 큰 직사각형의 넓이
//
// [알고리즘 힌트]
//
//	단조 증가 스택을 사용하여 각 막대의 좌우 경계를 찾는다.
//	현재 높이가 스택 최상위보다 작으면 팝하며 넓이를 계산한다.
func largestRectangle(h []int) int64 {
	n := len(h)
	// 인덱스를 저장하는 스택
	stack := []int{}
	var maxArea int64

	for i := 0; i <= n; i++ {
		// 현재 높이 (배열 끝에 도달하면 0으로 처리)
		var curHeight int
		if i < n {
			curHeight = h[i]
		}

		// 스택 최상위의 높이가 현재 높이보다 크면 넓이 계산
		for len(stack) > 0 && h[stack[len(stack)-1]] > curHeight {
			// 스택에서 팝
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 너비 계산: 스택이 비어있으면 왼쪽 경계는 0
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			// 넓이 계산 및 최대값 갱신
			area := int64(h[top]) * int64(width)
			if area > maxArea {
				maxArea = area
			}
		}

		// 현재 인덱스를 스택에 푸시
		stack = append(stack, i)
	}

	return maxArea
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 건물 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 건물 높이 입력
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}

	// 핵심 함수 호출
	result := largestRectangle(h)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
