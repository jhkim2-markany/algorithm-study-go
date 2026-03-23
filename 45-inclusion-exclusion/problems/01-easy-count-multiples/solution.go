package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N과 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// K개의 소수 입력
	primes := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &primes[i])
	}

	// 비트마스크를 이용한 포함 배제
	result := 0
	for mask := 1; mask < (1 << k); mask++ {
		// 선택된 소수들의 곱을 계산한다
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

		// 켜진 비트 수가 홀수이면 더하고, 짝수이면 뺀다
		count := n / product
		if bits.OnesCount(uint(mask))%2 == 1 {
			result += count
		} else {
			result -= count
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, result)
}
