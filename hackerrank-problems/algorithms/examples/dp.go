package main

import "fmt"

// 동적 프로그래밍 (Dynamic Programming) - 기본 패턴 예시
// 큰 문제를 작은 부분 문제로 나누어 풀고, 그 결과를 저장하여 중복 계산을 피하는 방법이다.
// DP의 두 가지 핵심 조건:
//   1. 최적 부분 구조 (Optimal Substructure): 큰 문제의 최적해가 부분 문제의 최적해로 구성된다
//   2. 중복 부분 문제 (Overlapping Subproblems): 동일한 부분 문제가 여러 번 반복된다
//
// DP 구현 방식:
//   - 메모이제이션 (Memoization, Top-Down): 재귀 + 캐싱으로 위에서 아래로 풀어나간다
//   - 타뷸레이션 (Tabulation, Bottom-Up): 반복문으로 아래에서 위로 테이블을 채워나간다
//
// 예시 1: 피보나치 수열 (메모이제이션 - Top-Down)
//   - 시간 복잡도: O(N) (각 부분 문제를 한 번만 계산)
//   - 공간 복잡도: O(N) (메모 배열 + 재귀 호출 스택)
//
// 예시 2: 피보나치 수열 (타뷸레이션 - Bottom-Up)
//   - 시간 복잡도: O(N)
//   - 공간 복잡도: O(N) (DP 테이블), 최적화 시 O(1)
//
// 예시 3: 동전 교환 문제 — 경우의 수 (Coin Change - Number of Ways)
//   - 시간 복잡도: O(N × M) (N: 목표 금액, M: 동전 종류 수)
//   - 공간 복잡도: O(N) (DP 테이블)
//
// 예시 4: LCS (최장 공통 부분 수열, Longest Common Subsequence)
//   - 시간 복잡도: O(N × M) (N: 문자열 a 길이, M: 문자열 b 길이)
//   - 공간 복잡도: O(N × M) (2차원 DP 테이블)
//   - 특징: 두 문자열에서 순서를 유지하면서 공통으로 나타나는 가장 긴 부분 수열의 길이를 구한다
//
// 예시 5: 0/1 배낭 문제 (0/1 Knapsack)
//   - 시간 복잡도: O(N × W) (N: 물건 수, W: 배낭 용량)
//   - 공간 복잡도: O(N × W) (2차원 DP 테이블), 1차원 최적화 시 O(W)
//   - 특징: 각 물건을 넣거나(1) 넣지 않거나(0) 선택하여 최대 가치를 구한다

