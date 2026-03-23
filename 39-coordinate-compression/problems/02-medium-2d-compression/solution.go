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
func calcUnionArea(rects []Rect) int {
	// 여기에 코드를 작성하세요
	return 0
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
