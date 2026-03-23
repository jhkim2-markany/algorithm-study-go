package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// sqrtRangeSum은 Sqrt Decomposition을 이용한 구간 합 자료구조를 구현한다.
// update는 점 갱신, query는 구간 합 쿼리를 처리한다.
//
// [매개변수]
//   - n: 배열의 크기
//   - a: 정수 배열 (0-indexed)
//
// [반환값]
//   - update func(idx, val int): a[idx]를 val로 변경하는 함수
//   - query func(l, r int) int: [l, r] 구간 합을 반환하는 함수
func sqrtRangeSum(n int, a []int) (update func(idx, val int), query func(l, r int) int) {
	// 여기에 코드를 작성하세요
	update = func(idx, val int) {}
	query = func(l, r int) int { return 0 }
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

	_ = math.Ceil(0) // math 패키지 사용

	update, query := sqrtRangeSum(n, a)

	for ; q > 0; q-- {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			var idx, val int
			fmt.Fscan(reader, &idx, &val)
			update(idx-1, val)
		} else {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, query(l-1, r-1))
		}
	}
}
