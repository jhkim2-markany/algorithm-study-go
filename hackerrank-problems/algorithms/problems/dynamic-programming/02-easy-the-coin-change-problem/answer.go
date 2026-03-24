package main

import (
	"bufio"
	"fmt"
	"os"
)

// coinChange는 주어진 동전으로 금액 n을 만드는 방법의 수를 반환한다.
//
// [매개변수]
//   - n: 목표 금액
//   - coins: 동전 액면가 배열
//
// [반환값]
//   - int64: 금액 n을 만드는 방법의 수
//
// [알고리즘 힌트]
//
//	1차원 DP 배열을 사용한다. 동전 종류별로 순회하며
//	dp[j] += dp[j - coin]으로 갱신하면 중복 조합을 방지한다.
func coinChange(n int, coins []int) int64 {
	// dp[i] = 금액 i를 만드는 방법의 수
	dp := make([]int64, n+1)

	// 금액 0을 만드는 방법은 1가지 (아무것도 선택하지 않음)
	dp[0] = 1

	// 각 동전에 대해 dp 테이블 갱신
	for _, coin := range coins {
		for j := coin; j <= n; j++ {
			// 현재 동전을 사용하는 경우를 추가
			dp[j] += dp[j-coin]
		}
	}

	return dp[n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 금액과 동전 종류 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 동전 액면가 입력
	coins := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &coins[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := coinChange(n, coins)
	fmt.Fprintln(writer, result)
}
