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

	// 원소 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 원소 입력
	a := make([]int, n)
	totalSum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		totalSum += a[i]
	}

	// dp[mask]: mask에 해당하는 원소들의 합
	// 모든 부분집합의 합을 비트마스크로 계산한다
	full := 1 << n
	subsetSum := make([]int, full)

	// 각 부분집합의 합을 구한다
	for mask := 1; mask < full; mask++ {
		// mask에서 가장 낮은 비트를 찾는다
		lsb := mask & (-mask)
		bit := 0
		temp := lsb
		for temp > 1 {
			bit++
			temp >>= 1
		}
		// 이전 부분집합의 합에 새 원소를 더한다
		subsetSum[mask] = subsetSum[mask^lsb] + a[bit]
	}

	// 모든 부분집합을 순회하며 최소 차이를 구한다
	// 한쪽 부분집합의 합이 s이면 다른 쪽은 totalSum - s이다
	// 차이 = |totalSum - 2*s|
	ans := totalSum // 최대 차이로 초기화
	for mask := 1; mask < full-1; mask++ {
		// 공집합과 전체 집합은 제외한다 (두 부분집합 모두 비어있지 않아야 함)
		s := subsetSum[mask]
		diff := totalSum - 2*s
		if diff < 0 {
			diff = -diff
		}
		if diff < ans {
			ans = diff
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, ans)
}
