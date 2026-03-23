package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Query는 쿼리 정보를 저장하는 구조체이다
type Query struct {
	l, r, idx int
}

// mosDistinctCount는 Mo's Algorithm으로 각 구간 쿼리에 대해 서로 다른 수의 개수를 반환한다.
//
// [매개변수]
//   - a: 정수 배열 (0-indexed)
//   - queries: 쿼리 배열 (각 쿼리는 0-indexed 구간 [l, r]과 원래 인덱스)
//
// [반환값]
//   - []int: 각 쿼리에 대한 서로 다른 수의 개수 (원래 쿼리 순서)
func mosDistinctCount(a []int, queries []Query) []int {
	// 여기에 코드를 작성하세요
	_ = math.Ceil(0)
	_ = sort.Search(0, nil)
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var q int
	fmt.Fscan(reader, &q)

	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].l, &queries[i].r)
		queries[i].l--
		queries[i].r--
		queries[i].idx = i
	}

	ans := mosDistinctCount(a, queries)

	for _, v := range ans {
		fmt.Fprintln(writer, v)
	}
}
