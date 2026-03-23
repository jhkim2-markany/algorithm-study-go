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
	var n, k int
	fmt.Fscan(reader, &n, &k)

	weight := make([]int, n+1)
	value := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &weight[i], &value[i])
	}

	// dp[j] = 용량 j일 때 담을 수 있는 최대 가치
	// 1차원 배열로 공간 최적화한 바텀업 DP
	dp := make([]int, k+1)

	for i := 1; i <= n; i++ {
		// 역순으로 순회하여 같은 물건을 중복 사용하지 않도록 한다
		for j := k; j >= weight[i]; j-- {
			// 점화식: i번째 물건을 넣는 경우와 넣지 않는 경우 중 최댓값
			if dp[j-weight[i]]+value[i] > dp[j] {
				dp[j] = dp[j-weight[i]] + value[i]
			}
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, dp[k])
}
