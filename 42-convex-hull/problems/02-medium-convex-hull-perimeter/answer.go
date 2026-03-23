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

// ccw는 세 점의 방향을 판별한다 (외적).
func ccw(a, b, c Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

// dist는 두 점 사이의 거리를 계산한다.
func dist(a, b Point) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// convexHullPerimeter는 주어진 점들의 볼록 껍질 둘레를 계산한다.
//
// [매개변수]
//   - points: 2차원 좌표 배열
//
// [반환값]
//   - float64: 볼록 껍질의 둘레 길이
//
// [알고리즘 힌트]
//   1. Andrew's Monotone Chain으로 볼록 껍질을 구한다.
//   2. 볼록 껍질의 인접한 꼭짓점 사이 거리를 모두 합산한다.
func convexHullPerimeter(points []Point) float64 {
	n := len(points)
	if n < 3 {
		if n == 2 {
			return 2 * dist(points[0], points[1])
		}
		return 0.0
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X != points[j].X {
			return points[i].X < points[j].X
		}
		return points[i].Y < points[j].Y
	})

	lower := []Point{}
	for _, p := range points {
		for len(lower) >= 2 && ccw(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	upper := []Point{}
	for i := n - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && ccw(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	hull := lower[:len(lower)-1]
	hull = append(hull, upper[:len(upper)-1]...)

	perimeter := 0.0
	for i := 0; i < len(hull); i++ {
		j := (i + 1) % len(hull)
		perimeter += dist(hull[i], hull[j])
	}

	return perimeter
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
}
