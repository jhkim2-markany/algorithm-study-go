package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// minMaxDistance는 x축 위의 점 P(t, 0)에서 주어진 점들까지의 최대 거리를
// 최소화하는 t와 그때의 최대 거리를 삼분 탐색으로 구한다.
//
// [매개변수]
//   - x, y: 각 점의 x좌표, y좌표 배열
//   - n: 점의 개수
//
// [반환값]
//   - float64: 최대 거리를 최소화하는 t 좌표
//   - float64: 해당 t에서의 최대 거리
//
// [알고리즘 힌트]
//   1. P(t,0)에서 점 (xi,yi)까지 거리: sqrt((t-xi)²+yi²)
//   2. 각 점까지의 거리는 t에 대해 볼록 함수이다
//   3. 볼록 함수들의 max도 볼록 → 삼분 탐색으로 최솟값을 찾을 수 있다
//   4. 탐색 구간은 x좌표의 최솟값~최댓값으로 설정한다
func minMaxDistance(x, y []float64, n int) (float64, float64) {
	maxDist := func(t float64) float64 {
		best := 0.0
		for i := 0; i < n; i++ {
			dx := t - x[i]
			d := math.Sqrt(dx*dx + y[i]*y[i])
			if d > best {
				best = d
			}
		}
		return best
	}

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

	for i := 0; i < 200; i++ {
		m1 := lo + (hi-lo)/3.0
		m2 := hi - (hi-lo)/3.0
		if maxDist(m1) < maxDist(m2) {
			hi = m2
		} else {
			lo = m1
		}
	}

	t := (lo + hi) / 2.0
	return t, maxDist(t)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	t, dist := minMaxDistance(x, y, n)
	fmt.Fprintf(writer, "%.6f\n", t)
	fmt.Fprintf(writer, "%.6f\n", dist)
}
