package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Point는 좌표와 가중치를 나타낸다.
type Point struct {
	x, w int
}

// Query는 구간 쿼리의 범위를 나타낸다.
type Query struct {
	l, r int
}

// compressRangeSum은 좌표 압축과 누적합을 이용하여 각 쿼리의 구간 가중치 합을 반환한다.
//
// [매개변수]
//   - points: 좌표와 가중치 배열
//   - queries: 구간 쿼리 배열 (각 쿼리는 [l, r] 범위)
//
// [반환값]
//   - []int64: 각 쿼리에 대한 구간 가중치 합 배열
func compressRangeSum(points []Point, queries []Query) []int64 {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].x, &points[i].w)
	}

	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].l, &queries[i].r)
	}

	results := compressRangeSum(points, queries)
	for _, ans := range results {
		fmt.Fprintln(writer, ans)
	}

	_ = sort.SearchInts // 패키지 사용 보장
}
