package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1
const NEG_INF = -(1<<31 - 1) // 음수 사이클 영향을 나타내는 값

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

	// 간선 입력 (같은 쌍에 여러 간선이 있으면 최솟값 저장)
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
				// 이미 음수 무한인 경우 건너뛴다
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

	// 음수 사이클 전파: dist[k][k] < 0인 정점 k를 경유하는 모든 쌍을 NEG_INF로 표시
	for k := 1; k <= n; k++ {
		if dist[k][k] < 0 {
			// k는 음수 사이클에 포함된다
			for i := 1; i <= n; i++ {
				for j := 1; j <= n; j++ {
					// i에서 k로 도달 가능하고 k에서 j로 도달 가능하면
					if dist[i][k] != INF && dist[k][j] != INF {
						dist[i][j] = NEG_INF
					}
				}
			}
		}
	}

	// 질의 처리
	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		var s, e int
		fmt.Fscan(reader, &s, &e)

		if dist[s][e] == INF {
			// 도달 불가능
			fmt.Fprintln(writer, -1)
		} else if dist[s][e] == NEG_INF {
			// 음수 사이클로 무한히 줄일 수 있음
			fmt.Fprintln(writer, -2)
		} else {
			// 일반 최단 거리
			fmt.Fprintln(writer, dist[s][e])
		}
	}
}
