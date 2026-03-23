package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		// 세 점의 좌표 입력
		var x1, y1, x2, y2, x3, y3 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x3, &y3)

		// 외적을 이용한 CCW 판별
		// 벡터 P1→P2 = (x2-x1, y2-y1)
		// 벡터 P1→P3 = (x3-x1, y3-y1)
		// 외적 = (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
		cross := (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)

		// 외적 부호에 따라 방향 판별
		if cross > 0 {
			fmt.Fprintln(writer, 1) // 반시계 방향
		} else if cross < 0 {
			fmt.Fprintln(writer, -1) // 시계 방향
		} else {
			fmt.Fprintln(writer, 0) // 일직선
		}
	}
}
