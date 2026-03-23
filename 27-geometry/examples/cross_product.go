package main

import (
	"fmt"
	"math"
)

// 외적 (Cross Product) - 기하학의 핵심 연산
// 두 벡터의 외적을 계산하고 다양한 활용 예시를 보여준다.
// 시간 복잡도: O(1)
// 공간 복잡도: O(1)

// Vector 구조체는 2차원 벡터를 나타낸다
type Vector struct {
	X, Y float64
}

// crossProduct 함수는 두 벡터의 외적을 계산한다
// 외적 = A.X * B.Y - A.Y * B.X
func crossProduct(a, b Vector) float64 {
	return a.X*b.Y - a.Y*b.X
}

// vectorFromPoints 함수는 두 점으로부터 벡터를 생성한다
func vectorFromPoints(x1, y1, x2, y2 float64) Vector {
	return Vector{x2 - x1, y2 - y1}
}

// triangleArea 함수는 세 점으로 이루어진 삼각형의 넓이를 계산한다
// 외적의 절댓값 / 2 = 삼각형 넓이
func triangleArea(x1, y1, x2, y2, x3, y3 float64) float64 {
	// 벡터 P1→P2와 P1→P3의 외적을 이용한다
	v1 := vectorFromPoints(x1, y1, x2, y2)
	v2 := vectorFromPoints(x1, y1, x3, y3)
	return math.Abs(crossProduct(v1, v2)) / 2.0
}

// parallelogramArea 함수는 두 벡터로 이루어진 평행사변형의 넓이를 계산한다
// 외적의 절댓값 = 평행사변형 넓이
func parallelogramArea(a, b Vector) float64 {
	return math.Abs(crossProduct(a, b))
}

func main() {
	// 예시 1: 기본 외적 계산
	a := Vector{3, 0}
	b := Vector{0, 4}
	cross := crossProduct(a, b)
	fmt.Printf("벡터 A(%.0f, %.0f) × 벡터 B(%.0f, %.0f) = %.0f\n", a.X, a.Y, b.X, b.Y, cross)

	// 예시 2: 외적 부호로 방향 판별
	v1 := Vector{1, 0}
	v2 := Vector{1, 1}
	cross2 := crossProduct(v1, v2)
	fmt.Printf("\n벡터 (%.0f,%.0f) × (%.0f,%.0f) = %.0f", v1.X, v1.Y, v2.X, v2.Y, cross2)
	if cross2 > 0 {
		fmt.Println(" → 반시계 방향")
	} else if cross2 < 0 {
		fmt.Println(" → 시계 방향")
	} else {
		fmt.Println(" → 평행")
	}

	// 예시 3: 삼각형 넓이 계산
	area := triangleArea(0, 0, 4, 0, 0, 3)
	fmt.Printf("\n삼각형 (0,0)-(4,0)-(0,3)의 넓이: %.1f\n", area)

	// 예시 4: 평행사변형 넓이 계산
	pArea := parallelogramArea(Vector{3, 0}, Vector{1, 4})
	fmt.Printf("평행사변형 넓이 (벡터 (3,0)과 (1,4)): %.1f\n", pArea)
}
