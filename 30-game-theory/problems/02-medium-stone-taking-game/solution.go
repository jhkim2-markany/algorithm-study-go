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

	// 돌의 개수 N과 가져갈 수 있는 경우의 수 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 가져갈 수 있는 돌의 개수 집합 입력
	moves := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &moves[i])
	}

	// dp[i] = true이면 돌 i개인 상태에서 선수 승리
	dp := make([]bool, n+1)
	// 돌 0개: 수를 둘 수 없으므로 패배 (false)

	for i := 1; i <= n; i++ {
		// 가능한 모든 이동을 시도한다
		for _, m := range moves {
			if i >= m && !dp[i-m] {
				// 상대를 패배 포지션으로 보낼 수 있으면 승리
				dp[i] = true
				break
			}
		}
	}

	// 돌 N개인 초기 상태의 승패를 출력한다
	if dp[n] {
		fmt.Fprintln(writer, "First")
	} else {
		fmt.Fprintln(writer, "Second")
	}
}
