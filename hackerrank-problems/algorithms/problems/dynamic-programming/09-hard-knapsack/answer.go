package main

import (
	"bufio"
	"fmt"
	"os"
)

// unboundedKnapsack은 원소를 중복 사용하여 K 이하의 최대 합을 반환한다.
//
// [매개변수]
//   - k: 목표값 (상한)
//   - arr: 정수 배열
//
// [반환값]
//   - int: K를 넘지 않는 최대 합
//
// [알고리즘 힌트]
//
//	dp[j] = 합이 j를 만들 수 있는지 여부를 추적한다.
//	각 원소를 무한히 사용할 수 있으므로 동전 교환과 유사하게 처리한다.
func unboundedKnapsack(k int, arr []int) int {
	// dp[j] = 합이 정확히 j가 가능한지 여부
	dp := make([]bool, k+1)
	dp[0] = true

	// 각 원소에 대해 dp 갱신
	for _, val := range arr {
		for j := val; j <= k; j++ {
			if dp[j-val] {
				dp[j] = true
			}
		}
	}

	// K부터 역순으로 탐색하여 가능한 최대 합 찾기
	for j := k; j >= 0; j-- {
		if dp[j] {
			return j
		}
	}

	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		// 배열 크기와 목표값 입력
		var n, k int
		fmt.Fscan(reader, &n, &k)

		// 배열 입력
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := unboundedKnapsack(k, arr)
		fmt.Fprintln(writer, result)
	}
}
