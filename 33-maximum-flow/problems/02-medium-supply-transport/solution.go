package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 유량 네트워크의 간선을 나타낸다.
type Edge struct {
	to, cap, flow, rev int
}

// FlowResult는 최대 유량과 최소 컷 간선 정보를 담는다.
type FlowResult struct {
	maxFlow  int
	cutEdges [][2]int // 최소 컷에 포함되는 간선 [u, v] 목록
}

// solveMaxFlowMinCut은 최대 유량과 최소 컷 간선을 구한다.
//
// [매개변수]
//   - n: 정점의 수
//   - edges: 간선 목록 (각 원소는 [u, v, cap])
//   - source: 소스 정점 번호
//   - sink: 싱크 정점 번호
//
// [반환값]
//   - FlowResult: 최대 유량과 최소 컷 간선 목록
func solveMaxFlowMinCut(n int, edges [][3]int, source, sink int) FlowResult {
	// 여기에 코드를 작성하세요
	return FlowResult{}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	edges := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	result := solveMaxFlowMinCut(n, edges, s, t)
	fmt.Fprintln(writer, result.maxFlow)
	fmt.Fprintln(writer, len(result.cutEdges))
	for _, e := range result.cutEdges {
		fmt.Fprintln(writer, e[0], e[1])
	}
}
