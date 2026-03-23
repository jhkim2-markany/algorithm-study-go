package main

import (
	"bufio"
	"fmt"
	"os"
)

// isPrime은 주어진 정수가 소수인지 판정한다.
//
// [매개변수]
//   - n: 판정할 양의 정수
//
// [반환값]
//   - bool: 소수이면 true, 아니면 false
//
// [알고리즘 힌트]
//
//	시행 나눗셈법을 사용한다.
//	2 미만은 소수가 아니고, 짝수를 먼저 제외한 뒤
//	3부터 √N까지 홀수로 나누어 본다.
//	시간복잡도: O(√N)
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
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

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		if isPrime(n) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
