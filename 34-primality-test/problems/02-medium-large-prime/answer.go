package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// modPow는 (base^exp) mod m을 big.Int로 계산한다.
func modPow(base, exp, m int64) int64 {
	b := big.NewInt(base)
	e := big.NewInt(exp)
	mod := big.NewInt(m)
	result := new(big.Int).Exp(b, e, mod)
	return result.Int64()
}

// isPrime은 밀러-라빈 소수 판정법으로 큰 정수가 소수인지 판정한다.
//
// [매개변수]
//   - n: 판정할 양의 정수 (int64 범위)
//
// [반환값]
//   - bool: 소수이면 true, 아니면 false
//
// [알고리즘 힌트]
//
//	결정적 밀러-라빈 판정법을 사용한다.
//	n-1 = 2^s * d로 분해한 뒤, 특정 밑 집합으로 테스트한다.
//	int64 범위에서 {2,3,5,7,11,13,17,19,23,29,31,37}이면 정확하다.
//	시간복잡도: O(k * log^2(N)) (k는 밑의 수)
func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	bases := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, p := range bases {
		if n == p {
			return true
		}
		if n%p == 0 {
			return false
		}
	}

	d := n - 1
	s := 0
	for d%2 == 0 {
		d /= 2
		s++
	}

	for _, a := range bases {
		if a >= n {
			continue
		}
		x := modPow(a, d, n)

		if x == 1 || x == n-1 {
			continue
		}

		passed := false
		for r := 0; r < s-1; r++ {
			x = modPow(x, 2, n)
			if x == n-1 {
				passed = true
				break
			}
		}
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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscan(reader, &n)

		if isPrime(n) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
