package main

import (
	"bufio"
	"fmt"
	"os"
)

// ccw 함수는 세 점의 방향 관계를 판별한다
func ccw(x1, y1, x2, y2, x3, y3 int) int {
	// 외적 계산: (P2-P1) × (P3-P1)
	cross := (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
	if cross > 0 {
		return 1 // 반시계 방향
	} else if cross < 0 {
		return -1 // 시계 방향
	}
	return 0 // 일직선
}

// min, max 헬퍼 함수
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// onSegment 함수는 일직선 위의 점이 선분 위에 있는지 확인한다
func onSegment(x1, y1, x2, y2, x3, y3 int) bool {
	// 점 (x3, y3)이 선분 (x1,y1)-(x2,y2) 위에 있는지 확인
	return minInt(x1, x2) <= x3 && x3 <= maxInt(x1, x2) &&
		minInt(y1, y2) <= y3 && y3 <= maxInt(y1, y2)
}

// intersects 함수는 두 선분의 교차 여부를 판정한다
func intersects(x1, y1, x2, y2, x3, y3, x4, y4 int) bool {
	// 선분 AB에 대한 C, D의 방향
	d1 := ccw(x1, y1, x2, y2, x3, y3)
	d2 := ccw(x1, y1, x2, y2, x4, y4)
	// 선분 CD에 대한 A, B의 방향
	d3 := ccw(x3, y3, x4, y4, x1, y1)
	d4 := ccw(x3, y3, x4, y4, x2, y2)

	// 일반적인 교차: 양쪽 모두 다른 방향에 위치
	if d1*d2 < 0 && d3*d4 < 0 {
		return true
	}

	// 일직선 위의 특수 케이스: 선분이 겹치는지 확인
	if d1 == 0 && onSegment(x1, y1, x2, y2, x3, y3) {
		return true
	}
	if d2 == 0 && onSegment(x1, y1, x2, y2, x4, y4) {
		return true
	}
	if d3 == 0 && onSegment(x3, y3, x4, y4, x1, y1) {
		return true
	}
	if d4 == 0 && onSegment(x3, y3, x4, y4, x2, y2) {
		return true
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		// 두 선분의 좌표 입력
		var x1, y1, x2, y2, x3, y3, x4, y4 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4)

		// 교차 여부 판정 후 출력
		if intersects(x1, y1, x2, y2, x3, y3, x4, y4) {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}
