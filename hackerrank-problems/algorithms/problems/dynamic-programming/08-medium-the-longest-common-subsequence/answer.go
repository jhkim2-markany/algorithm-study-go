package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestCommonSubsequence는 두 수열의 최장 공통 부분 수열을 반환한다.
//
// [매개변수]
//   - a: 첫 번째 수열
//   - b: 두 번째 수열
//
// [반환값]
//   - []int: 최장 공통 부분 수열
//
// [알고리즘 힌트]
//
//	2차원 DP 테이블로 LCS 길이를 구한 뒤,
//	역추적으로 실제 부분 수열을 복원한다.
func longestCommonSubsequence(a, b []int) []int {
	n, m := len(a), len(b)

	// dp[i][j] = a[:i]와 b[:j]의 LCS 길이
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// DP 테이블 채우기
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				// 같은 원소: LCS에 포함
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 다른 원소: 한쪽을 줄인 경우 중 최댓값
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	// 역추적으로 LCS 복원
	lcsLen := dp[n][m]
	lcs := make([]int, lcsLen)
	idx := lcsLen - 1
	i, j := n, m

	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			// 같은 원소: LCS에 추가
			lcs[idx] = a[i-1]
			idx--
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			// 위쪽으로 이동
			i--
		} else {
			// 왼쪽으로 이동
			j--
		}
	}

	return lcs
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 수열 길이 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 첫 번째 수열 입력
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 두 번째 수열 입력
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := longestCommonSubsequence(a, b)
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
