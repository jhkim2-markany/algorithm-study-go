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
func isPrime(n int) bool {
	// 여기에 코드를 작성하세요
	return false
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
