package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 도시 수 N, 도로 수 M
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 거리 행렬 초기화
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

	// 간선 입력 (같은 쌍에 여러 도로가 있으면 최솟값만 저장)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		if w < dist[u][v] {
			dist[u][v] = w
		}
	}

	// 플로이드-워셜: 3중 루프로 모든 쌍 최단 거리 계산
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

	// 출력: 도달 불가능하면 0 출력
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
