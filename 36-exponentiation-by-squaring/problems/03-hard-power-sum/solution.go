package main

import (
	"bufio"
	"fmt"
	"os"
)

// powerSum은 S(N) = A^1 + A^2 + ... + A^N mod M을 계산한다.
//
// [매개변수]
//   - a: 밑 (1 이상)
//   - n: 지수 합의 상한 (1 이상)
//   - m: 모듈러 값 (2 이상)
//
// [반환값]
//   - int64: (A^1 + A^2 + ... + A^N) mod M
func powerSum(a, n, m int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, n, m int64
	fmt.Fscan(reader, &a, &n, &m)

	fmt.Fprintln(writer, powerSum(a, n, m))
}
