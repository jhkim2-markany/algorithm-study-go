package main

import (
	"bufio"
	"fmt"
	"os"
)

// findMinimum은 삼분 탐색을 이용하여 볼록 함수 f(x) = ax² + bx + c의
// 구간 [lo, hi]에서 최솟값의 위치와 그 값을 반환한다.
//
// [매개변수]
//   - a, b, c: 이차 함수의 계수
//   - lo, hi: 탐색 구간의 양 끝점
//
// [반환값]
//   - float64: 최솟값이 되는 x 좌표
//   - float64: 해당 x에서의 함수값 f(x)
func findMinimum(a, b, c int, lo, hi float64) (float64, float64) {
	// 여기에 코드를 작성하세요
	return 0, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	var lo, hi float64
	fmt.Fscan(reader, &lo, &hi)

	x, fx := findMinimum(a, b, c, lo, hi)
	fmt.Fprintf(writer, "%.6f\n", x)
	fmt.Fprintf(writer, "%.6f\n", fx)
}
