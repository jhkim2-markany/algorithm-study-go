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
func modPow(a, n, m int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
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
