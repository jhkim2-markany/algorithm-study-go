package main

import (
	"bufio"
	"fmt"
	"os"
)

// 최소 비용 수송 문제 - MCMF 풀이
// P개의 공장에서 W개의 창고로 물건을 최소 비용으로 운송한다.

const INF = 1<<63 - 1

// Edge는 유량 네트워크의 간선을 나타낸다
type Edge struct {
	to, cap, cost, flow int
}

var (
	edges []Edge
	graph [][]int
	nodeN int
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

	// 입력: 공장 수 P, 창고 수 W
	var p, w int
	fmt.Fscan(reader, &p, &w)

	// 입력: 각 공장의 공급량
	supply := make([]int, p)
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &supply[i])
	}

	// 입력: 각 창고의 수요량
	demand := make([]int, w)
	for i := 0; i < w; i++ {
		fmt.Fscan(reader, &demand[i])
	}

	// 입력: 비용 행렬
	cost := make([][]int, p)
	for i := 0; i < p; i++ {
		cost[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	// 네트워크 구성
	// 노드: S(0), 공장(1~P), 창고(P+1~P+W), T(P+W+1)
	S := 0
	T := p + w + 1
	nodeN = p + w + 2
	edges = make([]Edge, 0)
	graph = make([][]int, nodeN)

	// S → 공장 i (용량 = supply[i], 비용 0)
	for i := 0; i < p; i++ {
		addEdge(S, i+1, supply[i], 0)
	}

	// 공장 i → 창고 j (용량 = min(supply[i], demand[j]), 비용 = cost[i][j])
	// 용량을 충분히 크게 잡아도 S→공장, 창고→T 간선이 제한한다
	for i := 0; i < p; i++ {
		for j := 0; j < w; j++ {
			cap := supply[i]
			if demand[j] < cap {
				cap = demand[j]
			}
			addEdge(i+1, p+1+j, cap, cost[i][j])
		}
	}

	// 창고 j → T (용량 = demand[j], 비용 0)
	for j := 0; j < w; j++ {
		addEdge(p+1+j, T, demand[j], 0)
	}

	// MCMF 수행
	_, minCost := mcmf(S, T)

	// 출력: 최소 운송 비용
	fmt.Fprintln(writer, minCost)
}
