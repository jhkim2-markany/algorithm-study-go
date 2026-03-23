package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcd는 두 정수의 최대공약수를 구한다.
//
// [매개변수]
//   - a: 첫 번째 양의 정수
//   - b: 두 번째 양의 정수
//
// [반환값]
//   - int: a와 b의 최대공약수
//
// [알고리즘 힌트]
//
//	유클리드 호제법을 사용한다.
//	시간복잡도: O(log(min(a, b)))
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// diceProbability는 주사위를 n번 던져서 합이 s가 되는 확률을 기약분수로 반환한다.
//
// [매개변수]
//   - n: 주사위를 던지는 횟수 (1 이상)
//   - s: 목표 합
//
// [반환값]
//   - int: 기약분수의 분자
//   - int: 기약분수의 분모
//
// [알고리즘 힌트]
//
//	DP로 경우의 수를 구한 뒤 6^N으로 나눠 기약분수를 만든다.
//	dp[i][j] = 주사위 i번 던져서 합이 j인 경우의 수.
//	시간복잡도: O(N * S * 6)
func diceProbability(n, s int) (int, int) {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, s+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			for face := 1; face <= 6; face++ {
				if j >= face {
					dp[i][j] += dp[i-1][j-face]
				}
			}
		}
	}

	numerator := dp[n][s]
	denominator := 1
	for i := 0; i < n; i++ {
		denominator *= 6
	}

	if numerator == 0 {
		return 0, 1
	}
	g := gcd(numerator, denominator)
	return numerator / g, denominator / g
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscan(reader, &n, &s)

	num, den := diceProbability(n, s)
	fmt.Fprintf(writer, "%d/%d\n", num, den)
}
