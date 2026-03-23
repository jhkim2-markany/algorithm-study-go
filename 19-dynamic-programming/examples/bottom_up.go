package main

import "fmt"

// 바텀업(타뷸레이션) 방식의 동적 프로그래밍 예시 - 피보나치 수
// 시간 복잡도: O(N)
// 공간 복잡도: O(N), 최적화 시 O(1)

// fibonacciTable 함수는 반복문으로 DP 테이블을 채워 N번째 피보나치 수를 구한다
func fibonacciTable(n int) int {
	if n <= 1 {
		return n
	}

	// DP 테이블 초기화
	dp := make([]int, n+1)

	// 기저 사례 설정
	dp[0] = 0
	dp[1] = 1

	// 작은 문제부터 큰 문제 방향으로 테이블을 채운다
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// fibonacciOptimized 함수는 공간을 O(1)로 최적화한 바텀업 방식이다
func fibonacciOptimized(n int) int {
	if n <= 1 {
		return n
	}

	// 이전 두 값만 유지하면 충분하다
	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

func main() {
	// 기본 바텀업 방식
	fmt.Println("=== 바텀업(타뷸레이션) 피보나치 ===")
	for i := 0; i <= 10; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacciTable(i))
	}

	// 공간 최적화 버전
	fmt.Println("\n=== 공간 최적화 바텀업 피보나치 ===")
	fmt.Printf("F(40) = %d\n", fibonacciOptimized(40))
	fmt.Printf("F(50) = %d\n", fibonacciOptimized(50))
}
