package main

import (
	"bufio"
	"fmt"
	"os"
)

// minHamiltonianPath는 모든 도시를 정확히 한 번 방문하는 최소 비용 경로를 반환한다.
//
// [매개변수]
//   - n: 도시의 수
//   - cost: n×n 비용 행렬 (cost[u][v] = u에서 v로 이동하는 비용, 0이면 이동 불가)
//
// [반환값]
//   - int: 최소 비용 해밀턴 경로의 비용 (-1이면 경로 없음)
func minHamiltonianPath(n int, cost [][]int) int {
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

	fmt.Fprintln(writer, minHamiltonianPath(n, cost))
}
