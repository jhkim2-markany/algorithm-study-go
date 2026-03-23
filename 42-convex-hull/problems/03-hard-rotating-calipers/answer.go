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

// dist2는 두 점 사이 거리의 제곱을 반환한다.
func dist2(a, b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}

// cross는 외적의 절댓값(평행사변형 넓이)을 반환한다.
func cross(o, a, b Point) int {
	v := (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
	if v < 0 {
		return -v
	}
	return v
}

// convexHull은 Andrew's Monotone Chain으로 볼록 껍질을 구한다.
func convexHull(points []Point) []Point {
	n := len(points)
	if n < 3 {
		return points
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
	return hull
}

// maxDistSquared는 주어진 점들 중 가장 먼 두 점 사이 거리의 제곱을 반환한다.
//
// [매개변수]
//   - points: 2차원 좌표 배열
//
// [반환값]
//   - int: 가장 먼 두 점 사이 거리의 제곱
//
// [알고리즘 힌트]
//   1. 볼록 껍질을 구한다 (Andrew's Monotone Chain).
//   2. 회전하는 캘리퍼스(Rotating Calipers)로 대척점 쌍을 순회한다.
//   3. 각 변에 대해 가장 먼 대척점을 찾아 최대 거리를 갱신한다.
func maxDistSquared(points []Point) int {
	n := len(points)
	if n == 2 {
		return dist2(points[0], points[1])
	}

	hull := convexHull(points)
	h := len(hull)
	if h == 2 {
		return dist2(hull[0], hull[1])
	}

	maxDist := 0
	j := 1
	for i := 0; i < h; i++ {
		ni := (i + 1) % h
		for {
			nj := (j + 1) % h
			if cross(hull[i], hull[ni], Point{
				hull[i].X + hull[nj].X - hull[j].X,
				hull[i].Y + hull[nj].Y - hull[j].Y,
			}) > cross(hull[i], hull[ni], Point{
				hull[i].X, hull[i].Y,
			}) {
				j = nj
			} else {
				break
			}
		}
		d := dist2(hull[i], hull[j])
		if d > maxDist {
			maxDist = d
		}
	}

	return maxDist
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
}
