package main

import (
	"bufio"
	"fmt"
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
func closestPairDist(points []Point) float64 {
	// 여기에 코드를 작성하세요
	return 0.0
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
