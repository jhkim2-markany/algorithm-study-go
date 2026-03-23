package main

import (
	"bufio"
	"fmt"
	"os"
)

// expectedMoves는 1번 칸에서 n번 칸까지 도착하기 위한 기대 주사위 횟수를 반환한다.
//
// [매개변수]
//   - n: 보드의 마지막 칸 번호 (2 이상)
//
// [반환값]
//   - float64: 1번 칸에서 출발하여 n번 칸에 도착하기까지의 기대 횟수
//
// [알고리즘 힌트]
//
//	역순 DP로 기댓값을 계산한다.
//	dp[i] = i번 칸에서 N번 칸까지의 기대 횟수.
//	N을 초과하면 제자리에 머무는 조건을 처리한다.
//	dp[i] = (6 + sum) / (6 - stay)
//	시간복잡도: O(N)
func expectedMoves(n int) float64 {
	dp := make([]float64, n+1)

	for i := n - 1; i >= 1; i-- {
		sum := 0.0
		stay := 0

		for face := 1; face <= 6; face++ {
			next := i + face
			if next == n {
				sum += 0
			} else if next > n {
				stay++
			} else {
				sum += dp[next]
			}
		}

		dp[i] = (6.0 + sum) / float64(6-stay)
	}

	return dp[1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	fmt.Fprintf(writer, "%.6f\n", expectedMoves(n))
}
