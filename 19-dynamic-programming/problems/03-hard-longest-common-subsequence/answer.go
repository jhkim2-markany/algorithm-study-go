package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestCommonSubsequence는 두 문자열의 LCS 길이와 LCS 문자열을 반환한다.
//
// [매개변수]
//   - a: 첫 번째 문자열
//   - b: 두 번째 문자열
//
// [반환값]
//   - int: LCS의 길이
//   - string: LCS 문자열 (길이가 0이면 빈 문자열)
//
// [알고리즘 힌트]
//
//	2차원 DP 테이블을 사용한다.
//	dp[i][j] = a의 처음 i글자와 b의 처음 j글자의 LCS 길이.
//	두 문자가 같으면 dp[i][j] = dp[i-1][j-1] + 1,
//	다르면 dp[i][j] = max(dp[i-1][j], dp[i][j-1]).
//	역추적으로 LCS 문자열을 복원한다.
func longestCommonSubsequence(a, b string) (int, string) {
	n := len(a)
	m := len(b)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	lcsLen := dp[n][m]
	if lcsLen == 0 {
		return 0, ""
	}

	lcs := make([]byte, lcsLen)
	idx := lcsLen - 1
	i, j := n, m
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			lcs[idx] = a[i-1]
			idx--
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcsLen, string(lcs)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	// 핵심 함수 호출
	lcsLen, lcsStr := longestCommonSubsequence(a, b)

	// LCS 길이 출력
	fmt.Fprintln(writer, lcsLen)

	// LCS 문자열 출력 (길이가 0이 아닌 경우)
	if lcsLen > 0 {
		fmt.Fprintln(writer, lcsStr)
	}
}
