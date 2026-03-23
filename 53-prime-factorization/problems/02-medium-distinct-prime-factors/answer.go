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
//
// [알고리즘 힌트]
//
//	SPF(최소 소인수) 체를 구축한 뒤, 각 수를 SPF로 반복 나누며 서로 다른 소인수를 센다.
//	전처리 O(N log log N), 쿼리 O(1).
func countDistinctPrimeFactors(maxN int, queries []int) []int {
	// SPF 체 구축
	spf := make([]int, maxN+1)
	for i := 2; i <= maxN; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= maxN; i++ {
		if spf[i] == i {
			for j := i * i; j <= maxN; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	// 각 수의 서로 다른 소인수 개수 전처리
	distinctCount := make([]int, maxN+1)
	for i := 2; i <= maxN; i++ {
		x := i
		cnt := 0
		for x > 1 {
			p := spf[x]
			cnt++
			for x%p == 0 {
				x /= p
			}
		}
		distinctCount[i] = cnt
	}

	results := make([]int, len(queries))
	for i, q := range queries {
		results[i] = distinctCount[q]
	}
	return results
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
