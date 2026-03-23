package main

import (
	"bufio"
	"fmt"
	"os"
)

// 시행 나눗셈으로 소수를 판정한다
func isPrime(n int) bool {
	// 2 미만은 소수가 아니다
	if n < 2 {
		return false
	}
	// 2는 소수이다
	if n == 2 {
		return true
	}
	// 짝수는 소수가 아니다
	if n%2 == 0 {
		return false
	}
	// 3부터 √N까지 홀수로 나누어 본다
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		// 소수 여부를 판별하여 출력한다
		if isPrime(n) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
