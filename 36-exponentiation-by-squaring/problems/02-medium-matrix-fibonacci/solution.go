package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

// Matrix는 2x2 행렬 타입이다.
type Matrix [2][2]int64

// fibonacci는 N번째 피보나치 수를 행렬 거듭제곱으로 구한다.
//
// [매개변수]
//   - n: 구할 피보나치 수의 인덱스 (0 이상)
//
// [반환값]
//   - int64: F(n) mod 10^9+7
func fibonacci(n int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int64
	fmt.Fscan(reader, &n)

	fmt.Fprintln(writer, fibonacci(n))
}
