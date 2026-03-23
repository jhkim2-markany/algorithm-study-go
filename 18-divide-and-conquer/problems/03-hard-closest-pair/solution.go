package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Point 구조체는 2차원 평면 위의 점을 나타낸다
type Point struct {
	x, y int
}

var points []Point

// dist 함수는 두 점 사이의 유클리드 거리를 계산한다
func dist(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}

// closestPair 함수는 분할 정복으로 최근접 점 쌍의 거리를 구한다
func closestPair(lo, hi int) float64 {
	// 기저 조건: 점이 3개 이하이면 브루트포스로 계산
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
		// 점이 2~3개인 경우 y좌표 기준으로 정렬해둔다
		sort.Slice(points[lo:hi], func(a, b int) bool {
			return points[lo+a].y < points[lo+b].y
		})
		return minDist
	}

	// 분할: 중간 지점을 기준으로 나눈다
	mid := (lo + hi) / 2
	midX := points[mid].x

	// 정복: 왼쪽과 오른쪽에서 최근접 거리를 구한다
	dLeft := closestPair(lo, mid)
	dRight := closestPair(mid, hi)

	// 두 결과 중 작은 값
	d := math.Min(dLeft, dRight)

	// 결합: y좌표 기준으로 병합 정렬 (두 정렬된 부분을 합친다)
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

	// 중간선에서 d 이내에 있는 점들만 모은다
	strip := []Point{}
	for i := lo; i < hi; i++ {
		if math.Abs(float64(points[i].x-midX)) < d {
			strip = append(strip, points[i])
		}
	}

	// 스트립 내에서 y좌표 차이가 d 미만인 점들만 비교
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)
	points = make([]Point, n)
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

	// 분할 정복으로 최근접 점 쌍의 거리 계산
	result := closestPair(0, n)

	// 결과 출력 (소수점 아래 6자리)
	fmt.Fprintf(writer, "%.6f\n", result)
}
