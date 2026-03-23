package main

import (
	"bufio"
	"fmt"
	"os"
)

// unboundedKnapsack은 무한 배낭 문제의 최대 가치를 반환한다.
//
// [매개변수]
//   - n: 물건 종류의 수
//   - k: 배낭의 용량
//   - weights: 각 물건의 무게 배열 (길이 n)
//   - values: 각 물건의 가치 배열 (길이 n)
//
// [반환값]
//   - int: 용량 k 이하로 담을 수 있는 최대 가치 (같은 물건 여러 번 선택 가능)
//
// [알고리즘 힌트]
//
//	1차원 DP 배열을 사용한다.
//	정순으로 순회하여 같은 물건을 여러 번 선택할 수 있게 한다.
//	dp[w] = 용량 w 이하로 담을 때의 최대 가치.
//	시간복잡도: O(N * K), 공간복잡도: O(K)
func unboundedKnapsack(n, k int, weights, values []int) int {
	dp := make([]int, k+1)

	for i := 0; i < n; i++ {
		for w := weights[i]; w <= k; w++ {
			if dp[w-weights[i]]+values[i] > dp[w] {
				dp[w] = dp[w-weights[i]] + values[i]
			}
		}
	}

	return dp[k]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	weights := make([]int, n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &weights[i], &values[i])
	}

	fmt.Fprintln(writer, unboundedKnapsack(n, k, weights, values))
}
