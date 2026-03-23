package main

import (
	"bufio"
	"fmt"
	"os"
)

// 유클리드 호제법으로 최대공약수 계산
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
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

	// 분수 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 결과 분수 초기화 (0/1)
	numResult := 0
	denResult := 1

	for i := 0; i < n; i++ {
		// 각 분수의 분자와 분모 입력
		var num, den int
		fmt.Fscan(reader, &num, &den)

		// 두 분수의 통분: a/b + c/d = (a*d + c*b) / (b*d)
		// LCM을 이용하여 통분하면 오버플로를 줄일 수 있음
		commonDen := lcm(denResult, den)
		numResult = numResult*(commonDen/denResult) + num*(commonDen/den)
		denResult = commonDen

		// 중간 결과를 약분하여 오버플로 방지
		g := gcd(numResult, denResult)
		if g != 0 {
			numResult /= g
			denResult /= g
		}
	}

	// 분모가 음수이면 부호를 분자로 이동
	if denResult < 0 {
		numResult = -numResult
		denResult = -denResult
	}

	// 최종 기약분수 출력
	fmt.Fprintln(writer, numResult, denResult)
}
