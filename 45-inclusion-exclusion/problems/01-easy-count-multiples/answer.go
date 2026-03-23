package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

// countMultiples는 포함 배제의 원리를 이용하여 1부터 n까지 주어진 소수 중
// 적어도 하나의 배수인 수의 개수를 구한다.
//
// [매개변수]
//   - n: 범위 상한값
//   - primes: 소수 배열
//
// [반환값]
//   - int: 1부터 n까지 primes 중 적어도 하나의 배수인 수의 개수
//
// [알고리즘 힌트]
//   1. 비트마스크로 소수의 모든 비어있지 않은 부분집합을 열거한다.
//   2. 각 부분집합에 대해 선택된 소수들의 곱(product)을 계산한다.
//   3. product > n이면 건너뛴다 (오버플로 방지).
//   4. 켜진 비트 수가 홀수이면 n/product를 더하고, 짝수이면 뺀다.
func countMultiples(n int, primes []int) int {
	k := len(primes)
	result := 0

	for mask := 1; mask < (1 << k); mask++ {
		product := 1
		overflow := false
		for i := 0; i < k; i++ {
			if mask&(1<<i) != 0 {
				product *= primes[i]
				if product > n {
					overflow = true
					break
				}
			}
		}

		if overflow {
			continue
		}

		count := n / product
		if bits.OnesCount(uint(mask))%2 == 1 {
			result += count
		} else {
			result -= count
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	primes := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &primes[i])
	}

	fmt.Fprintln(writer, countMultiples(n, primes))
}
