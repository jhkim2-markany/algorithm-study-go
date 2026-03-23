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

	// 아이템 종류의 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 각 아이템의 확률 입력
	prob := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &prob[i])
	}

	// 비트마스크 DP로 기댓값을 계산한다
	// dp[mask] = 이미 모은 아이템 집합이 mask일 때, 모든 아이템을 모으기까지의 기대 횟수
	// 전체 상태 수: 2^N
	full := (1 << n) - 1
	dp := make([]float64, 1<<n)

	// dp[full] = 0 (모든 아이템을 이미 모음)
	// 역순으로 계산한다
	for mask := full - 1; mask >= 0; mask-- {
		// 현재 상태에서 이미 가진 아이템을 다시 뽑을 확률
		pRepeat := 0.0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				// i번째 아이템은 이미 가지고 있다
				pRepeat += prob[i]
			}
		}

		// 새로운 아이템을 뽑을 확률
		pNew := 1.0 - pRepeat

		// dp[mask] = 1 + pRepeat × dp[mask] + Σ (prob[i] × dp[mask | (1<<i)]) (i가 mask에 없는 경우)
		// 정리: dp[mask] × (1 - pRepeat) = 1 + Σ (prob[i] × dp[mask | (1<<i)])
		// dp[mask] = (1 + Σ (prob[i] × dp[mask | (1<<i)])) / pNew

		sum := 0.0
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				// i번째 아이템을 새로 뽑는 경우
				sum += prob[i] * dp[mask|(1<<i)]
			}
		}

		dp[mask] = (1.0 + sum) / pNew
	}

	// 아무것도 모으지 않은 상태(mask=0)에서의 기댓값 출력
	fmt.Fprintf(writer, "%.6f\n", dp[0])
}
