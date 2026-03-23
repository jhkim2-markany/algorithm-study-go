package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 수열 길이
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// O(N²) DP로 LIS 길이와 개수를 동시에 구한다
	dp := make([]int, n)  // dp[i] = a[i]를 마지막으로 하는 LIS 길이
	cnt := make([]int, n) // cnt[i] = a[i]를 마지막으로 하는 LIS 개수

	lisLen := 1

	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1

		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				if dp[j]+1 > dp[i] {
					// 더 긴 LIS 발견: 길이 갱신, 개수 초기화
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					// 같은 길이의 LIS: 개수 누적
					cnt[i] = (cnt[i] + cnt[j]) % mod
				}
			}
		}

		if dp[i] > lisLen {
			lisLen = dp[i]
		}
	}

	// LIS 길이가 lisLen인 모든 위치의 개수를 합산한다
	totalCount := 0
	for i := 0; i < n; i++ {
		if dp[i] == lisLen {
			totalCount = (totalCount + cnt[i]) % mod
		}
	}

	// 출력: LIS 길이와 개수
	fmt.Fprintln(writer, lisLen)
	fmt.Fprintln(writer, totalCount)
}
