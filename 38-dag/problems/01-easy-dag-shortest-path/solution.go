package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e18)

// Edge는 가중치 간선을 나타낸다.
type Edge struct {
	to, weight int
}

// dagShortestPath는 DAG에서 시작 정점으로부터 각 정점까지의 최단 거리를 반환한다.
//
// [매개변수]
//   - n: 정점의 수 (0-indexed)
//   - graph: 인접 리스트 (graph[u] = u에서 나가는 간선 목록)
//   - s: 시작 정점 번호
//
// [반환값]
//   - []int: 각 정점까지의 최단 거리 (도달 불가능하면 INF)
func dagShortestPath(n int, graph [][]Edge, s int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s int
	fmt.Fscan(reader, &n, &m, &s)

	graph := make([][]Edge, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		graph[u] = append(graph[u], Edge{v, w})
	}

	dist := dagShortestPath(n, graph, s)
	for i := 0; i < n; i++ {
		if dist[i] == INF {
			fmt.Fprintln(writer, "INF")
		} else {
			fmt.Fprintln(writer, dist[i])
		}
	}
}
