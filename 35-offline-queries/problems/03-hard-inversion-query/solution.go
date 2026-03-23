package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Query는 구간 쿼리 정보를 저장한다.
type Query struct {
	left  int
	right int
	index int
}

// inversionQuery는 각 쿼리 구간 [l, r]에서 역전 쌍의 개수를 반환한다.
//
// [매개변수]
//   - n: 배열의 크기
//   - arr: 1-indexed 배열 (길이 n+1, arr[0]은 미사용)
//   - queries: 쿼리 목록 (각 쿼리는 left, right, index를 포함)
//
// [반환값]
//   - []int: 각 쿼리에 대한 역전 쌍의 개수 (원래 순서)
func inversionQuery(n int, arr []int, queries []Query) []int {
	// 여기에 코드를 작성하세요
	_ = math.Sqrt(0)
	_ = sort.Slice
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].left, &queries[i].right)
		queries[i].index = i
	}

	answers := inversionQuery(n, arr, queries)
	for _, ans := range answers {
		fmt.Fprintln(writer, ans)
	}
}
