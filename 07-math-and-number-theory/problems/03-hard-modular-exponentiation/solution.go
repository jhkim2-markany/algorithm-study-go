package main

import (
	"bufio"
	"fmt"
	"os"
)

// modPow 함수는 분할 정복을 이용한 빠른 거듭제곱으로 a^n mod m을 계산한다
func modPow(a, n, m int64) int64 {
	result := int64(1)
	a = a % m

	// n의 이진 표현을 이용하여 O(log n)에 계산한다
	for n > 0 {
		// n의 마지막 비트가 1이면 결과에 a를 곱한다
		if n%2 == 1 {
			result = result * a % m
		}
		// a를 제곱하고 n을 반으로 줄인다
		n /= 2
		a = a * a % m
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// A, B, M 입력
	var a, b, m int64
	fmt.Fscan(reader, &a, &b, &m)

	// 빠른 거듭제곱으로 A^B mod M 계산
	result := modPow(a, b, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
