package main

import (
	"bufio"
	"fmt"
	"os"
)

// 모듈러 분수의 합
// 각 분수 a/b를 a * b^(-1) mod M으로 변환하여 합산한다
// 역원은 페르마 소정리로 b^(M-2) mod M을 계산한다

const MOD = 1000000007

// modPow는 빠른 거듭제곱으로 a^b mod m을 계산한다
func modPow(a, b, m int64) int64 {
	a %= m
	if a < 0 {
		a += m
	}
	result := int64(1)
	for b > 0 {
		if b%2 == 1 {
			result = result * a % m
		}
		b /= 2
		a = a * a % m
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 분수의 개수
	var n int
	fmt.Fscan(reader, &n)

	var sum int64

	for i := 0; i < n; i++ {
		var a, b int64
		fmt.Fscan(reader, &a, &b)

		// a/b mod M = a * b^(-1) mod M
		// b의 역원을 페르마 소정리로 계산한다
		invB := modPow(b, MOD-2, MOD)

		// a * invB를 합산한다
		term := a % MOD * invB % MOD
		sum = (sum + term) % MOD
	}

	fmt.Fprintln(writer, sum)
}
