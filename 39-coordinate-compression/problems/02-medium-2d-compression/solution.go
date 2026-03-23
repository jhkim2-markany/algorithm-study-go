package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 직사각형 구조체
type Rect struct {
	x1, y1, x2, y2 int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 직사각형 수
	var n int
	fmt.Fscan(reader, &n)

	rects := make([]Rect, n)
	xs := []int{} // x 좌표 수집
	ys := []int{} // y 좌표 수집

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &rects[i].x1, &rects[i].y1, &rects[i].x2, &rects[i].y2)
		xs = append(xs, rects[i].x1, rects[i].x2)
		ys = append(ys, rects[i].y1, rects[i].y2)
	}

	// x, y 좌표 각각 정렬 후 중복 제거 (좌표 압축)
	xs = uniqueSorted(xs)
	ys = uniqueSorted(ys)

	// 압축된 격자에서 각 셀이 직사각형에 덮이는지 확인
	totalArea := 0
	for i := 0; i < len(xs)-1; i++ {
		for j := 0; j < len(ys)-1; j++ {
			// 현재 셀의 대표 좌표 (셀 내부의 한 점)
			mx := xs[i]
			my := ys[j]

			// 어떤 직사각형이 이 셀을 덮는지 확인
			for _, r := range rects {
				if r.x1 <= mx && mx < r.x2 && r.y1 <= my && my < r.y2 {
					// 셀의 실제 넓이를 더한다
					width := xs[i+1] - xs[i]
					height := ys[j+1] - ys[j]
					totalArea += width * height
					break // 한 번이라도 덮이면 넓이 추가 후 다음 셀로
				}
			}
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, totalArea)
}

// uniqueSorted는 배열을 정렬하고 중복을 제거한다
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
