package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Edge는 간선을 나타낸다 (출발, 도착, 가중치)
type Edge struct {
	from, to, weight int
}

// hasNegativeCycle 함수는 Bellman-Ford 알고리즘으로 음수 사이클을 탐지한다
// 그래프가 연결되지 않을 수 있으므로 모든 정점을 출발점으로 초기화한다
func hasNegativeCycle(n int, edges []Edge) bool {
	dist := make([]int, n+1)
	// 모든 정점의 거리를 0으로 초기화 (가상의 시작점에서 모든 정점으로 가중치 0 간선이 있다고 가정)
	for i := range dist {
		dist[i] = 0
	}

	// (N-1)번 반복하여 최단 거리를 확정한다
	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			// 완화(relaxation): 더 짧은 경로가 발견되면 갱신
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

// hasNegativeCycleSPFA 함수는 SPFA 알고리즘으로 음수 사이클을 탐지한다
// 어떤 정점이 큐에 N번 이상 들어가면 음수 사이클이 존재한다
func hasNegativeCycleSPFA(n int, edges []Edge) bool {
	// 인접 리스트 구성
	graph := make([][]struct{ to, w int }, n+1)
	for _, e := range edges {
		graph[e.from] = append(graph[e.from], struct{ to, w int }{e.to, e.weight})
	}

	dist := make([]int, n+1)
	inQueue := make([]bool, n+1)
	count := make([]int, n+1) // 큐에 들어간 횟수

	for i := range dist {
		dist[i] = math.MaxInt64
	}

	// 모든 정점을 큐에 넣어 비연결 그래프도 처리
	queue := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		dist[i] = 0
		inQueue[i] = true
		queue = append(queue, i)
		count[i] = 1
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		inQueue[u] = false

		for _, e := range graph[u] {
			if dist[u]+e.w < dist[e.to] {
				dist[e.to] = dist[u] + e.w
				if !inQueue[e.to] {
					inQueue[e.to] = true
					queue = append(queue, e.to)
					count[e.to]++
					// 큐에 N번 이상 들어가면 음수 사이클
					if count[e.to] >= n {
						return true
					}
				}
			}
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

		// Bellman-Ford로 음수 사이클 탐지
		if hasNegativeCycle(n, edges) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
