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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 프로젝트 수와 의존 관계 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 각 프로젝트의 이익 입력
	profit := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &profit[i])
	}

	// 그래프 초기화: 소스(0), 프로젝트(1~N), 싱크(N+1)
	source := 0
	sink := n + 1
	total := n + 2
	graph = make([][]Edge, total)
	for i := range graph {
		graph[i] = []Edge{}
	}

	// 프로젝트 선택 문제를 최소 컷으로 모델링한다
	// 이익이 양수인 프로젝트: 소스 → 프로젝트 (용량 = 이익)
	// 이익이 음수인 프로젝트: 프로젝트 → 싱크 (용량 = |이익|)
	// 의존 관계 (a→b): 프로젝트 a → 프로젝트 b (용량 = INF)
	//
	// 최소 컷에서:
	// - 소스 쪽에 남은 프로젝트 = 선택한 프로젝트
	// - 싱크 쪽으로 간 프로젝트 = 선택하지 않은 프로젝트
	// - 컷되는 간선 = 선택하지 않은 양수 이익 + 선택한 음수 이익(비용)
	// 최대 순이익 = 양수 이익 합 - 최소 컷

	sumPositive := 0
	for i := 1; i <= n; i++ {
		if profit[i] > 0 {
			// 이익이 양수: 소스에서 프로젝트로 간선 추가
			addEdge(source, i, profit[i])
			sumPositive += profit[i]
		} else if profit[i] < 0 {
			// 이익이 음수: 프로젝트에서 싱크로 간선 추가
			addEdge(i, sink, -profit[i])
		}
	}

	// 의존 관계 입력: a가 b에 의존 → a를 선택하면 b도 선택해야 함
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		// a를 선택하고 b를 선택하지 않으면 안 되므로 무한 용량 간선
		addEdge(a, b, INF)
	}

	// 최소 컷을 구한다
	minCut := edmondsKarp(source, sink, total)

	// 최대 순이익 = 양수 이익 합 - 최소 컷
	fmt.Fprintln(writer, sumPositive-minCut)
}
