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
//
// [알고리즘 힌트]
//   1. 모든 점의 x좌표를 수집하여 정렬 후 중복 제거(좌표 압축)한다.
//   2. 압축된 인덱스에 가중치를 합산한다.
//   3. 누적합 배열을 구성한다.
//   4. 각 쿼리에 대해 이진 탐색으로 구간 경계를 찾아 O(1)에 응답한다.
func compressRangeSum(points []Point, queries []Query) []int64 {
	// 좌표 압축
	allX := make([]int, len(points))
	for i, p := range points {
		allX[i] = p.x
	}
	sort.Ints(allX)

	unique := []int{allX[0]}
	for i := 1; i < len(allX); i++ {
		if allX[i] != allX[i-1] {
			unique = append(unique, allX[i])
		}
	}
	sz := len(unique)

	// 압축된 좌표에 가중치 합산
	weightSum := make([]int64, sz)
	for _, p := range points {
		idx := sort.SearchInts(unique, p.x)
		weightSum[idx] += int64(p.w)
	}

	// 누적합 배열 구성
	prefix := make([]int64, sz+1)
	for i := 0; i < sz; i++ {
		prefix[i+1] = prefix[i] + weightSum[i]
	}

	// 각 쿼리 처리
	results := make([]int64, len(queries))
	for i, qr := range queries {
		lo := sort.SearchInts(unique, qr.l)
		hi := sort.SearchInts(unique, qr.r+1)
		results[i] = prefix[hi] - prefix[lo]
	}
	return results
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
}
