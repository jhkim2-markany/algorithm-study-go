package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 도시 수 입력
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

	// dp[mask][i]: 방문한 도시 집합이 mask이고 현재 도시가 i일 때의 최소 비용
	full := 1 << n
	dp := make([][]int, full)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	// 시작 도시 0에서 출발
	dp[1][0] = 0

	// 모든 상태를 순회하며 DP를 채운다
	for mask := 1; mask < full; mask++ {
		for u := 0; u < n; u++ {
			if dp[mask][u] == INF {
				continue
			}
			// 현재 도시 u가 mask에 포함되어 있는지 확인한다
			if mask&(1<<u) == 0 {
				continue
			}

			// 아직 방문하지 않은 도시 v로 이동한다
			for v := 0; v < n; v++ {
				// 이미 방문한 도시는 건너뛴다
				if mask&(1<<v) != 0 {
					continue
				}
				// 이동할 수 없는 경로는 건너뛴다
				if cost[u][v] == 0 {
					continue
				}

				// 도시 v를 방문한 새로운 상태
				nextMask := mask | (1 << v)
				newCost := dp[mask][u] + cost[u][v]

				// 더 작은 비용이면 갱신한다
				if newCost < dp[nextMask][v] {
					dp[nextMask][v] = newCost
				}
			}
		}
	}

	// 모든 도시를 방문한 후 출발 도시(0)로 돌아오는 최소 비용을 구한다
	allVisited := full - 1
	ans := INF
	for u := 0; u < n; u++ {
		if dp[allVisited][u] == INF {
			continue
		}
		// 도시 u에서 도시 0으로 돌아갈 수 있는지 확인한다
		if cost[u][0] == 0 && u != 0 {
			continue
		}
		if u == 0 {
			// 이미 출발 도시에 있는 경우 (N=1일 때)
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

	// 결과 출력
	if ans == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}
