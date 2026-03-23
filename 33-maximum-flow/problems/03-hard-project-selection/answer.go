package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<60 - 1

// Edge는 유량 네트워크의 간선을 나타낸다.
type Edge struct {
	to, cap, flow, rev int
}

var graph [][]Edge

func addEdge(u, v, cap int) {
	graph[u] = append(graph[u], Edge{to: v, cap: cap, flow: 0, rev: len(graph[v])})
	graph[v] = append(graph[v], Edge{to: u, cap: 0, flow: 0, rev: len(graph[u]) - 1})
}

func edmondsKarp(s, t, n int) int {
	totalFlow := 0
	for {
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

// maxProfit은 의존 관계를 만족하면서 프로젝트를 선택했을 때의 최대 순이익을 반환한다.
//
// [매개변수]
//   - n: 프로젝트의 수
//   - profit: 각 프로젝트의 이익 배열 (1-indexed, 음수이면 비용)
//   - deps: 의존 관계 목록 (각 원소는 [a, b], a가 b에 의존)
//
// [반환값]
//   - int: 최대 순이익
//
// [알고리즘 힌트]
//
//	프로젝트 선택 문제를 최소 컷으로 모델링한다.
//	양수 이익 프로젝트: 소스→프로젝트 (용량=이익).
//	음수 이익 프로젝트: 프로젝트→싱크 (용량=|이익|).
//	의존 관계: 프로젝트→프로젝트 (용량=INF).
//	최대 순이익 = 양수 이익 합 - 최소 컷.
//	시간복잡도: O(V * E^2)
func maxProfit(n int, profit []int, deps [][2]int) int {
	source := 0
	sink := n + 1
	total := n + 2
	graph = make([][]Edge, total)
	for i := range graph {
		graph[i] = []Edge{}
	}

	sumPositive := 0
	for i := 1; i <= n; i++ {
		if profit[i] > 0 {
			addEdge(source, i, profit[i])
			sumPositive += profit[i]
		} else if profit[i] < 0 {
			addEdge(i, sink, -profit[i])
		}
	}

	for _, d := range deps {
		addEdge(d[0], d[1], INF)
	}

	minCut := edmondsKarp(source, sink, total)
	return sumPositive - minCut
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	profit := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &profit[i])
	}

	deps := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &deps[i][0], &deps[i][1])
	}

	fmt.Fprintln(writer, maxProfit(n, profit, deps))
}
