package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcd: 최대공약수를 구한다 (기약분수 변환용)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 주사위 횟수 N과 목표 합 S 입력
	var n, s int
	fmt.Fscan(reader, &n, &s)

	// dp[i][j] = 주사위를 i번 던져서 합이 j가 되는 경우의 수
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, s+1)
	}
	// 초기 상태: 0번 던져서 합이 0인 경우는 1가지
	dp[0][0] = 1

	// 각 주사위를 던질 때 1~6의 눈이 나올 수 있다
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			for face := 1; face <= 6; face++ {
				if j >= face {
					dp[i][j] += dp[i-1][j-face]
				}
			}
		}
	}

	// 분자: 합이 S가 되는 경우의 수
	numerator := dp[n][s]
	// 분모: 전체 경우의 수 = 6^N
	denominator := 1
	for i := 0; i < n; i++ {
		denominator *= 6
	}

	// 기약분수로 변환하여 출력
	if numerator == 0 {
		fmt.Fprintln(writer, "0/1")
	} else {
		g := gcd(numerator, denominator)
		fmt.Fprintf(writer, "%d/%d\n", numerator/g, denominator/g)
	}
}
