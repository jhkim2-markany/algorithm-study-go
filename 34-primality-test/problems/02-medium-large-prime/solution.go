package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// isPrime은 밀러-라빈 소수 판정법으로 큰 정수가 소수인지 판정한다.
//
// [매개변수]
//   - n: 판정할 양의 정수 (int64 범위)
//
// [반환값]
//   - bool: 소수이면 true, 아니면 false
func isPrime(n int64) bool {
	// 여기에 코드를 작성하세요
	_ = big.NewInt(0) // big 패키지 사용 힌트
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscan(reader, &n)

		if isPrime(n) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
