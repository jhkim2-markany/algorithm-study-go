package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 가중치 간선을 나타낸다.
type Edge struct {
	to, weight int
}

// dagLongestPath는 DAG에서 시작 정점에서 도착 정점까지의 최장 경로 비용을 반환한다.
//
// [매개변수]
//   - n: 정점의 수 (0-indexed)
//   - graph: 인접 리스트 (graph[u] = u에서 나가는 간선 목록)
//   - s: 시작 정점 번호
//   - t: 도착 정점 번호
//
// [반환값]
//   - int: 최장 경로의 비용 (도달 불가능하면 -1)
func dagLongestPath(n int, graph [][]Edge, s, t int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	graph := make([][]Edge, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		graph[u] = append(graph[u], Edge{v, w})
	}

	result := dagLongestPath(n, graph, s, t)
	if result == -1 {
		fmt.Fprintln(writer, "IMPOSSIBLE")
	} else {
		fmt.Fprintln(writer, result)
	}
}
