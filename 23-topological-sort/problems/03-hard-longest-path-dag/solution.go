package main

import (
	"bufio"
	"fmt"
	"os"
)

const NEG_INF = -1 << 60

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s int
	fmt.Fscan(reader, &n, &m, &s)

	// 인접 리스트와 진입 차수 배열 초기화
	type Edge struct {
		to, weight int
	}
	adj := make([][]Edge, n+1)
	inDegree := make([]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []Edge{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], Edge{v, w})
		inDegree[v]++
	}

	// Kahn 알고리즘으로 위상 정렬 수행
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := make([]int, 0, n)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		order = append(order, cur)

		for _, e := range adj[cur] {
			inDegree[e.to]--
			if inDegree[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	// 최장 경로 DP: dist[v] = 시작 정점에서 v까지의 최장 거리
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = NEG_INF
	}
	dist[s] = 0

	// 위상 정렬 순서대로 DP 갱신
	for _, u := range order {
		if dist[u] == NEG_INF {
			continue
		}
		// 현재 정점에서 나가는 간선을 통해 최장 거리 갱신
		for _, e := range adj[u] {
			if dist[u]+e.weight > dist[e.to] {
				dist[e.to] = dist[u] + e.weight
			}
		}
	}

	// 결과 출력
	for i := 1; i <= n; i++ {
		if dist[i] == NEG_INF {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, dist[i])
		}
	}
}
