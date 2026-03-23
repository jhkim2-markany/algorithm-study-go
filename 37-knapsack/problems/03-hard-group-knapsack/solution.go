package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	// dp[w] = 용량 w 이하로 담을 때의 최대 가치
	dp := make([]int, k+1)

	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(reader, &m)

		// 그룹 내 물건 정보를 읽는다
		groupWeights := make([]int, m)
		groupValues := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &groupWeights[j], &groupValues[j])
		}

		// 역순으로 순회하여 각 그룹에서 최대 1개만 선택한다
		for w := k; w >= 0; w-- {
			// 그룹 내 각 물건을 시도한다
			for j := 0; j < m; j++ {
				if w >= groupWeights[j] {
					val := dp[w-groupWeights[j]] + groupValues[j]
					if val > dp[w] {
						dp[w] = val
					}
				}
			}
		}
	}

	// 최대 가치를 출력한다
	fmt.Fprintln(writer, dp[k])
}
