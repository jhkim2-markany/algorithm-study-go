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

// maxDistSquared는 주어진 점들 중 가장 먼 두 점 사이 거리의 제곱을 반환한다.
//
// [매개변수]
//   - points: 2차원 좌표 배열
//
// [반환값]
//   - int: 가장 먼 두 점 사이 거리의 제곱
func maxDistSquared(points []Point) int {
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

	fmt.Fprintln(writer, maxDistSquared(points))

	_ = sort.Slice // 패키지 사용 보장
}
