package main

import "fmt"

// 탑다운(메모이제이션) 방식의 동적 프로그래밍 예시 - 피보나치 수
// 시간 복잡도: O(N)
// 공간 복잡도: O(N)

// memo 테이블: 이미 계산한 피보나치 값을 저장한다
var memo map[int]int

// fibonacci 함수는 재귀 + 메모이제이션으로 N번째 피보나치 수를 구한다
func fibonacci(n int) int {
	// 기저 사례: F(0) = 0, F(1) = 1
	if n <= 1 {
		return n
	}

	// 이미 계산한 값이 있으면 바로 반환한다
	if val, ok := memo[n]; ok {
		return val
	}

	// 점화식: F(n) = F(n-1) + F(n-2)
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func main() {
	memo = make(map[int]int)

	// 0번째부터 10번째 피보나치 수를 출력한다
	fmt.Println("=== 탑다운(메모이제이션) 피보나치 ===")
	for i := 0; i <= 10; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}

	// 큰 값도 중복 계산 없이 빠르게 구할 수 있다
	fmt.Printf("\nF(40) = %d\n", fibonacci(40))
}
