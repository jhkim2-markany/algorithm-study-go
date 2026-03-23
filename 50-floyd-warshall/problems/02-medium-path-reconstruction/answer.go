package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1

// floydWarshallWithPath는 플로이드-워셜 알고리즘으로 모든 쌍 최단 거리와
// 경로 복원용 next 행렬을 계산한다.
//
// [매개변수]
//   - n: 도시(정점) 수
//   - edges: 각 간선은 [u, v, w] 형태의 방향 간선
//
// [반환값]
//   - [][]int: dist[i][j] = i에서 j까지의 최단 거리
//   - [][]int: next[i][j] = i에서 j로 가는 최단 경로에서 i 다음에 방문할 정점
//
// [알고리즘 힌트]
//
//	플로이드-워셜 실행 시 dist 갱신과 함께 next[i][j] = next[i][k]로 경로를 추적한다.
//	경로 복원: s에서 e까지 next를 따라가며 경로를 구성한다.
func floydWarshallWithPath(n int, edges [][3]int) ([][]int, [][]int) {
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

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if w < dist[u][v] {
			dist[u][v] = w
			next[u][v] = v
		}
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					newDist := dist[i][k] + dist[k][j]
					if newDist < dist[i][j] {
						dist[i][j] = newDist
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}
	return dist, next
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

	dist, next := floydWarshallWithPath(n, edges)

	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		var s, e int
		fmt.Fscan(reader, &s, &e)

		if dist[s][e] == INF {
			fmt.Fprintln(writer, -1)
			fmt.Fprintln(writer, 0)
		} else {
			fmt.Fprintln(writer, dist[s][e])
			path := []int{s}
			cur := s
			for cur != e {
				cur = next[cur][e]
				path = append(path, cur)
			}
			fmt.Fprint(writer, len(path))
			for _, v := range path {
				fmt.Fprint(writer, " ", v)
			}
			fmt.Fprintln(writer)
		}
	}
}
