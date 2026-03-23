package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Point는 2차원 좌표를 나타낸다.
type Point struct {
	X, Y int
}

// convexHullCount는 주어진 점들의 볼록 껍질 꼭짓점 개수를 반환한다.
//
// [매개변수]
//   - points: 2차원 좌표 배열
//
// [반환값]
//   - int: 볼록 껍질의 꼭짓점 개수
func convexHullCount(points []Point) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].X, &points[i].Y)
	}

	fmt.Fprintln(writer, convexHullCount(points))

	_ = sort.Slice // 패키지 사용 보장
}
