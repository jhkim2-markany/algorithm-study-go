package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Point는 2차원 평면 위의 점을 나타낸다.
type Point struct {
	x, y int
}

// closestPairDist는 분할 정복으로 최근접 점 쌍의 거리를 반환한다.
//
// [매개변수]
//   - points: 2차원 점 배열 (x좌표 기준 정렬 필요)
//
// [반환값]
//   - float64: 최근접 점 쌍 사이의 유클리드 거리
//
// [알고리즘 힌트]
//
//	x좌표 기준으로 정렬한 뒤 분할 정복한다.
//	기저 조건: 점이 3개 이하이면 브루트포스로 계산한다.
//	분할: 중간 지점을 기준으로 왼쪽/오른쪽으로 나눈다.
//	정복: 양쪽의 최근접 거리 d를 구한다.
//	결합: 중간선에서 d 이내의 점들(스트립)에서 y좌표 차이가 d 미만인 쌍만 비교한다.
//	y좌표 기준 병합 정렬을 함께 수행하여 효율을 높인다.
func closestPairDist(points []Point) float64 {
	dist := func(a, b Point) float64 {
		dx := float64(a.x - b.x)
		dy := float64(a.y - b.y)
		return math.Sqrt(dx*dx + dy*dy)
	}

	var solve func(lo, hi int) float64
	solve = func(lo, hi int) float64 {
		if hi-lo < 3 {
			minDist := math.Inf(1)
			for i := lo; i < hi; i++ {
				for j := i + 1; j < hi; j++ {
					d := dist(points[i], points[j])
					if d < minDist {
						minDist = d
					}
				}
			}
			sort.Slice(points[lo:hi], func(a, b int) bool {
				return points[lo+a].y < points[lo+b].y
			})
			return minDist
		}

		mid := (lo + hi) / 2
		midX := points[mid].x

		dLeft := solve(lo, mid)
		dRight := solve(mid, hi)
		d := math.Min(dLeft, dRight)

		// y좌표 기준 병합
		merged := make([]Point, hi-lo)
		li, ri, k := lo, mid, 0
		for li < mid && ri < hi {
			if points[li].y <= points[ri].y {
				merged[k] = points[li]
				li++
			} else {
				merged[k] = points[ri]
				ri++
			}
			k++
		}
		for li < mid {
			merged[k] = points[li]
			li++
			k++
		}
		for ri < hi {
			merged[k] = points[ri]
			ri++
			k++
		}
		copy(points[lo:hi], merged)

		// 스트립 내 비교
		var strip []Point
		for i := lo; i < hi; i++ {
			if math.Abs(float64(points[i].x-midX)) < d {
				strip = append(strip, points[i])
			}
		}
		for i := 0; i < len(strip); i++ {
			for j := i + 1; j < len(strip) && float64(strip[j].y-strip[i].y) < d; j++ {
				dd := dist(strip[i], strip[j])
				if dd < d {
					d = dd
				}
			}
		}

		return d
	}

	return solve(0, len(points))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].x, &points[i].y)
	}

	// x좌표 기준으로 정렬
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	// 핵심 함수 호출
	result := closestPairDist(points)

	// 결과 출력 (소수점 아래 6자리)
	fmt.Fprintf(writer, "%.6f\n", result)
}
