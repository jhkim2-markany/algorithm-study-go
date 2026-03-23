package main

import (
	"bufio"
	"fmt"
	"os"
)

// modPow는 base^exp mod m을 O(log exp)에 계산한다.
//
// [매개변수]
//   - base: 밑 (0 이상)
//   - exp: 지수 (0 이상)
//   - mod: 모듈러 값 (1 이상)
//
// [반환값]
//   - int64: base^exp mod m의 결과
//
// [알고리즘 힌트]
//
//	분할 정복 거듭제곱을 사용한다.
//	지수를 반씩 줄이며, 홀수이면 결과에 밑을 곱한다.
//	시간복잡도: O(log exp)
func modPow(base, exp, mod int64) int64 {
	result := int64(1)
	base %= mod

	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp /= 2
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
		var a, b, m int64
		fmt.Fscan(reader, &a, &b, &m)

		fmt.Fprintln(writer, modPow(a, b, m))
	}
}
