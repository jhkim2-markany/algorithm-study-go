package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 1<<60 - 1

// Edge는 유량 네트워크의 간선을 나타낸다.
type Edge struct {
	to, cap, flow, rev int
}

// OrigEdge는 원본 간선 정보를 저장한다.
type OrigEdge struct {
	u, v, idx int
}

// FlowResult는 최대 유량과 최소 컷 간선 정보를 담는다.
type FlowResult struct {
	maxFlow  int
	cutEdges [][2]int
}

var graph [][]Edge

func addEdge(u, v, cap int) {
	graph[u] = append(graph[u], Edge{to: v, cap: cap, flow: 0, rev: len(graph[v])})
	graph[v] = append(graph[v], Edge{to: u, cap: 0, flow: 0, rev: len(graph[u]) - 1})
}

// solveMaxFlowMinCut은 최대 유량과 최소 컷 간선을 구한다.
//
// [매개변수]
//   - n: 정점의 수
//   - edges: 간선 목록 (각 원소는 [u, v, cap])
//   - source: 소스 정점 번호
//   - sink: 싱크 정점 번호
//
// [반환값]
//   - FlowResult: 최대 유량과 최소 컷 간선 목록
//
// [알고리즘 힌트]
//
//	에드몬드-카프로 최대 유량을 구한 뒤,
//	잔여 그래프에서 소스로부터 BFS로 도달 가능한 정점을 찾아 최소 컷을 구한다.
//	S 집합에서 T 집합으로 가는 원본 간선이 최소 컷이다.
//	시간복잡도: O(V * E^2)
func solveMaxFlowMinCut(n int, edges [][3]int, source, sink int) FlowResult {
	total := n + 1
	graph = make([][]Edge, total)
	for i := range graph {
		graph[i] = []Edge{}
	}

	origEdges := make([]OrigEdge, len(edges))
	for i, e := range edges {
		origEdges[i] = OrigEdge{u: e[0], v: e[1], idx: i}
		addEdge(e[0], e[1], e[2])
	}

	// 에드몬드-카프
	totalFlow := 0
	for {
		parent := make([]int, total)
		parentEdge := make([]int, total)
		for i := range parent {
			parent[i] = -1
		}
		parent[source] = source

		queue := []int{source}
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

		if parent[sink] == -1 {
			break
		}

		bottleneck := INF
		v := sink
		for v != source {
			u := parent[v]
			e := graph[u][parentEdge[v]]
			if e.cap-e.flow < bottleneck {
				bottleneck = e.cap - e.flow
			}
			v = u
		}

		v = sink
		for v != source {
			u := parent[v]
			idx := parentEdge[v]
			graph[u][idx].flow += bottleneck
			graph[v][graph[u][idx].rev].flow -= bottleneck
			v = u
		}

		totalFlow += bottleneck
	}

	// 최소 컷: 잔여 그래프에서 소스로부터 도달 가능한 정점
	reachable := make([]bool, total)
	queue := []int{source}
	reachable[source] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, e := range graph[cur] {
			if e.cap-e.flow > 0 && !reachable[e.to] {
				reachable[e.to] = true
				queue = append(queue, e.to)
			}
		}
	}

	var cutOrig []OrigEdge
	for _, oe := range origEdges {
		if reachable[oe.u] && !reachable[oe.v] {
			cutOrig = append(cutOrig, oe)
		}
	}

	sort.Slice(cutOrig, func(i, j int) bool {
		return cutOrig[i].idx < cutOrig[j].idx
	})

	cutEdges := make([][2]int, len(cutOrig))
	for i, oe := range cutOrig {
		cutEdges[i] = [2]int{oe.u, oe.v}
	}

	return FlowResult{maxFlow: totalFlow, cutEdges: cutEdges}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	edges := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	result := solveMaxFlowMinCut(n, edges, s, t)
	fmt.Fprintln(writer, result.maxFlow)
	fmt.Fprintln(writer, len(result.cutEdges))
	for _, e := range result.cutEdges {
		fmt.Fprintln(writer, e[0], e[1])
	}
}
