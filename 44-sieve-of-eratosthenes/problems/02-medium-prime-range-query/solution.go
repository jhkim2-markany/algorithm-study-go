package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N과 쿼리 수 Q 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 에라토스테네스의 체로 소수를 구한다
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

	// 누적합 배열을 구성한다
	// prefix[i] = 1부터 i까지의 소수 개수
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1]
		if isPrime[i] {
			prefix[i]++
		}
	}

	// 각 쿼리에 대해 구간 소수 개수를 O(1)에 응답한다
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		// [L, R] 구간의 소수 개수 = prefix[R] - prefix[L-1]
		fmt.Fprintln(writer, prefix[r]-prefix[l-1])
	}
}
