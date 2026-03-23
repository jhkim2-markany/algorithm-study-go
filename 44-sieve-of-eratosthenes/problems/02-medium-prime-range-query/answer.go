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
//
// [알고리즘 힌트]
//   1. 에라토스테네스의 체로 n 이하의 소수 여부를 구한다.
//   2. 누적합 배열 prefix[i] = 1부터 i까지의 소수 개수를 구성한다.
//   3. 각 쿼리 [l, r]에 대해 prefix[r] - prefix[l-1]로 O(1)에 응답한다.
func primeRangeQuery(n int, queries [][2]int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1]
		if isPrime[i] {
			prefix[i]++
		}
	}

	results := make([]int, len(queries))
	for i, qr := range queries {
		results[i] = prefix[qr[1]] - prefix[qr[0]-1]
	}
	return results
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
