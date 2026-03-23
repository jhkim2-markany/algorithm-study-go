package main

import (
	"bufio"
	"fmt"
	"os"
)

// extGCD는 확장 유클리드 알고리즘으로 ax + by = gcd(a, b)를 만족하는 g, x, y를 반환한다.
func extGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := extGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}

// modInverse는 확장 유클리드 알고리즘을 이용하여 a의 모듈러 역원을 구한다.
// a * x ≡ 1 (mod m)을 만족하는 x를 반환한다.
//
// [매개변수]
//   - a: 역원을 구할 정수
//   - m: 모듈러 값
//
// [반환값]
//   - int: 모듈러 역원 (0 이상 m 미만), 역원이 없으면 -1
//
// [알고리즘 힌트]
//   1. 확장 유클리드 알고리즘으로 ax + my = gcd(a, m)을 푼다.
//   2. gcd(a, m) ≠ 1이면 역원이 존재하지 않으므로 -1을 반환한다.
//   3. x를 (x % m + m) % m으로 조정하여 0 이상 m 미만으로 만든다.
func modInverse(a, m int) int {
	g, x, _ := extGCD(a, m)
	if g != 1 {
		return -1
	}
	return (x%m + m) % m
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var a, m int
		fmt.Fscan(reader, &a, &m)
		fmt.Fprintln(writer, modInverse(a, m))
	}
}
