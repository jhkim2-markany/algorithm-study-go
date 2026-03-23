package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 점의 수, 쿼리 수
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 점 정보 입력
	type Point struct {
		x, w int
	}
	points := make([]Point, n)
	allX := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].x, &points[i].w)
		allX = append(allX, points[i].x)
	}

	// 쿼리 입력
	type Query struct {
		l, r int
	}
	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].l, &queries[i].r)
	}

	// 좌표 압축: 정렬 후 중복 제거
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
		// 이진 탐색으로 압축된 인덱스 찾기
		idx := sort.SearchInts(unique, p.x)
		weightSum[idx] += int64(p.w)
	}

	// 누적합 배열 구성
	prefix := make([]int64, sz+1)
	for i := 0; i < sz; i++ {
		prefix[i+1] = prefix[i] + weightSum[i]
	}

	// 각 쿼리 처리
	for _, qr := range queries {
		// L 이상인 첫 번째 인덱스 (lower_bound)
		lo := sort.SearchInts(unique, qr.l)
		// R 이하인 마지막 인덱스 + 1 (upper_bound)
		hi := sort.SearchInts(unique, qr.r+1)

		// 구간 [lo, hi) 의 누적합
		ans := prefix[hi] - prefix[lo]
		fmt.Fprintln(writer, ans)
	}
}
