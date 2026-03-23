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
//
// [알고리즘 힌트]
//   1. dp[i] = min(dp[j] + (sum[i]-sum[j])²)로 점화식을 세운다
//   2. 전개하면 dp[i] = sum[i]² + min((-2·sum[j])·sum[i] + dp[j] + sum[j]²)
//   3. f_j(x) = (-2·sum[j])·x + (dp[j]+sum[j]²)로 일차 함수를 정의한다
//   4. sum[i]가 단조 증가하므로 단조 CHT로 O(N)에 해결한다
//   5. 기울기 감소 순으로 직선을 추가하고, 포인터를 이동하며 최솟값을 쿼리한다
func minCostSequence(n int, a []int64) int64 {
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + a[i]
	}

	type Line struct {
		m, b int64
	}

	lines := []Line{}
	ptr := 0

	bad := func(a, b, c Line) bool {
		return (c.b-a.b)*(a.m-b.m) <= (b.b-a.b)*(a.m-c.m)
	}

	addLine := func(m, b int64) {
		newLine := Line{m, b}
		for len(lines) >= 2 {
			nn := len(lines)
			if bad(lines[nn-2], lines[nn-1], newLine) {
				lines = lines[:nn-1]
			} else {
				break
			}
		}
		lines = append(lines, newLine)
	}

	query := func(x int64) int64 {
		for ptr+1 < len(lines) && lines[ptr+1].m*x+lines[ptr+1].b <= lines[ptr].m*x+lines[ptr].b {
			ptr++
		}
		return lines[ptr].m*x + lines[ptr].b
	}

	dp := make([]int64, n+1)
	addLine(-2*sum[0], dp[0]+sum[0]*sum[0])

	for i := 1; i <= n; i++ {
		minVal := query(sum[i])
		dp[i] = sum[i]*sum[i] + minVal
		addLine(-2*sum[i], dp[i]+sum[i]*sum[i])
	}

	return dp[n]
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
