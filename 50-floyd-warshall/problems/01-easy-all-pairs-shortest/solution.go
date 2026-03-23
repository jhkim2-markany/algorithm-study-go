package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1

// floydWarshall은 플로이드-워셜 알고리즘으로 모든 쌍 최단 거리를 계산한다.
//
// [매개변수]
//   - n: 도시(정점) 수
//   - edges: 각 간선은 [u, v, w] 형태의 방향 간선
//
// [반환값]
//   - [][]int: dist[i][j] = i에서 j까지의 최단 거리 (도달 불가 시 INF)
func floydWarshall(n int, edges [][3]int) [][]int {
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

	dist := floydWarshall(n, edges)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j > 1 {
				fmt.Fprint(writer, " ")
			}
			if dist[i][j] == INF {
				fmt.Fprint(writer, 0)
			} else {
				fmt.Fprint(writer, dist[i][j])
			}
		}
		fmt.Fprintln(writer)
	}
}
