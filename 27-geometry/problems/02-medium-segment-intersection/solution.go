package main

import (
	"bufio"
	"fmt"
	"os"
)

// ccw는 세 점의 방향 관계를 판별한다
func ccw(x1, y1, x2, y2, x3, y3 int) int {
	cross := (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
	if cross > 0 {
		return 1
	} else if cross < 0 {
		return -1
	}
	return 0
}

// minInt는 두 정수 중 작은 값을 반환한다
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// maxInt는 두 정수 중 큰 값을 반환한다
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// onSegment는 일직선 위의 점이 선분 위에 있는지 확인한다
func onSegment(x1, y1, x2, y2, x3, y3 int) bool {
	return minInt(x1, x2) <= x3 && x3 <= maxInt(x1, x2) &&
		minInt(y1, y2) <= y3 && y3 <= maxInt(y1, y2)
}

// intersects는 두 선분의 교차 여부를 판정한다.
//
// [매개변수]
//   - x1, y1, x2, y2: 첫 번째 선분의 양 끝점 좌표
//   - x3, y3, x4, y4: 두 번째 선분의 양 끝점 좌표
//
// [반환값]
//   - bool: 두 선분이 교차하면 true, 아니면 false
func intersects(x1, y1, x2, y2, x3, y3, x4, y4 int) bool {
	// 여기에 코드를 작성하세요
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var x1, y1, x2, y2, x3, y3, x4, y4 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4)

		// 핵심 함수 호출
		if intersects(x1, y1, x2, y2, x3, y3, x4, y4) {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}
