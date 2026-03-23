package main

import (
	"bufio"
	"fmt"
	"os"
)

// stoneTakingWinner는 돌 가져가기 게임의 승자를 판별한다.
//
// [매개변수]
//   - n: 돌의 개수
//   - moves: 한 번에 가져갈 수 있는 돌의 개수 집합
//
// [반환값]
//   - string: 선수 승리이면 "First", 후수 승리이면 "Second"
//
// [알고리즘 힌트]
//
//	게임 DP를 사용한다.
//	dp[i] = true이면 돌 i개인 상태에서 선수 승리.
//	돌 0개는 수를 둘 수 없으므로 패배(false).
//	각 상태에서 가능한 모든 이동을 시도하여,
//	상대를 패배 포지션으로 보낼 수 있으면 승리이다.
func stoneTakingWinner(n int, moves []int) string {
	dp := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		for _, m := range moves {
			if i >= m && !dp[i-m] {
				dp[i] = true
				break
			}
		}
	}

	if dp[n] {
		return "First"
	}
	return "Second"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	moves := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &moves[i])
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, stoneTakingWinner(n, moves))
}
