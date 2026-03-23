package main

import (
	"bufio"
	"fmt"
	"os"
)

// collectAllExpected는 n종류의 아이템을 모두 모으기 위한 기대 횟수를 반환한다.
//
// [매개변수]
//   - n: 아이템 종류의 수 (1 이상, 최대 20)
//   - prob: 각 아이템이 뽑힐 확률 배열 (길이 n, 합이 1.0)
//
// [반환값]
//   - float64: 모든 아이템을 모으기까지의 기대 횟수
//
// [알고리즘 힌트]
//
//	비트마스크 DP를 사용한다.
//	dp[mask] = 이미 모은 아이템 집합이 mask일 때 전부 모으기까지의 기대 횟수.
//	dp[full] = 0, 역순으로 계산한다.
//	dp[mask] = (1 + Σ prob[i]*dp[mask|(1<<i)]) / pNew
//	시간복잡도: O(2^N * N)
func collectAllExpected(n int, prob []float64) float64 {
	full := (1 << n) - 1
	dp := make([]float64, 1<<n)

	for mask := full - 1; mask >= 0; mask-- {
		pRepeat := 0.0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				pRepeat += prob[i]
			}
		}

		pNew := 1.0 - pRepeat

		sum := 0.0
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				sum += prob[i] * dp[mask|(1<<i)]
			}
		}

		dp[mask] = (1.0 + sum) / pNew
	}

	return dp[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	prob := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &prob[i])
	}

	fmt.Fprintf(writer, "%.6f\n", collectAllExpected(n, prob))
}
