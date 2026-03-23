package main

import (
	"bufio"
	"fmt"
	"os"
)

// 파스칼의 삼각형을 저장할 배열
var dp [31][31]int64

// 파스칼의 삼각형을 구축한다
func buildPascal() {
	for i := 0; i <= 30; i++ {
		dp[i][0] = 1 // nC0 = 1
		dp[i][i] = 1 // nCn = 1
		for j := 1; j < i; j++ {
			// 점화식: C(n, r) = C(n-1, r-1) + C(n-1, r)
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 파스칼의 삼각형 전처리
	buildPascal()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)

		// 전처리된 값을 바로 출력한다
		fmt.Fprintln(writer, dp[n][r])
	}
}
