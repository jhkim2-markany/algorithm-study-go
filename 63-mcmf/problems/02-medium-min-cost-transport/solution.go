package main

import (
	"bufio"
	"fmt"
	"os"
)

// minCostTransport는 P개의 공장에서 W개의 창고로 물건을 운송할 때
// MCMF를 이용하여 최소 운송 비용을 구한다.
//
// [매개변수]
//   - p: 공장 수
//   - w: 창고 수
//   - supply: 각 공장의 공급량
//   - demand: 각 창고의 수요량
//   - cost: p×w 비용 행렬 (cost[i][j] = 공장 i에서 창고 j로의 단위 운송 비용)
//
// [반환값]
//   - int: 최소 운송 비용
func minCostTransport(p, w int, supply, demand []int, cost [][]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var p, w int
	fmt.Fscan(reader, &p, &w)

	supply := make([]int, p)
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &supply[i])
	}

	demand := make([]int, w)
	for i := 0; i < w; i++ {
		fmt.Fscan(reader, &demand[i])
	}

	cost := make([][]int, p)
	for i := 0; i < p; i++ {
		cost[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	fmt.Fprintln(writer, minCostTransport(p, w, supply, demand, cost))
}
