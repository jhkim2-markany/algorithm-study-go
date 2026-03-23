package main

import (
	"bufio"
	"fmt"
	"os"
)

const NEG_INF = -int(1e18)

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
//
// [알고리즘 힌트]
//
//	위상 정렬 후 순서대로 간선을 완화한다 (최장 경로).
//	dist를 NEG_INF로 초기화하고, 더 큰 값으로 갱신한다.
//	시간복잡도: O(V + E)
func dagLongestPath(n int, graph [][]Edge, s, t int) int {
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
		dist[i] = NEG_INF
	}
	dist[s] = 0

	for _, u := range order {
		if dist[u] == NEG_INF {
			continue
		}
		for _, e := range graph[u] {
			if dist[u]+e.weight > dist[e.to] {
				dist[e.to] = dist[u] + e.weight
			}
		}
	}

	if dist[t] == NEG_INF {
		return -1
	}
	return dist[t]
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
