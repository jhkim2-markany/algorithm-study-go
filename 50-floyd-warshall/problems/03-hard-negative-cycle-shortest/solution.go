package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1
const NEG_INF = -(1<<31 - 1)

// floydWarshallNegCycle은 플로이드-워셜 알고리즘으로 모든 쌍 최단 거리를 계산하고,
// 음수 사이클의 영향을 받는 쌍을 NEG_INF로 표시한다.
//
// [매개변수]
//   - n: 도시(정점) 수
//   - edges: 각 간선은 [u, v, w] 형태의 방향 간선
//
// [반환값]
//   - [][]int: dist[i][j] = 최단 거리 (도달 불가 시 INF, 음수 사이클 영향 시 NEG_INF)
func floydWarshallNegCycle(n int, edges [][3]int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	dist := floydWarshallNegCycle(n, edges)

	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		var s, e int
		fmt.Fscan(reader, &s, &e)

		if dist[s][e] == INF {
			fmt.Fprintln(writer, -1)
		} else if dist[s][e] == NEG_INF {
			fmt.Fprintln(writer, -2)
		} else {
			fmt.Fprintln(writer, dist[s][e])
		}
	}
}
