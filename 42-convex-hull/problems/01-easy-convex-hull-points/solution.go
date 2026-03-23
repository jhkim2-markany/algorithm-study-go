package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 2차원 좌표
type Point struct {
	X, Y int
}

// 세 점의 방향 판별 (외적)
func ccw(a, b, c Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

// Andrew's Monotone Chain으로 볼록 껍질 구하기
func convexHull(points []Point) []Point {
	n := len(points)
	if n < 3 {
		return points
	}

	// x좌표 기준 정렬 (같으면 y좌표 기준)
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
	return hull
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 점의 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 점 좌표 입력
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].X, &points[i].Y)
	}

	// 볼록 껍질 구하기
	hull := convexHull(points)

	// 꼭짓점 개수 출력
	fmt.Fprintln(writer, len(hull))
}
