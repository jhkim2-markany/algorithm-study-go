package main

import (
	"bufio"
	"fmt"
	"os"
)

// 파스칼의 삼각형을 저장할 배열
var dp [31][31]int64

// buildPascal은 파스칼의 삼각형을 구축한다
func buildPascal() {
	for i := 0; i <= 30; i++ {
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
}

// binomialCoefficient는 파스칼의 삼각형으로 이항 계수 C(n, r)을 구한다.
//
// [매개변수]
//   - n: 전체 원소 수 (0 ≤ n ≤ 30)
//   - r: 선택할 원소 수 (0 ≤ r ≤ n)
//
// [반환값]
//   - int64: C(n, r) 값
//
// [알고리즘 힌트]
//
//	파스칼의 삼각형을 전처리하여 이항 계수를 구한다.
//	점화식: C(n, r) = C(n-1, r-1) + C(n-1, r)
//	기저 조건: C(n, 0) = C(n, n) = 1
//	전처리 후 O(1)로 조회한다.
func binomialCoefficient(n, r int) int64 {
	return dp[n][r]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 파스칼의 삼각형 전처리
	buildPascal()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)

		// 핵심 함수 호출
		fmt.Fprintln(writer, binomialCoefficient(n, r))
	}
}
