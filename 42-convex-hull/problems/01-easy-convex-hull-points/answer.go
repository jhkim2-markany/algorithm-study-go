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

// ccw는 세 점의 방향을 판별한다 (외적).
func ccw(a, b, c Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

// convexHullCount는 주어진 점들의 볼록 껍질 꼭짓점 개수를 반환한다.
//
// [매개변수]
//   - points: 2차원 좌표 배열
//
// [반환값]
//   - int: 볼록 껍질의 꼭짓점 개수
//
// [알고리즘 힌트]
//   1. 점들을 x좌표 기준으로 정렬한다 (같으면 y좌표 기준).
//   2. Andrew's Monotone Chain으로 하부 껍질과 상부 껍질을 각각 구성한다.
//   3. 하부/상부 껍질을 합쳐 볼록 껍질을 완성하고 꼭짓점 수를 반환한다.
func convexHullCount(points []Point) int {
	n := len(points)
	if n < 3 {
		return n
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X != points[j].X {
			return points[i].X < points[j].X
		}
		return points[i].Y < points[j].Y
	})

	// 하부 껍질
	lower := []Point{}
	for _, p := range points {
		for len(lower) >= 2 && ccw(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	// 상부 껍질
	upper := []Point{}
	for i := n - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && ccw(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	// 합치기 (끝점 중복 제거)
	hull := lower[:len(lower)-1]
	hull = append(hull, upper[:len(upper)-1]...)
	return len(hull)
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
}
