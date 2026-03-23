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

// 두 점 사이 거리의 제곱
func dist2(a, b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}

// 외적의 절댓값 (평행사변형 넓이)
func cross(o, a, b Point) int {
	v := (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
	if v < 0 {
		return -v
	}
	return v
}

// Andrew's Monotone Chain으로 볼록 껍질 구하기
func convexHull(points []Point) []Point {
	n := len(points)
	if n < 3 {
		return points
	}

	// x좌표 기준 정렬
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

	// 합치기
	hull := lower[:len(lower)-1]
	hull = append(hull, upper[:len(upper)-1]...)
	return hull
}

// 회전하는 캘리퍼스로 가장 먼 두 점의 거리 제곱 구하기
func rotatingCalipers(hull []Point) int {
	n := len(hull)
	if n == 2 {
		return dist2(hull[0], hull[1])
	}

	maxDist := 0

	// 대척점 j를 초기화 (가장 먼 점에서 시작)
	j := 1
	for i := 0; i < n; i++ {
		// 현재 변: hull[i] → hull[(i+1)%n]
		ni := (i + 1) % n
		// j를 회전시켜 현재 변에서 가장 먼 점을 찾음
		for {
			nj := (j + 1) % n
			// 외적 비교: hull[ni]-hull[i] 방향으로 hull[nj]가 hull[j]보다 더 먼지 확인
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

		// 현재 대척점 쌍의 거리 갱신
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

	// 점의 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 점 좌표 입력
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].X, &points[i].Y)
	}

	// 점이 2개인 경우 바로 거리 계산
	if n == 2 {
		fmt.Fprintln(writer, dist2(points[0], points[1]))
		return
	}

	// 볼록 껍질 구하기
	hull := convexHull(points)

	// 회전하는 캘리퍼스로 최대 거리 제곱 구하기
	fmt.Fprintln(writer, rotatingCalipers(hull))
}
