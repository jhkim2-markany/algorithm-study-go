package main

import (
	"bufio"
	"fmt"
	"os"
)

// minCostAssignment는 n명의 작업자에게 n개의 작업을 1:1 배정할 때 최소 비용을 반환한다.
//
// [매개변수]
//   - n: 작업자/작업의 수
//   - cost: n×n 비용 행렬 (cost[i][j] = 작업자 i가 작업 j를 수행하는 비용)
//
// [반환값]
//   - int: 모든 작업을 배정했을 때의 최소 총 비용
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
