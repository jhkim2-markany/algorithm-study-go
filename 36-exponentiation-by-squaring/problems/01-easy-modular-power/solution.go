package main

import (
	"bufio"
	"fmt"
	"os"
)

// 모듈러 거듭제곱: base^exp mod m을 O(log exp)에 계산한다
func modPow(base, exp, mod int64) int64 {
	result := int64(1)
	base %= mod

	// 지수를 반씩 줄이며 거듭제곱을 계산한다
	for exp > 0 {
		// 지수가 홀수이면 결과에 밑을 곱한다
		if exp%2 == 1 {
			result = result * base % mod
		}
		// 밑을 제곱하고 지수를 반으로 나눈다
		base = base * base % mod
		exp /= 2
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		// A, B, M 입력
		var a, b, m int64
		fmt.Fscan(reader, &a, &b, &m)

		// A^B mod M 계산 후 출력
		fmt.Fprintln(writer, modPow(a, b, m))
	}
}
