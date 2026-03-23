package main

import (
	"bufio"
	"fmt"
	"os"
)

// 최소 비용 할당 문제 - MCMF 풀이
// N명의 작업자를 N개의 작업에 1:1 배정하여 최소 비용을 구한다.

const INF = 1<<63 - 1

// Edge는 유량 네트워크의 간선을 나타낸다
type Edge struct {
	to, cap, cost, flow int
}

var (
	edges []Edge
	graph [][]int
	nodeN int // 전체 노드 수
)

// addEdge는 u→v 간선과 역방향 간선을 쌍으로 추가한다
func addEdge(u, v, cap, cost int) {
	graph[u] = append(graph[u], len(edges))
	edges = append(edges, Edge{v, cap, cost, 0})
	graph[v] = append(graph[v], len(edges))
	edges = append(edges, Edge{u, 0, -cost, 0})
}

// mcmf는 소스 s에서 싱크 t로 최소 비용 최대 유량을 구한다
func mcmf(s, t int) (int, int) {
	totalFlow := 0
	totalCost := 0

	for {
		// SPFA로 최소 비용 경로 탐색
		dist := make([]int, nodeN)
		for i := range dist {
			dist[i] = INF
		}
		dist[s] = 0

		inQueue := make([]bool, nodeN)
		inQueue[s] = true

		prevEdge := make([]int, nodeN)
		for i := range prevEdge {
			prevEdge[i] = -1
		}

		queue := []int{s}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			inQueue[u] = false

			for _, idx := range graph[u] {
				e := &edges[idx]
				if e.cap-e.flow > 0 && dist[u]+e.cost < dist[e.to] {
					dist[e.to] = dist[u] + e.cost
					prevEdge[e.to] = idx
					if !inQueue[e.to] {
						inQueue[e.to] = true
						queue = append(queue, e.to)
					}
				}
			}
		}

		// 싱크에 도달 불가 → 종료
		if dist[t] == INF {
			break
		}

		// 경로 상 최소 잔여 용량
		f := INF
		for v := t; v != s; {
			idx := prevEdge[v]
			if edges[idx].cap-edges[idx].flow < f {
				f = edges[idx].cap - edges[idx].flow
			}
			v = edges[idx^1].to
		}

		// 유량 전송
		for v := t; v != s; {
			idx := prevEdge[v]
			edges[idx].flow += f
			edges[idx^1].flow -= f
			v = edges[idx^1].to
		}

		totalFlow += f
		totalCost += f * dist[t]
	}

	return totalFlow, totalCost
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 작업자/작업 수
	var n int
	fmt.Fscan(reader, &n)

	// 입력: 비용 행렬
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	// 네트워크 구성
	// 노드: S(0), 작업자(1~N), 작업(N+1~2N), T(2N+1)
	S := 0
	T := 2*n + 1
	nodeN = 2*n + 2
	edges = make([]Edge, 0)
	graph = make([][]int, nodeN)

	// S → 작업자 (용량 1, 비용 0)
	for i := 1; i <= n; i++ {
		addEdge(S, i, 1, 0)
	}

	// 작업자 → 작업 (용량 1, 비용 = cost[i][j])
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			addEdge(i+1, n+1+j, 1, cost[i][j])
		}
	}

	// 작업 → T (용량 1, 비용 0)
	for j := 0; j < n; j++ {
		addEdge(n+1+j, T, 1, 0)
	}

	// MCMF 수행
	_, minCost := mcmf(S, T)

	// 출력: 최소 비용
	fmt.Fprintln(writer, minCost)
}
