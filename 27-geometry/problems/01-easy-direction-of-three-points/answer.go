package main

import (
	"bufio"
	"fmt"
	"os"
)

// ccwDirection은 세 점의 방향 관계를 판별한다.
//
// [매개변수]
//   - x1, y1: 첫 번째 점의 좌표
//   - x2, y2: 두 번째 점의 좌표
//   - x3, y3: 세 번째 점의 좌표
//
// [반환값]
//   - int: 반시계 방향이면 1, 시계 방향이면 -1, 일직선이면 0
//
// [알고리즘 힌트]
//
//	외적(Cross Product)을 이용한 CCW 판별.
//	벡터 P1→P2 = (x2-x1, y2-y1), 벡터 P1→P3 = (x3-x1, y3-y1)
//	외적 = (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
//	양수면 반시계, 음수면 시계, 0이면 일직선이다.
func ccwDirection(x1, y1, x2, y2, x3, y3 int) int {
	cross := (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)

	if cross > 0 {
		return 1
	} else if cross < 0 {
		return -1
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var x1, y1, x2, y2, x3, y3 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x3, &y3)

		// 핵심 함수 호출
		fmt.Fprintln(writer, ccwDirection(x1, y1, x2, y2, x3, y3))
	}
}
