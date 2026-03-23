package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Point는 2차원 좌표를 나타낸다.
type Point struct {
	X, Y int
}

// convexHullPerimeter는 주어진 점들의 볼록 껍질 둘레를 계산한다.
//
// [매개변수]
//   - points: 2차원 좌표 배열
//
// [반환값]
//   - float64: 볼록 껍질의 둘레 길이
func convexHullPerimeter(points []Point) float64 {
	// 여기에 코드를 작성하세요
	return 0.0
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

	fmt.Fprintf(writer, "%.2f\n", convexHullPerimeter(points))

	_ = sort.Slice // 패키지 사용 보장
	_ = math.Sqrt  // 패키지 사용 보장
}
