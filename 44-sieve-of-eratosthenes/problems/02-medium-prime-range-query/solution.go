package main

import (
	"bufio"
	"fmt"
	"os"
)

// primeRangeQuery는 에라토스테네스의 체와 누적합을 이용하여 구간 소수 개수 쿼리에 응답한다.
//
// [매개변수]
//   - n: 소수를 구할 상한값
//   - queries: 구간 쿼리 배열 (각 원소는 [2]int{l, r})
//
// [반환값]
//   - []int: 각 쿼리에 대한 [l, r] 구간의 소수 개수 배열
func primeRangeQuery(n int, queries [][2]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	queries := make([][2]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	results := primeRangeQuery(n, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
