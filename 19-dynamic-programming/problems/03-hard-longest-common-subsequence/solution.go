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

	// 입력 처리
	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	n := len(a)
	m := len(b)

	// dp[i][j] = a의 처음 i글자와 b의 처음 j글자의 LCS 길이
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 바텀업으로 DP 테이블을 채운다
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				// 두 문자가 같으면 LCS에 포함한다
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 다르면 한쪽을 제외한 경우 중 큰 값을 선택한다
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	// LCS 길이 출력
	lcsLen := dp[n][m]
	fmt.Fprintln(writer, lcsLen)

	// LCS 길이가 0이면 부분 수열을 출력하지 않는다
	if lcsLen == 0 {
		return
	}

	// 역추적으로 LCS 문자열을 복원한다
	lcs := make([]byte, lcsLen)
	idx := lcsLen - 1
	i, j := n, m
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			// 두 문자가 같으면 LCS에 포함된 문자이다
			lcs[idx] = a[i-1]
			idx--
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			// 위쪽 값이 더 크면 위로 이동한다
			i--
		} else {
			// 왼쪽 값이 더 크거나 같으면 왼쪽으로 이동한다
			j--
		}
	}

	fmt.Fprintln(writer, string(lcs))
}
