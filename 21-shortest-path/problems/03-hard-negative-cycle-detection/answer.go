package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 간선을 나타낸다 (출발, 도착, 가중치)
type Edge struct {
	from, to, weight int
}

// hasNegativeCycle은 그래프에서 음수 사이클의 존재 여부를 판별한다.
//
// [매개변수]
//   - n: 정점의 수
//   - edges: 간선 목록 (출발, 도착, 가중치)
//
// [반환값]
//   - bool: 음수 사이클이 존재하면 true, 아니면 false
//
// [알고리즘 힌트]
//
//	Bellman-Ford 알고리즘을 사용한다.
//	모든 정점의 거리를 0으로 초기화하여 비연결 그래프도 처리한다.
//	(N-1)번 반복하여 최단 거리를 확정한 후,
//	N번째 반복에서 갱신이 발생하면 음수 사이클이 존재한다.
func hasNegativeCycle(n int, edges []Edge) bool {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = 0
	}

	// (N-1)번 반복하여 최단 거리를 확정한다
	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if dist[e.from]+e.weight < dist[e.to] {
				dist[e.to] = dist[e.from] + e.weight
			}
		}
	}

	// N번째 반복에서 갱신이 발생하면 음수 사이클 존재
	for _, e := range edges {
		if dist[e.from]+e.weight < dist[e.to] {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		edges := make([]Edge, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &edges[i].from, &edges[i].to, &edges[i].weight)
		}

		// 핵심 함수 호출
		if hasNegativeCycle(n, edges) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
