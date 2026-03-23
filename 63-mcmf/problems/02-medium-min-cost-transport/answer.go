package main

import (
	"bufio"
	"fmt"
	"os"
)

// minCostTransport는 P개의 공장에서 W개의 창고로 물건을 운송할 때
// MCMF를 이용하여 최소 운송 비용을 구한다.
//
// [매개변수]
//   - p: 공장 수
//   - w: 창고 수
//   - supply: 각 공장의 공급량
//   - demand: 각 창고의 수요량
//   - cost: p×w 비용 행렬 (cost[i][j] = 공장 i에서 창고 j로의 단위 운송 비용)
//
// [반환값]
//   - int: 최소 운송 비용
//
// [알고리즘 힌트]
//   1. 네트워크 구성: S→공장(용량=supply), 공장→창고(용량=min(supply,demand),비용=cost), 창고→T(용량=demand)
//   2. SPFA로 최소 비용 경로를 반복 탐색한다
//   3. 경로 상 최소 잔여 용량만큼 유량을 전송하고 비용을 누적한다
//   4. 싱크에 도달 불가능하면 종료하고 총 비용을 반환한다
func minCostTransport(p, w int, supply, demand []int, cost [][]int) int {
	const INF = 1<<63 - 1

	type Edge struct {
		to, cap, cost, flow int
	}

	nodeN := p + w + 2
	S, T := 0, p+w+1
	edges := make([]Edge, 0)
	graph := make([][]int, nodeN)

	addEdge := func(u, v, cap, cost int) {
		graph[u] = append(graph[u], len(edges))
		edges = append(edges, Edge{v, cap, cost, 0})
		graph[v] = append(graph[v], len(edges))
		edges = append(edges, Edge{u, 0, -cost, 0})
	}

	for i := 0; i < p; i++ {
		addEdge(S, i+1, supply[i], 0)
	}
	for i := 0; i < p; i++ {
		for j := 0; j < w; j++ {
			cap := supply[i]
			if demand[j] < cap {
				cap = demand[j]
			}
			addEdge(i+1, p+1+j, cap, cost[i][j])
		}
	}
	for j := 0; j < w; j++ {
		addEdge(p+1+j, T, demand[j], 0)
	}

	totalCost := 0
	for {
		dist := make([]int, nodeN)
		for i := range dist {
			dist[i] = INF
		}
		dist[S] = 0
		inQueue := make([]bool, nodeN)
		inQueue[S] = true
		prevEdge := make([]int, nodeN)
		for i := range prevEdge {
			prevEdge[i] = -1
		}
		queue := []int{S}

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

		if dist[T] == INF {
			break
		}

		f := INF
		for v := T; v != S; {
			idx := prevEdge[v]
			if edges[idx].cap-edges[idx].flow < f {
				f = edges[idx].cap - edges[idx].flow
			}
			v = edges[idx^1].to
		}
		for v := T; v != S; {
			idx := prevEdge[v]
			edges[idx].flow += f
			edges[idx^1].flow -= f
			v = edges[idx^1].to
		}
		totalCost += f * dist[T]
	}

	return totalCost
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var p, w int
	fmt.Fscan(reader, &p, &w)

	supply := make([]int, p)
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &supply[i])
	}

	demand := make([]int, w)
	for i := 0; i < w; i++ {
		fmt.Fscan(reader, &demand[i])
	}

	cost := make([][]int, p)
	for i := 0; i < p; i++ {
		cost[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	fmt.Fprintln(writer, minCostTransport(p, w, supply, demand, cost))
}
