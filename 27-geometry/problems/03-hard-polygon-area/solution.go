package main

import (
	"bufio"
	"fmt"
	"os"
)

// polygonArea는 다각형의 넓이를 구한다.
//
// [매개변수]
//   - x: 꼭짓점의 x 좌표 배열
//   - y: 꼭짓점의 y 좌표 배열
//   - n: 꼭짓점의 수
//
// [반환값]
//   - float64: 다각형의 넓이
func polygonArea(x, y []int64, n int) float64 {
	// 여기에 코드를 작성하세요
	return 0.0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	x := make([]int64, n)
	y := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	// 핵심 함수 호출
	area := polygonArea(x, y, n)

	// 소수점 첫째 자리까지 출력
	fmt.Fprintf(writer, "%.1f\n", area)
}
