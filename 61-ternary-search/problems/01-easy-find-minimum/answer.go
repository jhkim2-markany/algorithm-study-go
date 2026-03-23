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
//
// [알고리즘 힌트]
//   1. f(x) = ax² + bx + c 를 정의한다
//   2. 구간 [lo, hi]를 삼등분하여 m1, m2를 구한다
//   3. f(m1) < f(m2)이면 hi = m2, 아니면 lo = m1로 구간을 좁힌다
//   4. 충분한 반복(약 200회) 후 구간 중점을 최솟값 위치로 반환한다
func findMinimum(a, b, c int, lo, hi float64) (float64, float64) {
	f := func(x float64) float64 {
		return float64(a)*x*x + float64(b)*x + float64(c)
	}

	for i := 0; i < 200; i++ {
		m1 := lo + (hi-lo)/3.0
		m2 := hi - (hi-lo)/3.0
		if f(m1) < f(m2) {
			hi = m2
		} else {
			lo = m1
		}
	}

	x := (lo + hi) / 2.0
	return x, f(x)
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
