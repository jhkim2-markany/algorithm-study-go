package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// sqrtRangeUpdateQuery는 Sqrt Decomposition(Lazy 블록)을 이용한 구간 갱신 + 구간 합 자료구조를 구현한다.
//
// [매개변수]
//   - n: 배열의 크기
//   - a: 정수 배열 (0-indexed)
//
// [반환값]
//   - rangeUpdate func(l, r, v int): [l, r] 구간에 v를 더하는 함수
//   - rangeQuery func(l, r int) int: [l, r] 구간 합을 반환하는 함수
func sqrtRangeUpdateQuery(n int, a []int) (rangeUpdate func(l, r, v int), rangeQuery func(l, r int) int) {
	// 여기에 코드를 작성하세요
	_ = math.Ceil(0)
	rangeUpdate = func(l, r, v int) {}
	rangeQuery = func(l, r int) int { return 0 }
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	rangeUpdate, rangeQuery := sqrtRangeUpdateQuery(n, a)

	for ; q > 0; q-- {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			var l, r, v int
			fmt.Fscan(reader, &l, &r, &v)
			rangeUpdate(l-1, r-1, v)
		} else {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, rangeQuery(l-1, r-1))
		}
	}
}
