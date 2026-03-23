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

	// 거리 행렬과 경로 복원용 next 행렬 초기화
	dist := make([][]int, n+1)
	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = make([]int, n+1)
		next[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
				next[i][j] = j
			} else {
				dist[i][j] = INF
				next[i][j] = -1
			}
		}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		if w < dist[u][v] {
			dist[u][v] = w
			next[u][v] = v
		}
	}

	// 플로이드-워셜: 최단 거리와 경로 동시 갱신
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					newDist := dist[i][k] + dist[k][j]
					if newDist < dist[i][j] {
						dist[i][j] = newDist
						// 경로 복원: i에서 j로 갈 때 k를 경유하므로 next 갱신
						next[i][j] = next[i][k]
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
			fmt.Fprintln(writer, 0)
		} else {
			// 최소 비용 출력
			fmt.Fprintln(writer, dist[s][e])

			// 경로 복원: next를 따라가며 경로를 구성한다
			path := []int{s}
			cur := s
			for cur != e {
				cur = next[cur][e]
				path = append(path, cur)
			}

			// 경로 출력: 도시 수와 경로
			fmt.Fprint(writer, len(path))
			for _, v := range path {
				fmt.Fprint(writer, " ", v)
			}
			fmt.Fprintln(writer)
		}
	}
}
