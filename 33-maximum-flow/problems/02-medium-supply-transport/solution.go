package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 1<<60 - 1

// 간선 구조체
type Edge struct {
	to, cap, flow, rev int
}

// 원본 간선 정보 (최소 컷 출력용)
type OrigEdge struct {
	u, v, idx int
}

var graph [][]Edge

// addEdge: 용량 cap인 간선 u→v를 추가한다
func addEdge(u, v, cap int) {
	graph[u] = append(graph[u], Edge{to: v, cap: cap, flow: 0, rev: len(graph[v])})
	graph[v] = append(graph[v], Edge{to: u, cap: 0, flow: 0, rev: len(graph[u]) - 1})
}

// edmondsKarp: BFS로 최대 유량을 구한다
func edmondsKarp(s, t, n int) int {
	totalFlow := 0

	for {
		// BFS로 증가 경로를 찾는다
		parent := make([]int, n)
		parentEdge := make([]int, n)
		for i := range parent {
			parent[i] = -1
		}
		parent[s] = s

		queue := []int{s}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for idx, e := range graph[cur] {
				if e.cap-e.flow > 0 && parent[e.to] == -1 {
					parent[e.to] = cur
					parentEdge[e.to] = idx
					queue = append(queue, e.to)
				}
			}
		}

		if parent[t] == -1 {
			break
		}

		// 병목 용량을 구한다
		bottleneck := INF
		v := t
		for v != s {
			u := parent[v]
			e := graph[u][parentEdge[v]]
			if e.cap-e.flow < bottleneck {
				bottleneck = e.cap - e.flow
			}
			v = u
		}

		// 유량을 갱신한다
		v = t
		for v != s {
			u := parent[v]
			idx := parentEdge[v]
			graph[u][idx].flow += bottleneck
			graph[v][graph[u][idx].rev].flow -= bottleneck
			v = u
		}

		totalFlow += bottleneck
	}

	return totalFlow
}

// findMinCut: 최대 유량 계산 후 최소 컷에 포함되는 간선을 찾는다
func findMinCut(s, n int) []bool {
	// 잔여 그래프에서 소스로부터 도달 가능한 정점을 BFS로 찾는다
	reachable := make([]bool, n)
	queue := []int{s}
	reachable[s] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, e := range graph[cur] {
			// 잔여 용량이 있고 미방문이면 도달 가능하다
			if e.cap-e.flow > 0 && !reachable[e.to] {
				reachable[e.to] = true
				queue = append(queue, e.to)
			}
		}
	}

	return reachable
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 정점 수, 간선 수, 소스, 싱크
	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	// 그래프 초기화
	graph = make([][]Edge, n+1)
	for i := range graph {
		graph[i] = []Edge{}
	}

	// 원본 간선 정보를 저장한다
	origEdges := make([]OrigEdge, m)
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		origEdges[i] = OrigEdge{u: u, v: v, idx: i}
		addEdge(u, v, c)
	}

	// 최대 유량을 구한다
	maxFlow := edmondsKarp(s, t, n+1)
	fmt.Fprintln(writer, maxFlow)

	// 최소 컷을 구한다: 소스에서 도달 가능한 정점 집합을 찾는다
	reachable := findMinCut(s, n+1)

	// 최소 컷에 포함되는 간선을 찾는다
	// S 집합에서 T 집합으로 가는 원본 간선이 최소 컷이다
	var cutEdges []OrigEdge
	for _, oe := range origEdges {
		if reachable[oe.u] && !reachable[oe.v] {
			cutEdges = append(cutEdges, oe)
		}
	}

	// 입력 순서 기준으로 정렬한다
	sort.Slice(cutEdges, func(i, j int) bool {
		return cutEdges[i].idx < cutEdges[j].idx
	})

	// 최소 컷 간선 출력
	fmt.Fprintln(writer, len(cutEdges))
	for _, e := range cutEdges {
		fmt.Fprintln(writer, e.u, e.v)
	}
}
