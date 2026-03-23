package main

import (
	"bufio"
	"fmt"
	"os"
)

// binomialCoefficient는 파스칼의 삼각형으로 이항 계수 C(n, r)을 구한다.
//
// [매개변수]
//   - n: 전체 원소 수 (0 ≤ n ≤ 30)
//   - r: 선택할 원소 수 (0 ≤ r ≤ n)
//
// [반환값]
//   - int64: C(n, r) 값
func binomialCoefficient(n, r int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)

		// 핵심 함수 호출
		fmt.Fprintln(writer, binomialCoefficient(n, r))
	}
}
