package main

import (
	"bufio"
	"fmt"
	"os"
)

// modInverse는 페르마 소정리를 이용하여 a의 모듈러 역원 a^(m-2) mod m을 반환한다.
//
// [매개변수]
//   - a: 역원을 구할 정수
//   - m: 소수인 모듈러 값
//
// [반환값]
//   - int64: a^(-1) mod m
//
// [알고리즘 힌트]
//
//	페르마 소정리: m이 소수일 때 a^(-1) ≡ a^(m-2) (mod m)
//	빠른 거듭제곱으로 O(log m)에 계산한다.
func modInverse(a, m int64) int64 {
	a %= m
	result := int64(1)
	b := m - 2
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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var a, m int64
		fmt.Fscan(reader, &a, &m)
		fmt.Fprintln(writer, modInverse(a, m))
	}
}
