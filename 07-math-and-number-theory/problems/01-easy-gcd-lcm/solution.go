package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcd 함수는 유클리드 호제법으로 최대공약수를 구한다
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 두 자연수 입력
	var a, b int
	fmt.Fscan(reader, &a, &b)

	// 최대공약수 계산
	g := gcd(a, b)

	// 최소공배수 계산: LCM = a / GCD * b (오버플로 방지)
	l := a / g * b

	// 결과 출력
	fmt.Fprintln(writer, g)
	fmt.Fprintln(writer, l)
}
