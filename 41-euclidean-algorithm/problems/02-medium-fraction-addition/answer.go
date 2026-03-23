package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcd는 유클리드 호제법으로 최대공약수를 계산한다 (음수 처리 포함).
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

// lcm은 GCD를 이용하여 최소공배수를 계산한다.
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// addFractions는 여러 분수의 합을 기약분수로 반환한다.
//
// [매개변수]
//   - nums: 각 분수의 분자 배열
//   - dens: 각 분수의 분모 배열
//
// [반환값]
//   - int: 결과 기약분수의 분자
//   - int: 결과 기약분수의 분모
//
// [알고리즘 힌트]
//   1. 결과 분수를 0/1로 초기화한다.
//   2. 각 분수를 순회하며 LCM으로 통분하여 더한다.
//   3. 매 단계마다 GCD로 약분하여 오버플로를 방지한다.
//   4. 최종 결과에서 분모가 음수이면 부호를 분자로 이동한다.
func addFractions(nums, dens []int) (int, int) {
	numResult := 0
	denResult := 1

	for i := 0; i < len(nums); i++ {
		commonDen := lcm(denResult, dens[i])
		numResult = numResult*(commonDen/denResult) + nums[i]*(commonDen/dens[i])
		denResult = commonDen

		g := gcd(numResult, denResult)
		if g != 0 {
			numResult /= g
			denResult /= g
		}
	}

	if denResult < 0 {
		numResult = -numResult
		denResult = -denResult
	}

	return numResult, denResult
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	dens := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i], &dens[i])
	}

	num, den := addFractions(nums, dens)
	fmt.Fprintln(writer, num, den)
}
