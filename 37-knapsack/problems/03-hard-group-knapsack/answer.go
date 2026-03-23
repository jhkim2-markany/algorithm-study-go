package main

import (
	"bufio"
	"fmt"
	"os"
)

// Group은 그룹 내 물건들의 무게와 가치를 저장한다.
type Group struct {
	weights []int
	values  []int
}

// groupKnapsack은 그룹 배낭 문제의 최대 가치를 반환한다.
//
// [매개변수]
//   - n: 그룹의 수
//   - k: 배낭의 용량
//   - groups: 각 그룹의 물건 정보 배열 (길이 n)
//
// [반환값]
//   - int: 각 그룹에서 최대 1개씩 선택하여 담을 수 있는 최대 가치
//
// [알고리즘 힌트]
//
//	1차원 DP + 그룹별 역순 순회를 사용한다.
//	각 그룹에 대해 역순으로 순회하며 그룹 내 물건을 시도한다.
//	역순이므로 각 그룹에서 최대 1개만 선택된다.
//	시간복잡도: O(N * K * M_max), 공간복잡도: O(K)
func groupKnapsack(n, k int, groups []Group) int {
	dp := make([]int, k+1)

	for i := 0; i < n; i++ {
		m := len(groups[i].weights)
		for w := k; w >= 0; w-- {
			for j := 0; j < m; j++ {
				if w >= groups[i].weights[j] {
					val := dp[w-groups[i].weights[j]] + groups[i].values[j]
					if val > dp[w] {
						dp[w] = val
					}
				}
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

	groups := make([]Group, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(reader, &m)
		groups[i].weights = make([]int, m)
		groups[i].values = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &groups[i].weights[j], &groups[i].values[j])
		}
	}

	fmt.Fprintln(writer, groupKnapsack(n, k, groups))
}
