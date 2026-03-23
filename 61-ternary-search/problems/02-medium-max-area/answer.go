package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// maxArea는 길이 l의 철사를 정사각형과 원으로 나눌 때,
// 삼분 탐색으로 넓이 합의 최솟값 위치를 찾고 양 끝점과 비교하여
// 최대 넓이를 달성하는 분할 지점과 그 넓이를 반환한다.
//
// [매개변수]
//   - l: 철사의 전체 길이
//
// [반환값]
//   - float64: 최대 넓이를 달성하는 정사각형 사용 길이 x
//   - float64: 해당 분할에서의 최대 넓이
//
// [알고리즘 힌트]
//
//	정사각형 둘레 x → 넓이 (x/4)², 원 둘레 (l-x) → 넓이 (l-x)²/(4π).
//	총 넓이 f(x) = x²/16 + (l-x)²/(4π)는 볼록 함수이므로
//	삼분 탐색으로 최솟값을 찾을 수 있다.
//	최댓값은 양 끝점(x=0 또는 x=l)에서 발생하므로,
//	세 점(x=0, 최솟값 위치, x=l)을 비교하여 최대를 반환한다.
func maxArea(l float64) (float64, float64) {
	// 넓이 함수: 정사각형(x) + 원(l-x)
	area := func(x float64) float64 {
		sq := x / 4.0
		circlePerimeter := l - x
		return sq*sq + circlePerimeter*circlePerimeter/(4.0*math.Pi)
	}

	// 삼분 탐색으로 최솟값 위치를 찾는다 (볼록 함수)
	lo, hi := 0.0, l
	for i := 0; i < 200; i++ {
		m1 := lo + (hi-lo)/3.0
		m2 := hi - (hi-lo)/3.0
		if area(m1) < area(m2) {
			hi = m2
		} else {
			lo = m1
		}
	}

	// 세 후보 비교: x=0, 최솟값 위치, x=l
	candidates := []float64{0, (lo + hi) / 2.0, l}
	bestX := candidates[0]
	bestArea := area(candidates[0])
	for _, x := range candidates[1:] {
		a := area(x)
		if a > bestArea {
			bestArea = a
			bestX = x
		}
	}

	return bestX, bestArea
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l float64
	fmt.Fscan(reader, &l)

	bestX, bestArea := maxArea(l)
	fmt.Fprintf(writer, "%.6f\n", bestX)
	fmt.Fprintf(writer, "%.6f\n", bestArea)
}
