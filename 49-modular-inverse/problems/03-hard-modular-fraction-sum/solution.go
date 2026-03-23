package main

import (
	"bufio"
	"fmt"
	"os"
)

// modFractionSum은 분수 a_i/b_i들의 합을 모듈러 연산으로 계산한다.
// 각 분수를 a * b^(-1) mod M으로 변환하여 합산한다.
//
// [매개변수]
//   - fractions: 각 원소가 [a, b]인 분수 배열
//   - mod: 소수인 모듈러 값
//
// [반환값]
//   - int64: 모든 분수의 합 mod M
func modFractionSum(fractions [][2]int64, mod int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	fractions := make([][2]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &fractions[i][0], &fractions[i][1])
	}

	const MOD int64 = 1000000007
	fmt.Fprintln(writer, modFractionSum(fractions, MOD))
}
