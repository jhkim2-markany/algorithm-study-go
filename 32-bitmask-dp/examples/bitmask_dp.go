package main

import "fmt"

// 비트필드 DP (Bitmask DP) - 작업 배정 문제 (Assignment Problem)
// N명의 작업자에게 N개의 작업을 1:1로 배정할 때 최소 비용을 구한다.
// 시간 복잡도: O(2^N × N)
// 공간 복잡도: O(2^N)

const INF = 1<<60 - 1

// assignmentDP: 비트마스크 DP로 최소 비용 배정을 구한다
// cost[i][j]는 작업자 i가 작업 j를 수행하는 비용이다
func assignmentDP(n int, cost [][]int) int {
	// dp[mask]: mask에 해당하는 작업들이 배정된 상태에서의 최소 비용
	// mask의 비트 수 = 지금까지 배정된 작업자 수
	full := 1 << n
	dp := make([]int, full)
	for i := range dp {
		dp[i] = INF
	}

	// 초기 상태: 아무 작업도 배정하지 않은 상태
	dp[0] = 0

	// 모든 상태를 순회한다
	for mask := 0; mask < full; mask++ {
		if dp[mask] == INF {
			continue
		}

		// 현재까지 배정된 작업자 수를 구한다 (mask에서 1인 비트 수)
		worker := popcount(mask)
		if worker >= n {
			continue
		}

		// worker번째 작업자에게 아직 배정되지 않은 작업 j를 배정한다
		for j := 0; j < n; j++ {
			// 작업 j가 이미 배정되었으면 건너뛴다
			if mask&(1<<j) != 0 {
				continue
			}

			// 작업 j를 배정한 새로운 상태
			nextMask := mask | (1 << j)
			newCost := dp[mask] + cost[worker][j]

			// 더 작은 비용이면 갱신한다
			if newCost < dp[nextMask] {
				dp[nextMask] = newCost
			}
		}
	}

	// 모든 작업이 배정된 상태의 최소 비용을 반환한다
	return dp[full-1]
}

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
	// 예시: 4명의 작업자, 4개의 작업
	// cost[i][j] = 작업자 i가 작업 j를 수행하는 비용
	cost := [][]int{
		{9, 2, 7, 8},
		{6, 4, 3, 7},
		{5, 8, 1, 8},
		{7, 6, 9, 4},
	}
	n := len(cost)

	// 비트마스크 DP로 최소 비용 배정을 구한다
	result := assignmentDP(n, cost)
	fmt.Printf("작업자 수: %d\n", n)
	fmt.Println("비용 행렬:")
	for i := 0; i < n; i++ {
		fmt.Println(" ", cost[i])
	}
	fmt.Printf("최소 배정 비용: %d\n", result)

	// 최적 배정을 역추적한다
	fmt.Println("\n최적 배정 역추적:")
	full := 1 << n
	dp := make([]int, full)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	// DP 테이블을 다시 채운다
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

	// 역추적: 최종 상태에서 거꾸로 배정을 복원한다
	mask := full - 1
	assignment := make([]int, n)
	for w := n - 1; w >= 0; w-- {
		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				continue
			}
			prevMask := mask ^ (1 << j)
			if dp[prevMask]+cost[w][j] == dp[mask] {
				assignment[w] = j
				mask = prevMask
				break
			}
		}
	}

	for w := 0; w < n; w++ {
		fmt.Printf("  작업자 %d → 작업 %d (비용: %d)\n", w, assignment[w], cost[w][assignment[w]])
	}
}
