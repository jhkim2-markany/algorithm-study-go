package main

import (
	"bufio"
	"fmt"
	"os"
)

// 팩토리얼을 재귀로 계산
func factorial(n int) int {
	// 기저 조건: 0! = 1, 1! = 1
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 피보나치를 메모이제이션 재귀로 계산
func fibonacci(n int, memo map[int]int) int {
	// 캐시 확인
	if val, ok := memo[n]; ok {
		return val
	}
	// 기저 조건
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 재귀 호출 후 결과 저장
	result := fibonacci(n-1, memo) + fibonacci(n-2, memo)
	memo[n] = result
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: N과 M
	var n, m int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)

	// 팩토리얼 계산 및 출력
	fmt.Fprintln(writer, factorial(n))

	// 피보나치 계산 및 출력
	memo := make(map[int]int)
	fmt.Fprintln(writer, fibonacci(m, memo))
}
