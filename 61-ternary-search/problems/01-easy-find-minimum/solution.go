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

	// 이차 함수 계수 입력
	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	// 탐색 구간 입력
	var lo, hi float64
	fmt.Fscan(reader, &lo, &hi)

	// 볼록 함수 f(x) = ax² + bx + c 정의
	f := func(x float64) float64 {
		return float64(a)*x*x + float64(b)*x + float64(c)
	}

	// 삼분 탐색으로 최솟값 위치 탐색
	// 반복 횟수를 200회로 고정하여 충분한 정밀도 보장
	for iter := 0; iter < 200; iter++ {
		// 구간을 3등분하는 두 점
		m1 := lo + (hi-lo)/3
		m2 := hi - (hi-lo)/3

		if f(m1) < f(m2) {
			// 최솟값은 m2 오른쪽에 없다
			hi = m2
		} else {
			// 최솟값은 m1 왼쪽에 없다
			lo = m1
		}
	}

	// 결과 출력
	x := (lo + hi) / 2
	fmt.Fprintf(writer, "%.6f\n", x)
	fmt.Fprintf(writer, "%.6f\n", f(x))
}
