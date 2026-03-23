package main

import (
	"bufio"
	"fmt"
	"os"
)

// 페르마 소정리를 이용한 모듈러 역원 계산
// A^(-1) ≡ A^(M-2) (mod M), M이 소수일 때

// modPow는 빠른 거듭제곱으로 a^b mod m을 계산한다
func modPow(a, b, m int64) int64 {
	a %= m
	result := int64(1)
	for b > 0 {
		// 지수의 마지막 비트가 1이면 결과에 곱한다
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

	// 입력: 테스트 케이스 수
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var a, m int64
		fmt.Fscan(reader, &a, &m)

		// 페르마 소정리: A의 역원 = A^(M-2) mod M
		inverse := modPow(a, m-2, m)
		fmt.Fprintln(writer, inverse)
	}
}
