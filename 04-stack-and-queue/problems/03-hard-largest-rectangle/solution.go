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
func largestRectangle(heights []int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
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
