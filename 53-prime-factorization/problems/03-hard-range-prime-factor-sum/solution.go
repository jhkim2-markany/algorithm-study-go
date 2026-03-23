package main

import (
	"bufio"
	"fmt"
	"os"
)

// rangePrimeFactorSum은 구간 [l, r]에 속하는 각 수의 소인수 합의 총합을 반환한다.
// 소인수가 여러 번 나누어지면 그 횟수만큼 더한다.
//
// [매개변수]
//   - l: 구간 시작 (l >= 2)
//   - r: 구간 끝
//
// [반환값]
//   - int64: 구간 내 모든 수의 소인수 합의 총합
func rangePrimeFactorSum(l, r int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, r int
	fmt.Fscan(reader, &l, &r)

	fmt.Fprintln(writer, rangePrimeFactorSum(l, r))
}
