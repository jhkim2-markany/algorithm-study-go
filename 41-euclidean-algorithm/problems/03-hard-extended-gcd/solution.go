package main

import (
	"bufio"
	"fmt"
	"os"
)

// 확장 유클리드 알고리즘
// ax + by = GCD(a, b)를 만족하는 g, x, y를 반환
func extGCD(a, b int) (int, int, int) {
	// 기저 조건: b = 0이면 GCD = a
	if b == 0 {
		return a, 1, 0
	}
	// 재귀적으로 해를 구함
	g, x1, y1 := extGCD(b, a%b)
	// 역추적하여 현재 단계의 계수를 계산
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}

// 모듈러 역원 계산
// a * x ≡ 1 (mod m)을 만족하는 x를 반환
func modInverse(a, m int) int {
	g, x, _ := extGCD(a, m)
	// GCD가 1이 아니면 역원이 존재하지 않음
	if g != 1 {
		return -1
	}
	// 결과를 0 이상 m 미만으로 조정
	return (x%m + m) % m
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 쿼리 수 입력
	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		// A와 M 입력
		var a, m int
		fmt.Fscan(reader, &a, &m)

		// 모듈러 역원 계산 및 출력
		fmt.Fprintln(writer, modInverse(a, m))
	}
}
