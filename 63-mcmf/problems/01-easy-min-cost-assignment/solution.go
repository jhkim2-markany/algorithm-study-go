package main

import (
	"bufio"
	"fmt"
	"os"
)

// minCostAssignment은 N명의 작업자를 N개의 작업에 1:1 배정할 때
// MCMF(최소 비용 최대 유량)를 이용하여 최소 비용을 구한다.
//
// [매개변수]
//   - n: 작업자/작업 수
//   - cost: n×n 비용 행렬 (cost[i][j] = 작업자 i가 작업 j를 수행하는 비용)
//
// [반환값]
//   - int: 최소 배정 비용
func minCostAssignment(n int, cost [][]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	fmt.Fprintln(writer, minCostAssignment(n, cost))
}
