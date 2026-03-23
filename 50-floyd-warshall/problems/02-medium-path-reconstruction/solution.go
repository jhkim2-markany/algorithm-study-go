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
func floydWarshallWithPath(n int, edges [][3]int) ([][]int, [][]int) {
	// 여기에 코드를 작성하세요
	return nil, nil
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
