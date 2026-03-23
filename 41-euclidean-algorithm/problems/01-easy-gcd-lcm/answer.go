package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcd는 유클리드 호제법으로 최대공약수를 계산한다.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// gcdLcm은 유클리드 호제법을 이용하여 두 수의 최대공약수와 최소공배수를 구한다.
//
// [매개변수]
//   - a: 첫 번째 자연수
//   - b: 두 번째 자연수
//
// [반환값]
//   - int: 최대공약수 (GCD)
//   - int: 최소공배수 (LCM)
//
// [알고리즘 힌트]
//   1. 유클리드 호제법: gcd(a, b) = gcd(b, a%b), b=0이면 a가 GCD이다.
//   2. LCM = a / gcd(a, b) * b (오버플로 방지를 위해 나눗셈을 먼저 수행).
func gcdLcm(a, b int) (int, int) {
	g := gcd(a, b)
	l := a / g * b
	return g, l
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b int
	fmt.Fscan(reader, &a, &b)

	g, l := gcdLcm(a, b)
	fmt.Fprintln(writer, g)
	fmt.Fprintln(writer, l)
}
