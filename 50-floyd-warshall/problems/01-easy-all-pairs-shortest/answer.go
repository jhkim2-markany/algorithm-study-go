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
//
// [알고리즘 힌트]
//
//	3중 루프: 경유 정점 k, 출발 i, 도착 j 순서로 순회하며
//	dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])로 갱신한다.
//	시간 복잡도 O(N^3).
func floydWarshall(n int, edges [][3]int) [][]int {
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
				if dist[i][k] != INF && dist[k][j] != INF {
					if dist[i][k]+dist[k][j] < dist[i][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
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
