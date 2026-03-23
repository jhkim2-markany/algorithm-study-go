package main

import (
	"bufio"
	"fmt"
	"os"
)

// 최소 비용 수열 분할
// dp[i] = min(dp[j] + (sum[i] - sum[j])²)  (0 ≤ j < i)
// 전개: dp[i] = sum[i]² + min((-2·sum[j])·sum[i] + dp[j] + sum[j]²)
// 일차 함수: f_j(x) = (-2·sum[j])·x + (dp[j] + sum[j]²), 쿼리 x = sum[i]
//
// sum[i]가 단조 증가하므로 쿼리 x가 단조 증가한다.
// 기울기 -2·sum[j]는 j가 증가하면 단조 감소한다.
// → 단조 CHT로 O(N)에 해결 가능

// Line은 일차 함수 y = m*x + b를 나타낸다
type Line struct {
	m, b int64
}

func (l Line) eval(x int64) int64 {
	return l.m*x + l.b
}

var (
	lines []Line // 볼록 껍질
	ptr   int    // 단조 쿼리 포인터
)

// bad는 직선 b가 a와 c 사이에서 불필요한지 판정한다
func bad(a, b, c Line) bool {
	return (c.b-a.b)*(a.m-b.m) <= (b.b-a.b)*(a.m-c.m)
}

// addLine은 기울기 감소 순서로 직선을 추가한다
func addLine(m, b int64) {
	newLine := Line{m, b}
	for len(lines) >= 2 {
		n := len(lines)
		if bad(lines[n-2], lines[n-1], newLine) {
			lines = lines[:n-1]
		} else {
			break
		}
	}
	lines = append(lines, newLine)
}

// query는 단조 증가하는 x에 대해 최솟값을 반환한다
func query(x int64) int64 {
	for ptr+1 < len(lines) && lines[ptr+1].eval(x) <= lines[ptr].eval(x) {
		ptr++
	}
	return lines[ptr].eval(x)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 수열 길이
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 누적합 계산
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + a[i]
	}

	// DP 계산
	dp := make([]int64, n+1)
	lines = nil
	ptr = 0

	// 초기 직선: j=0일 때 f₀(x) = 0·x + 0
	addLine(-2*sum[0], dp[0]+sum[0]*sum[0])

	for i := 1; i <= n; i++ {
		// 쿼리: x = sum[i]에서 최솟값
		minVal := query(sum[i])
		dp[i] = sum[i]*sum[i] + minVal

		// 직선 추가: f_i(x) = (-2·sum[i])·x + (dp[i] + sum[i]²)
		addLine(-2*sum[i], dp[i]+sum[i]*sum[i])
	}

	// 출력: 최소 비용
	fmt.Fprintln(writer, dp[n])
}
