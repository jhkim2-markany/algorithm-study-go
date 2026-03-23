package main

import (
	"bufio"
	"fmt"
	"os"
)

// minCostAssignment은 N명의 작업자를 N개의 작업에 1:1 배정할 때
// MCMF(최소 비용 최대 유량)를 이용하여 최소 비용을 구한다.
//
// [매개변수]
//   - n: 작업자/작업 수
//   - cost: n×n 비용 행렬 (cost[i][j] = 작업자 i가 작업 j를 수행하는 비용)
//
// [반환값]
//   - int: 최소 배정 비용
//
// [알고리즘 힌트]
//   1. 네트워크를 구성한다: S→작업자(용량1,비용0), 작업자→작업(용량1,비용cost[i][j]), 작업→T(용량1,비용0)
//   2. SPFA로 최소 비용 경로를 반복 탐색한다
//   3. 경로 상 최소 잔여 용량만큼 유량을 전송하고 비용을 누적한다
//   4. 싱크에 도달 불가능하면 종료하고 총 비용을 반환한다
func minCostAssignment(n int, cost [][]int) int {
	const INF = 1<<63 - 1

	type Edge struct {
		to, cap, cost, flow int
	}

	nodeN := 2*n + 2
	S, T := 0, 2*n+1
	edges := make([]Edge, 0)
	graph := make([][]int, nodeN)

	addEdge := func(u, v, cap, cost int) {
		graph[u] = append(graph[u], len(edges))
		edges = append(edges, Edge{v, cap, cost, 0})
		graph[v] = append(graph[v], len(edges))
		edges = append(edges, Edge{u, 0, -cost, 0})
	}

	for i := 1; i <= n; i++ {
		addEdge(S, i, 1, 0)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			addEdge(i+1, n+1+j, 1, cost[i][j])
		}
	}
	for j := 0; j < n; j++ {
		addEdge(n+1+j, T, 1, 0)
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

	var n int
	fmt.Fscan(reader, &n)

	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &cost[i][j])
		}
	}

	fmt.Fprintln(writer, minCostAssignment(n, cost))
}
