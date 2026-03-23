package main

import (
	"bufio"
	"fmt"
	"os"
)

// 유클리드 호제법으로 최대공약수 계산
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GCD를 이용한 최소공배수 계산
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 두 자연수 입력
	var a, b int
	fmt.Fscan(reader, &a, &b)

	// GCD와 LCM 출력
	fmt.Fprintln(writer, gcd(a, b))
	fmt.Fprintln(writer, lcm(a, b))
}
