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

	// 철사 길이 입력
	var l float64
	fmt.Fscan(reader, &l)

	// 넓이 합 함수: x를 정사각형에 사용할 때의 총 넓이
	// 정사각형 넓이 = (x/4)², 원 넓이 = π × ((l-x)/(2π))²
	area := func(x float64) float64 {
		// 정사각형 넓이
		side := x / 4.0
		squareArea := side * side

		// 원 넓이
		r := (l - x) / (2.0 * math.Pi)
		circleArea := math.Pi * r * r

		return squareArea + circleArea
	}

	// 이 함수는 오목 함수(위로 볼록)이므로 삼분 탐색으로 최댓값을 찾는다
	// 실제로 이 문제에서는 넓이 합이 x=0 또는 x=L에서 최대이다
	// 삼분 탐색으로 최솟값을 찾고, 양 끝점과 비교하여 최댓값을 결정한다

	// 넓이 합 함수의 최솟값 위치를 삼분 탐색으로 찾기
	lo, hi := 0.0, l
	for iter := 0; iter < 200; iter++ {
		m1 := lo + (hi-lo)/3
		m2 := hi - (hi-lo)/3

		if area(m1) < area(m2) {
			// 최솟값 쪽으로 이동 (m1이 더 작으므로 최솟값은 왼쪽)
			hi = m2
		} else {
			lo = m1
		}
	}

	// 최솟값 위치
	// minX := (lo + hi) / 2

	// 최댓값은 양 끝점 중 하나에서 발생
	bestX := 0.0
	bestArea := area(0)
	if area(l) > bestArea {
		bestX = l
		bestArea = area(l)
	}

	// 결과 출력
	fmt.Fprintf(writer, "%.6f\n", bestX)
	fmt.Fprintf(writer, "%.6f\n", bestArea)
}
