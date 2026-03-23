package main

import (
	"fmt"
	"math"
	"sort"
)

// 볼록 껍질 - Graham Scan과 Andrew's Monotone Chain
// 시간 복잡도: O(N log N)
// 공간 복잡도: O(N)

// Point - 2차원 좌표를 나타내는 구조체
type Point struct {
	X, Y int
}

// ccw - 세 점의 방향을 판별하는 외적 계산
// 양수: 반시계 방향, 음수: 시계 방향, 0: 일직선
func ccw(a, b, c Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}

// dist2 - 두 점 사이 거리의 제곱
func dist2(a, b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}

// grahamScan - Graham Scan 알고리즘으로 볼록 껍질 구하기
func grahamScan(points []Point) []Point {
	n := len(points)
	if n < 3 {
		return points
	}

	// 기준점 선택: y좌표가 가장 작은 점 (같으면 x좌표가 가장 작은 점)
	pivot := 0
	for i := 1; i < n; i++ {
		if points[i].Y < points[pivot].Y ||
			(points[i].Y == points[pivot].Y && points[i].X < points[pivot].X) {
			pivot = i
		}
	}
	points[0], points[pivot] = points[pivot], points[0]
	base := points[0]

	// 기준점에서의 각도 기준으로 정렬
	sort.Slice(points[1:], func(i, j int) bool {
		a, b := points[i+1], points[j+1]
		c := ccw(base, a, b)
		if c != 0 {
			return c > 0 // 반시계 방향이 먼저
		}
		// 같은 각도면 가까운 점이 먼저
		return dist2(base, a) < dist2(base, b)
	})

	// 스택을 이용하여 볼록 껍질 구성
	stack := []Point{points[0], points[1]}
	for i := 2; i < n; i++ {
		// 좌회전이 아니면 스택에서 제거
		for len(stack) >= 2 && ccw(stack[len(stack)-2], stack[len(stack)-1], points[i]) <= 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, points[i])
	}

	return stack
}

// monotoneChain - Andrew's Monotone Chain 알고리즘으로 볼록 껍질 구하기
func monotoneChain(points []Point) []Point {
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

	// 하부 껍질 구성 (왼쪽 → 오른쪽)
	lower := []Point{}
	for _, p := range points {
		for len(lower) >= 2 && ccw(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	// 상부 껍질 구성 (오른쪽 → 왼쪽)
	upper := []Point{}
	for i := n - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && ccw(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	// 하부와 상부 합치기 (각 끝점은 중복이므로 제거)
	hull := lower[:len(lower)-1]
	hull = append(hull, upper[:len(upper)-1]...)
	return hull
}

func main() {
	// 예제 점 집합
	points := []Point{
		{0, 0}, {1, 1}, {2, 2}, {4, 4},
		{0, 4}, {4, 0}, {2, 1}, {3, 3},
		{1, 3}, {3, 1},
	}

	fmt.Println("=== 입력 점 집합 ===")
	for _, p := range points {
		fmt.Printf("(%d, %d) ", p.X, p.Y)
	}
	fmt.Println()

	// Graham Scan
	pts1 := make([]Point, len(points))
	copy(pts1, points)
	hull1 := grahamScan(pts1)
	fmt.Println("\n=== Graham Scan 결과 ===")
	for _, p := range hull1 {
		fmt.Printf("(%d, %d) ", p.X, p.Y)
	}
	fmt.Println()

	// Andrew's Monotone Chain
	pts2 := make([]Point, len(points))
	copy(pts2, points)
	hull2 := monotoneChain(pts2)
	fmt.Println("\n=== Andrew's Monotone Chain 결과 ===")
	for _, p := range hull2 {
		fmt.Printf("(%d, %d) ", p.X, p.Y)
	}
	fmt.Println()

	// 볼록 껍질 둘레 계산
	perimeter := 0.0
	for i := 0; i < len(hull2); i++ {
		j := (i + 1) % len(hull2)
		perimeter += math.Sqrt(float64(dist2(hull2[i], hull2[j])))
	}
	fmt.Printf("\n볼록 껍질 둘레: %.2f\n", perimeter)

	// 볼록 껍질 넓이 계산 (신발끈 공식)
	area := 0
	for i := 0; i < len(hull2); i++ {
		j := (i + 1) % len(hull2)
		area += hull2[i].X*hull2[j].Y - hull2[j].X*hull2[i].Y
	}
	if area < 0 {
		area = -area
	}
	fmt.Printf("볼록 껍질 넓이: %.1f\n", float64(area)/2.0)
}
