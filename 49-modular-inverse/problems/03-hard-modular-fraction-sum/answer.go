package main

import (
	"bufio"
	"fmt"
	"os"
)

// modFractionSum은 분수 a_i/b_i들의 합을 모듈러 연산으로 계산한다.
// 각 분수를 a * b^(-1) mod M으로 변환하여 합산한다.
//
// [매개변수]
//   - fractions: 각 원소가 [a, b]인 분수 배열
//   - mod: 소수인 모듈러 값
//
// [반환값]
//   - int64: 모든 분수의 합 mod M
//
// [알고리즘 힌트]
//
//	a/b mod M = a * b^(-1) mod M
//	b의 역원은 페르마 소정리로 b^(M-2) mod M을 빠른 거듭제곱으로 계산한다.
func modFractionSum(fractions [][2]int64, mod int64) int64 {
	modPow := func(a, b, m int64) int64 {
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

	var sum int64
	for _, f := range fractions {
		a, b := f[0], f[1]
		invB := modPow(b, mod-2, mod)
		term := a % mod * invB % mod
		sum = (sum + term) % mod
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	fractions := make([][2]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &fractions[i][0], &fractions[i][1])
	}

	const MOD int64 = 1000000007
	fmt.Fprintln(writer, modFractionSum(fractions, MOD))
}
