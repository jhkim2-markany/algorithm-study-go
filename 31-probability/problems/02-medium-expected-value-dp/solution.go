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

	// 보드의 마지막 칸 번호 입력
	var n int
	fmt.Fscan(reader, &n)

	// dp[i] = i번 칸에서 N번 칸에 도착하기까지의 기대 주사위 횟수
	// dp[n] = 0 (이미 도착)
	// dp[i] = 1 + (1/6) × Σ dp[next] (next = i+1 ~ i+6, N 초과 시 제자리)
	dp := make([]float64, n+1)

	// N번 칸부터 역순으로 기댓값을 계산한다
	for i := n - 1; i >= 1; i-- {
		sum := 0.0
		stay := 0 // N을 넘어서 제자리에 머무는 경우의 수

		for face := 1; face <= 6; face++ {
			next := i + face
			if next == n {
				// 정확히 도착: 기댓값 기여 0
				sum += 0
			} else if next > n {
				// N을 넘어가면 제자리에 머문다
				stay++
			} else {
				sum += dp[next]
			}
		}

		// dp[i] = 1 + (sum + stay × dp[i]) / 6
		// 6 × dp[i] = 6 + sum + stay × dp[i]
		// (6 - stay) × dp[i] = 6 + sum
		// dp[i] = (6 + sum) / (6 - stay)
		dp[i] = (6.0 + sum) / float64(6-stay)
	}

	// 1번 칸에서의 기댓값 출력
	fmt.Fprintf(writer, "%.6f\n", dp[1])
}
