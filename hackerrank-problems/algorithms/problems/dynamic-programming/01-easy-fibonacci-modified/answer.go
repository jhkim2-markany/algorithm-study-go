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
//
// [알고리즘 힌트]
//
//	이전 두 항만 유지하며 반복적으로 계산한다.
//	값이 매우 커지므로 math/big 패키지를 사용한다.
func fibonacciModified(t1, t2 int, n int) *big.Int {
	// 큰 수로 초기화
	a := big.NewInt(int64(t1))
	b := big.NewInt(int64(t2))

	// 3번째 항부터 n번째 항까지 반복 계산
	for i := 3; i <= n; i++ {
		// t(i) = t(i-2) + t(i-1)²
		temp := new(big.Int)
		temp.Mul(b, b)    // t(i-1)²
		temp.Add(a, temp) // t(i-2) + t(i-1)²

		// 다음 반복을 위해 이동
		a = b
		b = temp
	}

	return b
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
