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
func modPow(base, exp, mod int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
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
