package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcdLcm은 두 자연수의 최대공약수와 최소공배수를 반환한다.
//
// [매개변수]
//   - a: 첫 번째 자연수
//   - b: 두 번째 자연수
//
// [반환값]
//   - int: 최대공약수(GCD)
//   - int: 최소공배수(LCM)
//
// [알고리즘 힌트]
//
//	유클리드 호제법으로 GCD를 구한다: gcd(a, b) = gcd(b, a%b)
//	b가 0이 될 때까지 반복하면 a가 GCD이다.
//	LCM = a / GCD * b (오버플로 방지를 위해 나눗셈을 먼저 수행)
//
//	시간복잡도: O(log(min(a, b)))
func gcdLcm(a, b int) (int, int) {
	origA, origB := a, b
	for b != 0 {
		a, b = b, a%b
	}
	g := a
	l := origA / g * origB
	return g, l
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 두 자연수 입력
	var a, b int
	fmt.Fscan(reader, &a, &b)

	// 핵심 함수 호출
	g, l := gcdLcm(a, b)

	// 결과 출력
	fmt.Fprintln(writer, g)
	fmt.Fprintln(writer, l)
}
