package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// popcount는 정수에서 1인 비트의 개수를 센다.
//
// [매개변수]
//   - x: 비트 수를 셀 정수
//
// [반환값]
//   - int: x에서 1인 비트의 개수
func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// minCostAssignment는 n명의 작업자에게 n개의 작업을 1:1 배정할 때 최소 비용을 반환한다.
//
// [매개변수]
//   - n: 작업자/작업의 수
//   - cost: n×n 비용 행렬 (cost[i][j] = 작업자 i가 작업 j를 수행하는 비용)
//
// [반환값]
//   - int: 모든 작업을 배정했을 때의 최소 총 비용
//
// [알고리즘 힌트]
//
//	비트마스크 DP를 사용한다.
//	dp[mask] = mask에 해당하는 작업들이 배정된 상태의 최소 비용.
//	popcount(mask)번째 작업자에게 미배정 작업을 배정한다.
//	시간복잡도: O(2^N * N)
func minCostAssignment(n int, cost [][]int) int {
	full := 1 << n
	dp := make([]int, full)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	for mask := 0; mask < full; mask++ {
		if dp[mask] == INF {
			continue
		}

		worker := popcount(mask)
		if worker >= n {
			continue
		}

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

	return dp[full-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	fmt.Fprintln(writer, minCostAssignment(n, cost))
}
