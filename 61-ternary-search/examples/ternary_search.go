package main

import (
	"fmt"
	"math"
)

// 삼분 탐색 (Ternary Search) - 유니모달 함수의 극값을 O(log N)에 탐색
// 시간 복잡도: O(log((hi-lo)/eps))
// 공간 복잡도: O(1)

// ternarySearchMin은 볼록 함수 f의 구간 [lo, hi]에서 최솟값의 x 좌표를 반환한다.
// 실수 도메인에서 동작하며, 반복 횟수를 고정하여 정밀도를 보장한다.
func ternarySearchMin(lo, hi float64, f func(float64) float64) float64 {
	// 충분한 반복 횟수로 정밀도 보장 (약 10^-18 수준)
	for iter := 0; iter < 200; iter++ {
		// 구간을 3등분하는 두 점 계산
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
	return (lo + hi) / 2
}

// ternarySearchMax는 오목 함수 f의 구간 [lo, hi]에서 최댓값의 x 좌표를 반환한다.
func ternarySearchMax(lo, hi float64, f func(float64) float64) float64 {
	for iter := 0; iter < 200; iter++ {
		m1 := lo + (hi-lo)/3
		m2 := hi - (hi-lo)/3

		if f(m1) > f(m2) {
			// 최댓값은 m2 오른쪽에 없다
			hi = m2
		} else {
			// 최댓값은 m1 왼쪽에 없다
			lo = m1
		}
	}
	return (lo + hi) / 2
}

// ternarySearchIntMin은 정수 도메인에서 볼록 함수의 최솟값 위치를 반환한다.
func ternarySearchIntMin(lo, hi int, f func(int) int64) int {
	for hi-lo > 2 {
		m1 := lo + (hi-lo)/3
		m2 := hi - (hi-lo)/3
		if f(m1) > f(m2) {
			lo = m1
		} else {
			hi = m2
		}
	}
	// 남은 후보를 직접 비교
	best := lo
	for x := lo + 1; x <= hi; x++ {
		if f(x) < f(best) {
			best = x
		}
	}
	return best
}

func main() {
	// 예제 1: 볼록 함수 f(x) = (x-3)² + 1 의 최솟값 찾기
	f1 := func(x float64) float64 {
		return (x-3)*(x-3) + 1
	}

	minX := ternarySearchMin(0, 10, f1)
	fmt.Printf("예제 1: f(x) = (x-3)² + 1\n")
	fmt.Printf("최솟값 위치: x = %.6f\n", minX)
	fmt.Printf("최솟값: f(x) = %.6f\n\n", f1(minX))

	// 예제 2: 오목 함수 f(x) = -2x² + 8x + 3 의 최댓값 찾기
	f2 := func(x float64) float64 {
		return -2*x*x + 8*x + 3
	}

	maxX := ternarySearchMax(-10, 10, f2)
	fmt.Printf("예제 2: f(x) = -2x² + 8x + 3\n")
	fmt.Printf("최댓값 위치: x = %.6f\n", maxX)
	fmt.Printf("최댓값: f(x) = %.6f\n\n", f2(maxX))

	// 예제 3: 정수 도메인에서 최솟값 찾기
	// f(x) = |x - 7| + |x - 2| + |x - 15| (x=7에서 최솟값)
	f3 := func(x int) int64 {
		abs := func(a int) int64 {
			if a < 0 {
				return int64(-a)
			}
			return int64(a)
		}
		return abs(x-7) + abs(x-2) + abs(x-15)
	}

	bestX := ternarySearchIntMin(0, 20, f3)
	fmt.Printf("예제 3: f(x) = |x-7| + |x-2| + |x-15| (정수)\n")
	fmt.Printf("최솟값 위치: x = %d\n", bestX)
	fmt.Printf("최솟값: f(x) = %d\n\n", f3(bestX))

	// 예제 4: 두 점 사이 거리 합 최소화
	// 점 A(1,3), B(5,1)에서 x축 위의 점 P(x,0)까지 거리 합 최소화
	f4 := func(x float64) float64 {
		dA := math.Sqrt((x-1)*(x-1) + 9)
		dB := math.Sqrt((x-5)*(x-5) + 1)
		return dA + dB
	}

	optX := ternarySearchMin(0, 10, f4)
	fmt.Printf("예제 4: A(1,3), B(5,1)에서 x축 위 점까지 거리 합 최소\n")
	fmt.Printf("최적 위치: x = %.6f\n", optX)
	fmt.Printf("최소 거리 합: %.6f\n", f4(optX))
}
