package main

import (
	"bufio"
	"fmt"
	"os"
)

// minCostSequence는 수열을 분할할 때 각 구간 합의 제곱의 합을 최소화하는
// 비용을 볼록 껍질 트릭(CHT)으로 구한다.
//
// [매개변수]
//   - n: 수열의 길이
//   - a: 수열 (1-indexed)
//
// [반환값]
//   - int64: 최소 분할 비용
func minCostSequence(n int, a []int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	fmt.Fprintln(writer, minCostSequence(n, a))
}
