package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// 간선 구조체
type Edge struct {
	to, cap, flow, rev int
}

var graph [][]Edge

// addEdge: 용량 cap인 간선 u→v를 추가한다
func addEdge(u, v, cap int) {
	graph[u] = append(graph[u], Edge{to: v, cap: cap, flow: 0, rev: len(graph[v])})
	graph[v] = append(graph[v], Edge{to: u, cap: 0, flow: 0, rev: len(graph[u]) - 1})
}

// edmondsKarp: BFS로 증가 경로를 찾아 최대 유량을 구한다
func edmondsKarp(s, t, n int) int {
	totalFlow := 0

	for {
		// BFS로 증가 경로를 탐색한다
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
				// 잔여 용량이 있고 미방문 정점이면 탐색한다
				if e.cap-e.flow > 0 && parent[e.to] == -1 {
					parent[e.to] = cur
					parentEdge[e.to] = idx
					queue = append(queue, e.to)
				}
			}
		}

		// 싱크에 도달 불가능하면 종료한다
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

		// 경로를 따라 유량을 갱신한다
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정점 수와 간선 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 그래프 초기화
	graph = make([][]Edge, n+1)
	for i := range graph {
		graph[i] = []Edge{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		addEdge(u, v, c)
	}

	// 소스: 1, 싱크: N
	result := edmondsKarp(1, n, n+1)
	fmt.Fprintln(writer, result)
}
