package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// solveTSP는 비트마스크 DP로 외판원 문제의 최소 비용을 구한다.
//
// [매개변수]
//   - cost: 도시 간 비용 행렬 (n × n)
//   - n: 도시의 수
//
// [반환값]
//   - int: 모든 도시를 방문하고 출발 도시로 돌아오는 최소 비용 (-1이면 불가능)
func solveTSP(cost [][]int, n int) int {
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

	// 핵심 함수 호출
	result := solveTSP(cost, n)

	if result == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, result)
	}
}
