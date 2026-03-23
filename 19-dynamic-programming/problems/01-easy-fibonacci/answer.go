package main

import (
	"bufio"
	"fmt"
	"os"
)

// fibonacci는 N번째 피보나치 수를 반환한다.
//
// [매개변수]
//   - n: 구하고자 하는 피보나치 수의 인덱스 (0-indexed)
//
// [반환값]
//   - int: N번째 피보나치 수
//
// [알고리즘 힌트]
//
//	바텀업 DP로 계산한다.
//	dp[0] = 0, dp[1] = 1로 초기화하고,
//	점화식 dp[i] = dp[i-1] + dp[i-2]를 적용한다.
//	n이 0 또는 1이면 바로 반환한다.
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	result := fibonacci(n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
