package main

import (
	"bufio"
	"fmt"
	"os"
)

// minPartitionDiff는 집합을 두 부분집합으로 나눌 때 합의 차이의 최솟값을 반환한다.
//
// [매개변수]
//   - n: 원소의 수
//   - a: 원소 배열 (길이 n)
//
// [반환값]
//   - int: 두 부분집합 합의 차이의 최솟값
//
// [알고리즘 힌트]
//
//	비트마스크로 모든 부분집합의 합을 구한 뒤 최소 차이를 찾는다.
//	subsetSum[mask] = mask에 해당하는 원소들의 합.
//	차이 = |totalSum - 2*s|.
//	시간복잡도: O(2^N)
func minPartitionDiff(n int, a []int) int {
	totalSum := 0
	for i := 0; i < n; i++ {
		totalSum += a[i]
	}

	full := 1 << n
	subsetSum := make([]int, full)

	for mask := 1; mask < full; mask++ {
		lsb := mask & (-mask)
		bit := 0
		temp := lsb
		for temp > 1 {
			bit++
			temp >>= 1
		}
		subsetSum[mask] = subsetSum[mask^lsb] + a[bit]
	}

	ans := totalSum
	for mask := 1; mask < full-1; mask++ {
		s := subsetSum[mask]
		diff := totalSum - 2*s
		if diff < 0 {
			diff = -diff
		}
		if diff < ans {
			ans = diff
		}
	}

	return ans
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

	fmt.Fprintln(writer, minPartitionDiff(n, a))
}
