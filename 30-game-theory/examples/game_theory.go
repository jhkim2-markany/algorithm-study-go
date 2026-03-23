package main

import "fmt"

// 게임 이론 (Game Theory) - Nim 게임과 Sprague-Grundy 정리
// 두 플레이어가 번갈아 최적의 수를 두는 조합 게임의 승패를 판별한다.
// Nim 게임 시간 복잡도: O(N)
// 그런디 수 계산 시간 복잡도: O(S × M) (S: 상태 수, M: 이동 수)

// nimGameResult: N개 돌 더미의 Nim 게임 승패를 판별한다
// 모든 더미의 XOR이 0이 아니면 선수 승리, 0이면 후수 승리
func nimGameResult(piles []int) bool {
	xorSum := 0
	for _, p := range piles {
		xorSum ^= p
	}
	// XOR 합이 0이 아니면 선수(첫 번째 플레이어) 승리
	return xorSum != 0
}

// mex: 집합에 포함되지 않는 가장 작은 음이 아닌 정수를 구한다
// Minimum Excludant - Sprague-Grundy 정리의 핵심 연산
func mex(set map[int]bool) int {
	val := 0
	for set[val] {
		val++
	}
	return val
}

// grundy: 돌 n개인 상태에서 moves에 지정된 개수만큼 가져갈 수 있을 때
// 그런디 수를 계산한다 (메모이제이션 사용)
func grundy(n int, moves []int, memo map[int]int) int {
	// 이미 계산된 상태면 바로 반환한다
	if val, ok := memo[n]; ok {
		return val
	}

	// 이동 가능한 상태들의 그런디 수를 수집한다
	reachable := make(map[int]bool)
	for _, m := range moves {
		if n >= m {
			reachable[grundy(n-m, moves, memo)] = true
		}
	}

	// mex를 적용하여 현재 상태의 그런디 수를 결정한다
	result := mex(reachable)
	memo[n] = result
	return result
}

// winLosePosition: 돌 n개에서 1~3개를 가져갈 수 있을 때
// 각 상태의 승패를 DP로 판별한다 (true: 선수 승리)
func winLosePosition(n int) []bool {
	// dp[i] = true이면 돌 i개인 상태에서 선수 승리
	dp := make([]bool, n+1)
	// 돌 0개: 수를 둘 수 없으므로 패배 (false)
	for i := 1; i <= n; i++ {
		// 1~3개를 가져가서 상대를 패배 포지션으로 보낼 수 있는지 확인
		for take := 1; take <= 3 && take <= i; take++ {
			if !dp[i-take] {
				// 상대가 패배 포지션에 놓이므로 현재 상태는 승리
				dp[i] = true
				break
			}
		}
	}
	return dp
}

func main() {
	// === Nim 게임 예시 ===
	fmt.Println("=== Nim 게임 ===")
	piles1 := []int{3, 4, 5}
	fmt.Printf("돌 더미: %v\n", piles1)
	if nimGameResult(piles1) {
		fmt.Println("결과: 선수 승리")
	} else {
		fmt.Println("결과: 후수 승리")
	}

	piles2 := []int{1, 2, 3}
	fmt.Printf("\n돌 더미: %v\n", piles2)
	if nimGameResult(piles2) {
		fmt.Println("결과: 선수 승리")
	} else {
		fmt.Println("결과: 후수 승리")
	}

	// === 승리/패배 포지션 예시 ===
	fmt.Println("\n=== 승리/패배 포지션 (1~3개 가져가기) ===")
	n := 12
	dp := winLosePosition(n)
	for i := 0; i <= n; i++ {
		status := "패배(L)"
		if dp[i] {
			status = "승리(W)"
		}
		fmt.Printf("돌 %2d개: %s\n", i, status)
	}

	// === Sprague-Grundy 예시 ===
	fmt.Println("\n=== Sprague-Grundy (1, 3, 4개 가져가기) ===")
	moves := []int{1, 3, 4}
	memo := make(map[int]int)
	fmt.Printf("가져갈 수 있는 개수: %v\n", moves)
	for i := 0; i <= 10; i++ {
		g := grundy(i, moves, memo)
		fmt.Printf("돌 %2d개: 그런디 수 = %d\n", i, g)
	}

	// === 복합 게임 예시 ===
	fmt.Println("\n=== 복합 게임 (여러 독립 게임의 XOR) ===")
	// 3개의 독립적인 게임, 각각 돌 5, 3, 7개
	games := []int{5, 3, 7}
	moves2 := []int{1, 3, 4}
	memo2 := make(map[int]int)
	totalXor := 0
	for _, g := range games {
		gn := grundy(g, moves2, memo2)
		fmt.Printf("게임(돌 %d개): 그런디 수 = %d\n", g, gn)
		totalXor ^= gn
	}
	fmt.Printf("전체 XOR = %d\n", totalXor)
	if totalXor != 0 {
		fmt.Println("결과: 선수 승리")
	} else {
		fmt.Println("결과: 후수 승리")
	}
}
