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
func largestRectangle(h []int) int64 {
	// 여기에 코드를 작성하세요
	return 0
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
