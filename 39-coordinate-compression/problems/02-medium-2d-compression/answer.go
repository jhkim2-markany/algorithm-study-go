package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Rect는 직사각형의 좌표를 나타낸다.
type Rect struct {
	x1, y1, x2, y2 int
}

// uniqueSorted는 배열을 정렬하고 중복을 제거한다.
func uniqueSorted(arr []int) []int {
	sort.Ints(arr)
	result := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result
}

// calcUnionArea는 2차원 좌표 압축을 이용하여 직사각형들의 합집합 넓이를 구한다.
//
// [매개변수]
//   - rects: 직사각형 배열 (각 직사각형은 왼쪽 아래, 오른쪽 위 좌표)
//
// [반환값]
//   - int: 직사각형 합집합의 넓이
//
// [알고리즘 힌트]
//   1. 모든 직사각형의 x, y 좌표를 수집하여 각각 좌표 압축한다.
//   2. 압축된 격자의 각 셀에 대해 어떤 직사각형이 덮는지 확인한다.
//   3. 덮이는 셀의 실제 넓이(압축 전 좌표 차이)를 합산한다.
func calcUnionArea(rects []Rect) int {
	xs := []int{}
	ys := []int{}
	for _, r := range rects {
		xs = append(xs, r.x1, r.x2)
		ys = append(ys, r.y1, r.y2)
	}

	xs = uniqueSorted(xs)
	ys = uniqueSorted(ys)

	totalArea := 0
	for i := 0; i < len(xs)-1; i++ {
		for j := 0; j < len(ys)-1; j++ {
			mx := xs[i]
			my := ys[j]

			for _, r := range rects {
				if r.x1 <= mx && mx < r.x2 && r.y1 <= my && my < r.y2 {
					width := xs[i+1] - xs[i]
					height := ys[j+1] - ys[j]
					totalArea += width * height
					break
				}
			}
		}
	}

	return totalArea
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	rects := make([]Rect, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &rects[i].x1, &rects[i].y1, &rects[i].x2, &rects[i].y2)
	}

	fmt.Fprintln(writer, calcUnionArea(rects))
}