// fibMemo 함수는 메모이제이션(Top-Down) 방식으로 n번째 피보나치 수를 구한다.
// 재귀 호출 시 이미 계산한 값은 memo 맵에서 바로 꺼내 중복 계산을 방지한다.
func fibMemo(n int, memo map[int]int) int {
	// 기저 조건: F(0) = 0, F(1) = 1
	if n <= 1 {
		return n
	}

	// 이미 계산한 값이 있으면 바로 반환 (메모이제이션 핵심)
	if val, ok := memo[n]; ok {
		return val
	}

	// 아직 계산하지 않은 경우: 재귀로 계산 후 저장
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

// fibTab 함수는 타뷸레이션(Bottom-Up) 방식으로 n번째 피보나치 수를 구한다.
// 작은 값부터 차례로 테이블을 채워 올라가므로 재귀 호출이 필요 없다.
func fibTab(n int) int {
	// 기저 조건 처리
	if n <= 1 {
		return n
	}

	// DP 테이블 생성 및 초기값 설정
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	// 아래에서 위로 테이블을 채워나간다
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// fibTabOptimized 함수는 공간 최적화된 타뷸레이션 방식으로 피보나치 수를 구한다.
// 직전 두 값만 필요하므로 변수 2개로 O(1) 공간에 해결할 수 있다.
func fibTabOptimized(n int) int {
	if n <= 1 {
		return n
	}

	// 직전 두 값만 유지
	prev2, prev1 := 0, 1

	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

// coinChangeWays 함수는 주어진 동전으로 목표 금액을 만드는 경우의 수를 구한다.
// 타뷸레이션(Bottom-Up) 방식의 DP를 사용한다.
// coins: 사용 가능한 동전 종류, amount: 목표 금액
// 반환값: 목표 금액을 만드는 서로 다른 조합의 수
func coinChangeWays(coins []int, amount int) int {
	// dp[i] = 금액 i를 만드는 경우의 수
	dp := make([]int, amount+1)

	// 금액 0을 만드는 방법은 아무 동전도 사용하지 않는 1가지
	dp[0] = 1

	// 각 동전에 대해 가능한 금액을 갱신한다
	// 동전을 바깥 루프에 두면 같은 조합을 중복 없이 셀 수 있다
	// 예: {1,2}로 3을 만들 때 (1+2)와 (2+1)을 하나로 센다
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

// lcs 함수는 두 문자열의 최장 공통 부분 수열(LCS)의 길이를 구한다.
// 2차원 DP 테이블을 사용하는 대표적인 패턴이다.
// dp[i][j] = a[:i]와 b[:j]의 LCS 길이
// 점화식:
//   - a[i-1] == b[j-1] 이면: dp[i][j] = dp[i-1][j-1] + 1
//   - 다르면: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
func lcs(a, b string) int {
	n, m := len(a), len(b)

	// 2차원 DP 테이블 생성 (0으로 초기화)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// 테이블 채우기
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				// 두 문자가 같으면 대각선 값 + 1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 다르면 위쪽과 왼쪽 중 큰 값을 선택
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[n][m]
}

// knapsack01 함수는 0/1 배낭 문제를 2차원 DP로 풀어 최대 가치를 구한다.
// weights: 각 물건의 무게, values: 각 물건의 가치, capacity: 배낭 용량
// dp[i][w] = 처음 i개의 물건으로 용량 w를 채울 때의 최대 가치
// 점화식:
//   - weights[i-1] > w 이면: dp[i][w] = dp[i-1][w] (물건을 넣을 수 없음)
//   - 아니면: dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]] + values[i-1])
func knapsack01(weights, values []int, capacity int) int {
	n := len(weights)

	// 2차원 DP 테이블 생성
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// 테이블 채우기
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			// 기본: 현재 물건을 넣지 않는 경우
			dp[i][w] = dp[i-1][w]

			// 현재 물건을 넣을 수 있고, 넣는 것이 더 이득인 경우
			if weights[i-1] <= w {
				withItem := dp[i-1][w-weights[i-1]] + values[i-1]
				if withItem > dp[i][w] {
					dp[i][w] = withItem
				}
			}
		}
	}

	return dp[n][capacity]
}

func main() {
	// === 피보나치 수열: 메모이제이션 (Top-Down) ===
	fmt.Println("=== 피보나치 수열: 메모이제이션 (Top-Down) ===")
	fmt.Println()

	memo := make(map[int]int)
	n := 10
	fmt.Printf("F(%d) = %d\n", n, fibMemo(n, memo))
	fmt.Println("메모 테이블 내용:")
	for i := 0; i <= n; i++ {
		if val, ok := memo[i]; ok {
			fmt.Printf("  F(%d) = %d\n", i, val)
		}
	}

	// === 피보나치 수열: 타뷸레이션 (Bottom-Up) ===
	fmt.Println("\n=== 피보나치 수열: 타뷸레이션 (Bottom-Up) ===")
	fmt.Println()

	fmt.Println("일반 타뷸레이션:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  F(%d) = %d\n", i, fibTab(i))
	}

	fmt.Println("\n공간 최적화 타뷸레이션:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  F(%d) = %d\n", i, fibTabOptimized(i))
	}

	// === 동전 교환 문제 — 경우의 수 ===
	fmt.Println("\n=== 동전 교환 문제 — 경우의 수 (Coin Change) ===")
	fmt.Println()

	coins := []int{1, 2, 3}
	amount := 4

	fmt.Printf("동전 종류: %v\n", coins)
	fmt.Printf("목표 금액: %d\n", amount)
	fmt.Printf("경우의 수: %d\n", coinChangeWays(coins, amount))
	fmt.Println()
	fmt.Println("가능한 조합:")
	fmt.Println("  {1, 1, 1, 1}")
	fmt.Println("  {1, 1, 2}")
	fmt.Println("  {1, 3}")
	fmt.Println("  {2, 2}")

	// 추가 예시: 더 큰 금액
	fmt.Println()
	coins2 := []int{1, 5, 10, 25}
	amount2 := 50

	fmt.Printf("동전 종류: %v\n", coins2)
	fmt.Printf("목표 금액: %d\n", amount2)
	fmt.Printf("경우의 수: %d\n", coinChangeWays(coins2, amount2))

	// === LCS (최장 공통 부분 수열) ===
	fmt.Println("\n=== LCS (최장 공통 부분 수열) ===")
	fmt.Println()

	s1, s2 := "ABCBDAB", "BDCAB"
	fmt.Printf("문자열 A: %q\n", s1)
	fmt.Printf("문자열 B: %q\n", s2)
	fmt.Printf("LCS 길이: %d\n", lcs(s1, s2))
	fmt.Println("설명: 공통 부분 수열 예시 → \"BCAB\" (길이 4)")

	fmt.Println()
	s3, s4 := "AGGTAB", "GXTXAYB"
	fmt.Printf("문자열 A: %q\n", s3)
	fmt.Printf("문자열 B: %q\n", s4)
	fmt.Printf("LCS 길이: %d\n", lcs(s3, s4))
	fmt.Println("설명: 공통 부분 수열 예시 → \"GTAB\" (길이 4)")

	// === 0/1 배낭 문제 ===
	fmt.Println("\n=== 0/1 배낭 문제 (Knapsack) ===")
	fmt.Println()

	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 8

	fmt.Printf("물건 (무게, 가치): ")
	for i := range weights {
		fmt.Printf("(%d, %d) ", weights[i], values[i])
	}
	fmt.Println()
	fmt.Printf("배낭 용량: %d\n", capacity)
	fmt.Printf("최대 가치: %d\n", knapsack01(weights, values, capacity))
	fmt.Println("설명: 물건 0(무게2,가치3) + 물건 1(무게3,가치4) + 물건 2(무게4,가치5) 또는")
	fmt.Println("       물건 1(무게3,가치4) + 물건 3(무게5,가치6) = 10")
}
