package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// solveTSP는 비트마스크 DP로 외판원 문제의 최소 비용을 구한다.
//
// [매개변수]
//   - cost: 도시 간 비용 행렬 (n × n)
//   - n: 도시의 수
//
// [반환값]
//   - int: 모든 도시를 방문하고 출발 도시로 돌아오는 최소 비용 (-1이면 불가능)
//
// [알고리즘 힌트]
//
//	비트마스크 DP를 사용한다.
//	dp[mask][i] = 방문한 도시 집합이 mask이고 현재 도시가 i일 때의 최소 비용.
//	dp[1][0] = 0으로 초기화 (도시 0에서 출발).
//	모든 상태를 순회하며 미방문 도시로의 전이를 계산한다.
//	모든 도시를 방문한 후 출발 도시(0)로 돌아오는 최소 비용을 구한다.
func solveTSP(cost [][]int, n int) int {
	full := 1 << n
	dp := make([][]int, full)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	dp[1][0] = 0

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
		if dp[allVisited][u] == INF {
			continue
		}
		if cost[u][0] == 0 && u != 0 {
			continue
		}
		if u == 0 {
			if dp[allVisited][u] < ans {
				ans = dp[allVisited][u]
			}
		} else {
			total := dp[allVisited][u] + cost[u][0]
			if total < ans {
				ans = total
			}
		}
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

	// 핵심 함수 호출
	result := solveTSP(cost, n)

	if result == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, result)
	}
}
