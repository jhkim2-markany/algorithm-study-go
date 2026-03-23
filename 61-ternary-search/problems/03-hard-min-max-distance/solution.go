package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 점의 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 각 점의 좌표 입력
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	// f(t) = max(dist(P(t,0), Pi)) 함수 정의
	// 각 점까지의 거리 중 최댓값을 반환
	f := func(t float64) float64 {
		maxDist := 0.0
		for i := 0; i < n; i++ {
			dx := t - x[i]
			dist := math.Sqrt(dx*dx + y[i]*y[i])
			if dist > maxDist {
				maxDist = dist
			}
		}
		return maxDist
	}

	// f(t)는 볼록 함수이다 (각 점까지의 거리 함수는 볼록이고, 볼록 함수의 max도 볼록)
	// 삼분 탐색으로 최솟값을 찾는다

	// 탐색 구간 설정: x 좌표의 최솟값 ~ 최댓값으로 충분
	lo := x[0]
	hi := x[0]
	for i := 1; i < n; i++ {
		if x[i] < lo {
			lo = x[i]
		}
		if x[i] > hi {
			hi = x[i]
		}
	}
	// 여유를 두고 구간 확장
	lo -= 1.0
	hi += 1.0

	// 삼분 탐색 수행
	for iter := 0; iter < 200; iter++ {
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
	t := (lo + hi) / 2
	fmt.Fprintf(writer, "%.6f\n", t)
	fmt.Fprintf(writer, "%.6f\n", f(t))
}
