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
//
// [알고리즘 힌트]
//
//	위상 정렬 후 순서대로 간선을 완화한다.
//	Kahn's Algorithm으로 위상 정렬을 수행한다.
//	시간복잡도: O(V + E)
func dagShortestPath(n int, graph [][]Edge, s int) []int {
	inDegree := make([]int, n)
	for u := 0; u < n; u++ {
		for _, e := range graph[u] {
			inDegree[e.to]++
		}
	}

	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := []int{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, e := range graph[u] {
			inDegree[e.to]--
			if inDegree[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[s] = 0

	for _, u := range order {
		if dist[u] == INF {
			continue
		}
		for _, e := range graph[u] {
			if dist[u]+e.weight < dist[e.to] {
				dist[e.to] = dist[u] + e.weight
			}
		}
	}

	return dist
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
