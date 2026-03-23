package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// popcount: 정수에서 1인 비트의 개수를 센다
func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 작업자/작업 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 비용 행렬 입력
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	// dp[mask]: mask에 해당하는 작업들이 배정된 상태에서의 최소 비용
	full := 1 << n
	dp := make([]int, full)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	// 모든 상태를 순회하며 DP를 채운다
	for mask := 0; mask < full; mask++ {
		if dp[mask] == INF {
			continue
		}

		// 현재까지 배정된 작업자 수 (= mask에서 1인 비트 수)
		worker := popcount(mask)
		if worker >= n {
			continue
		}

		// worker번째 작업자에게 아직 배정되지 않은 작업 j를 배정한다
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				continue
			}
			nextMask := mask | (1 << j)
			newCost := dp[mask] + cost[worker][j]
			if newCost < dp[nextMask] {
				dp[nextMask] = newCost
			}
		}
	}

	// 모든 작업이 배정된 상태의 최소 비용 출력
	fmt.Fprintln(writer, dp[full-1])
}
