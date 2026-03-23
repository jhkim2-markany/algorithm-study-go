package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)

	// 기저 사례 처리
	if n <= 1 {
		fmt.Fprintln(writer, n)
		return
	}

	// 바텀업 DP로 피보나치 수를 계산한다
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	// 점화식: dp[i] = dp[i-1] + dp[i-2]
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 결과 출력
	fmt.Fprintln(writer, dp[n])
}
