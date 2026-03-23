package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// minHamiltonianPath는 모든 도시를 정확히 한 번 방문하는 최소 비용 경로를 반환한다.
//
// [매개변수]
//   - n: 도시의 수
//   - cost: n×n 비용 행렬 (cost[u][v] = u에서 v로 이동하는 비용, 0이면 이동 불가)
//
// [반환값]
//   - int: 최소 비용 해밀턴 경로의 비용 (-1이면 경로 없음)
//
// [알고리즘 힌트]
//
//	비트마스크 DP를 사용한다.
//	dp[mask][i] = 방문 집합이 mask이고 현재 도시가 i일 때의 최소 비용.
//	모든 도시를 출발점으로 시도한다.
//	시간복잡도: O(2^N * N^2)
func minHamiltonianPath(n int, cost [][]int) int {
	full := 1 << n
	dp := make([][]int, full)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 0
	}

	for mask := 1; mask < full; mask++ {
		for u := 0; u < n; u++ {
			if dp[mask][u] == INF {
				continue
			}
			if mask&(1<<u) == 0 {
				continue
			}

			for v := 0; v < n; v++ {
				if mask&(1<<v) != 0 {
					continue
				}
				if cost[u][v] == 0 {
					continue
				}

				nextMask := mask | (1 << v)
				newCost := dp[mask][u] + cost[u][v]
				if newCost < dp[nextMask][v] {
					dp[nextMask][v] = newCost
				}
			}
		}
	}

	allVisited := full - 1
	ans := INF
	for u := 0; u < n; u++ {
		if dp[allVisited][u] < ans {
			ans = dp[allVisited][u]
		}
	}

	if ans == INF {
		return -1
	}
	return ans
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

	fmt.Fprintln(writer, minHamiltonianPath(n, cost))
}
