package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// fibonacciModified는 변형 피보나치 수열의 n번째 항을 반환한다.
// 점화식: t(i+2) = t(i) + t(i+1)²
//
// [매개변수]
//   - t1: 첫 번째 항
//   - t2: 두 번째 항
//   - n: 구하고자 하는 항의 번호
//
// [반환값]
//   - *big.Int: n번째 항의 값
func fibonacciModified(t1, t2 int, n int) *big.Int {
	// 여기에 코드를 작성하세요
	return big.NewInt(0)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: t1, t2, n
	var t1, t2, n int
	fmt.Fscan(reader, &t1, &t2, &n)

	// 핵심 함수 호출 및 결과 출력
	result := fibonacciModified(t1, t2, n)
	fmt.Fprintln(writer, result.String())
}
