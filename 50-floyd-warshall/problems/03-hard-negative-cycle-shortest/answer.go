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
//
// [알고리즘 힌트]
//
//	1단계: 표준 플로이드-워셜로 최단 거리를 계산한다.
//	2단계: dist[k][k] < 0인 정점 k를 찾아, i→k→j 경로가 존재하는 모든 (i,j) 쌍을 NEG_INF로 표시한다.
func floydWarshallNegCycle(n int, edges [][3]int) [][]int {
	dist := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if w < dist[u][v] {
			dist[u][v] = w
		}
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if dist[i][j] == NEG_INF {
					continue
				}
				if dist[i][k] != INF && dist[i][k] != NEG_INF &&
					dist[k][j] != INF && dist[k][j] != NEG_INF {
					newDist := dist[i][k] + dist[k][j]
					if newDist < dist[i][j] {
						dist[i][j] = newDist
					}
				}
			}
		}
	}

	// 음수 사이클 전파
	for k := 1; k <= n; k++ {
		if dist[k][k] < 0 {
			for i := 1; i <= n; i++ {
				for j := 1; j <= n; j++ {
					if dist[i][k] != INF && dist[k][j] != INF {
						dist[i][j] = NEG_INF
					}
				}
			}
		}
	}
	return dist
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
