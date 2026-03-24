package main

import (
	"bufio"
	"fmt"
	"os"
)

// abs는 정수의 절댓값을 반환한다.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// max는 두 정수 중 큰 값을 반환한다.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// sherlockAndCost는 인접 원소 간 절댓값 차이의 합의 최댓값을 반환한다.
//
// [매개변수]
//   - b: 각 위치의 상한값 배열
//
// [반환값]
//   - int: S의 최댓값
//
// [알고리즘 힌트]
//
//	각 위치에서 A[i]를 1 또는 B[i]로 선택하는 두 가지 상태를 유지한다.
//	이전 상태로부터 전이하여 최대 차이 합을 구한다.
func sherlockAndCost(b []int) int {
	// low: 현재 위치에서 A[i] = 1을 선택했을 때의 최대 S
	// high: 현재 위치에서 A[i] = B[i]를 선택했을 때의 최대 S
	low := 0
	high := 0

	for i := 1; i < len(b); i++ {
		// 이전 상태에서 현재 상태로 전이
		newLow := max(
			low,                // 이전도 1, 현재도 1: 차이 0
			high+abs(1-b[i-1]), // 이전 B[i-1], 현재 1: 차이 |1-B[i-1]|
		)
		newHigh := max(
			low+abs(b[i]-1),       // 이전 1, 현재 B[i]: 차이 |B[i]-1|
			high+abs(b[i]-b[i-1]), // 이전 B[i-1], 현재 B[i]: 차이 |B[i]-B[i-1]|
		)

		low = newLow
		high = newHigh
	}

	return max(low, high)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		// 배열 크기 입력
		var n int
		fmt.Fscan(reader, &n)

		// 배열 B 입력
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := sherlockAndCost(b)
		fmt.Fprintln(writer, result)
	}
}
