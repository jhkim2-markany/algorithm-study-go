package main

import "fmt"

// CCW (Counter-Clockwise) 판별 - 기하학의 기본 연산
// 세 점의 방향 관계를 외적을 이용하여 판별한다.
// 시간 복잡도: O(1)
// 공간 복잡도: O(1)

// Point 구조체는 2차원 평면 위의 점을 나타낸다
type Point struct {
	X, Y int
}

// ccw 함수는 세 점 P1, P2, P3의 방향 관계를 판별한다.
// 반환값: 양수(반시계), 음수(시계), 0(일직선)
func ccw(p1, p2, p3 Point) int {
	// 벡터 P1→P2와 P1→P3의 외적을 계산한다
	cross := (p2.X-p1.X)*(p3.Y-p1.Y) - (p3.X-p1.X)*(p2.Y-p1.Y)
	if cross > 0 {
		return 1 // 반시계 방향 (CCW)
	} else if cross < 0 {
		return -1 // 시계 방향 (CW)
	}
	return 0 // 일직선 (Collinear)
}

// 방향 결과를 한국어 문자열로 변환한다
func directionName(result int) string {
	switch result {
	case 1:
		return "반시계 방향 (CCW)"
	case -1:
		return "시계 방향 (CW)"
	default:
		return "일직선 (Collinear)"
	}
}

func main() {
	// 예시 1: 반시계 방향
	p1 := Point{0, 0}
	p2 := Point{4, 0}
	p3 := Point{2, 3}
	result := ccw(p1, p2, p3)
	fmt.Printf("P1%v, P2%v, P3%v → %s\n", p1, p2, p3, directionName(result))

	// 예시 2: 시계 방향
	p4 := Point{0, 0}
	p5 := Point{2, 3}
	p6 := Point{4, 0}
	result2 := ccw(p4, p5, p6)
	fmt.Printf("P1%v, P2%v, P3%v → %s\n", p4, p5, p6, directionName(result2))

	// 예시 3: 일직선
	p7 := Point{0, 0}
	p8 := Point{2, 2}
	p9 := Point{4, 4}
	result3 := ccw(p7, p8, p9)
	fmt.Printf("P1%v, P2%v, P3%v → %s\n", p7, p8, p9, directionName(result3))
}
