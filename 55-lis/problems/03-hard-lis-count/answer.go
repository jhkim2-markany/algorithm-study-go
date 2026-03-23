package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

// lisCount는 최장 증가 부분 수열(LIS)의 길이와 개수를 반환한다.
//
// [매개변수]
//   - a: 정수 수열
//
// [반환값]
//   - int: LIS의 길이
//   - int: LIS의 개수 (mod 1,000,000,007)
//
// [알고리즘 힌트]
//   - O(N²) DP로 LIS 길이와 개수를 동시에 구한다
//   - dp[i] = a[i]를 마지막으로 하는 LIS 길이, cnt[i] = 해당 LIS 개수
//   - a[j] < a[i]이고 dp[j]+1 > dp[i]이면 더 긴 LIS 발견 → 길이 갱신, 개수 초기화
//   - a[j] < a[i]이고 dp[j]+1 == dp[i]이면 같은 길이 → 개수 누적
//   - 최종적으로 dp[i] == lisLen인 모든 위치의 cnt를 합산한다
func lisCount(a []int) (int, int) {
	n := len(a)
	dp := make([]int, n)
	cnt := make([]int, n)
	lisLen := 1

	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1

		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] = (cnt[i] + cnt[j]) % mod
				}
			}
		}

		if dp[i] > lisLen {
			lisLen = dp[i]
		}
	}

	totalCount := 0
	for i := 0; i < n; i++ {
		if dp[i] == lisLen {
			totalCount = (totalCount + cnt[i]) % mod
		}
	}

	return lisLen, totalCount
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	lisLen, totalCount := lisCount(a)

	fmt.Fprintln(writer, lisLen)
	fmt.Fprintln(writer, totalCount)
}
