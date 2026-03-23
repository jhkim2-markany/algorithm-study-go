package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MOD     = 1000000007
	NEG_INF = -int(1e18)
)

// Edge는 가중치 간선을 나타낸다.
type Edge struct {
	to, weight int
}

// DAGResult는 DAG 경로 분석 결과를 담는다.
type DAGResult struct {
	totalPaths int // S에서 T까지의 전체 경로 수 (mod)
	maxDist    int // S에서 T까지의 최장 거리
	maxPaths   int // 최장 경로를 달성하는 경로 수 (mod)
}

// dagPathAnalysis는 DAG에서 경로 수, 최장 거리, 최장 경로 수를 구한다.
//
// [매개변수]
//   - n: 정점의 수 (0-indexed)
//   - graph: 인접 리스트 (graph[u] = u에서 나가는 간선 목록)
//   - s: 시작 정점 번호
//   - t: 도착 정점 번호
//
// [반환값]
//   - DAGResult: 전체 경로 수, 최장 거리, 최장 경로 수
func dagPathAnalysis(n int, graph [][]Edge, s, t int) DAGResult {
	// 여기에 코드를 작성하세요
	return DAGResult{}
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

	result := dagPathAnalysis(n, graph, s, t)
	if result.totalPaths == 0 {
		fmt.Fprintln(writer, 0)
		fmt.Fprintln(writer, 0)
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, result.totalPaths)
		fmt.Fprintln(writer, result.maxDist)
		fmt.Fprintln(writer, result.maxPaths)
	}
}
