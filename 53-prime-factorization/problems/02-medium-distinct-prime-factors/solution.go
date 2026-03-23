package main

import (
	"bufio"
	"fmt"
	"os"
)

// countDistinctPrimeFactors는 SPF 체를 이용하여 1부터 maxN까지 각 수의
// 서로 다른 소인수 개수를 전처리하고, 쿼리에 O(1)로 응답한다.
//
// [매개변수]
//   - maxN: 전처리할 최대 범위
//   - queries: 서로 다른 소인수 개수를 구할 정수 배열
//
// [반환값]
//   - []int: 각 쿼리에 대한 서로 다른 소인수 개수
func countDistinctPrimeFactors(maxN int, queries []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	queries := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i])
	}

	results := countDistinctPrimeFactors(n, queries)
	for _, v := range results {
		fmt.Fprintln(writer, v)
	}
}
