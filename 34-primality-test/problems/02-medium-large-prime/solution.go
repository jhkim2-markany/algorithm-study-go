package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// 모듈러 거듭제곱: (base^exp) mod m 을 계산한다
// 오버플로우 방지를 위해 big.Int를 사용한다
func modPow(base, exp, m int64) int64 {
	b := big.NewInt(base)
	e := big.NewInt(exp)
	mod := big.NewInt(m)
	result := new(big.Int).Exp(b, e, mod)
	return result.Int64()
}

// 밀러-라빈 소수 판정법
// 결정적 밀러-라빈: 특정 밑 집합으로 int64 범위 내에서 정확히 판정한다
func isPrime(n int64) bool {
	// 2 미만은 소수가 아니다
	if n < 2 {
		return false
	}
	// 작은 소수를 밑으로 사용한다
	bases := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, p := range bases {
		if n == p {
			return true
		}
		if n%p == 0 {
			return false
		}
	}

	// n-1 = 2^s × d 형태로 분해한다 (d는 홀수)
	d := n - 1
	s := 0
	for d%2 == 0 {
		d /= 2
		s++
	}

	// 각 밑에 대해 밀러-라빈 테스트를 수행한다
	for _, a := range bases {
		if a >= n {
			continue
		}
		// x = a^d mod n 을 계산한다
		x := modPow(a, d, n)

		// x가 1이거나 n-1이면 이 밑에 대해 통과
		if x == 1 || x == n-1 {
			continue
		}

		// s-1번 제곱하며 n-1이 나오는지 확인한다
		passed := false
		for r := 0; r < s-1; r++ {
			x = modPow(x, 2, n)
			if x == n-1 {
				passed = true
				break
			}
		}
		// 통과하지 못하면 합성수이다
		if !passed {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscan(reader, &n)

		// 밀러-라빈으로 소수 여부를 판정한다
		if isPrime(n) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
