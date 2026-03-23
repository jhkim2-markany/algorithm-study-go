package main

import (
	"bufio"
	"fmt"
	"os"
)

// modPow는 a^n mod m을 빠른 거듭제곱으로 계산하여 반환한다.
//
// [매개변수]
//   - a: 밑 (base)
//   - n: 지수 (exponent)
//   - m: 나눌 수 (modulus)
//
// [반환값]
//   - int64: a^n mod m의 결과
//
// [알고리즘 힌트]
//
//	분할 정복을 이용한 빠른 거듭제곱(Exponentiation by Squaring).
//	n의 이진 표현을 이용하여 O(log n)에 계산한다.
//	n의 마지막 비트가 1이면 결과에 a를 곱하고,
//	매 단계에서 a를 제곱하고 n을 반으로 줄인다.
//	모든 곱셈 후 mod m을 취하여 오버플로를 방지한다.
//
//	시간복잡도: O(log N)
func modPow(a, n, m int64) int64 {
	result := int64(1)
	a = a % m
	for n > 0 {
		if n%2 == 1 {
			result = result * a % m
		}
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

	// 핵심 함수 호출
	result := modPow(a, b, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
