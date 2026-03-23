package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 유량 네트워크의 간선을 나타낸다.
type Edge struct {
	to, cap, flow, rev int
}

// maxProfit은 의존 관계를 만족하면서 프로젝트를 선택했을 때의 최대 순이익을 반환한다.
//
// [매개변수]
//   - n: 프로젝트의 수
//   - profit: 각 프로젝트의 이익 배열 (1-indexed, 음수이면 비용)
//   - deps: 의존 관계 목록 (각 원소는 [a, b], a가 b에 의존)
//
// [반환값]
//   - int: 최대 순이익
func maxProfit(n int, profit []int, deps [][2]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	profit := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &profit[i])
	}

	deps := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &deps[i][0], &deps[i][1])
	}

	fmt.Fprintln(writer, maxProfit(n, profit, deps))
}
