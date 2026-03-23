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

	weights := make([]int, n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &weights[i], &values[i])
	}

	// dp[w] = 용량 w 이하로 담을 때의 최대 가치
	dp := make([]int, k+1)

	for i := 0; i < n; i++ {
		// 정순으로 순회하여 같은 물건을 여러 번 선택할 수 있게 한다
		for w := weights[i]; w <= k; w++ {
			if dp[w-weights[i]]+values[i] > dp[w] {
				dp[w] = dp[w-weights[i]] + values[i]
			}
		}
	}

	// 최대 가치를 출력한다
	fmt.Fprintln(writer, dp[k])
}
