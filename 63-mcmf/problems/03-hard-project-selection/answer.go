package main

import (
	"bufio"
	"fmt"
	"os"
)

// maxProfit은 프로젝트와 직원의 의존 관계에서 최대 가중 폐합(Maximum Weight Closure)
// 모델을 이용하여 최대 순이익을 구한다.
//
// [매개변수]
//   - n: 프로젝트 수
//   - m: 직원 수
//   - profit: 각 프로젝트의 이익
//   - salary: 각 직원의 급여
//   - requires: 각 프로젝트가 요구하는 직원 번호 목록 (1-indexed)
//
// [반환값]
//   - int: 최대 순이익 (음수이면 0)
//
// [알고리즘 힌트]
//   1. S→프로젝트(용량=profit), 직원→T(용량=salary), 프로젝트→직원(용량=INF)로 네트워크를 구성한다
//   2. Edmonds-Karp(BFS 기반 최대 유량)로 최소 컷을 구한다
//   3. 최대 순이익 = 전체 프로젝트 이익 합 - 최소 컷(최대 유량)
//   4. 결과가 음수이면 0을 반환한다
func maxProfit(n, m int, profit, salary []int, requires [][]int) int {
	const INF = 1<<63 - 1

	type Edge struct {
		to, cap, flow int
	}

	nodeN := n + m + 2
	S, T := 0, n+m+1
	edges := make([]Edge, 0)
	graph := make([][]int, nodeN)

	addEdge := func(u, v, cap int) {
		graph[u] = append(graph[u], len(edges))
		edges = append(edges, Edge{v, cap, 0})
		graph[v] = append(graph[v], len(edges))
		edges = append(edges, Edge{u, 0, 0})
	}

	totalProfit := 0
	for i := 0; i < n; i++ {
		addEdge(S, i+1, profit[i])
		totalProfit += profit[i]
	}
	for j := 0; j < m; j++ {
		addEdge(n+1+j, T, salary[j])
	}
	for i := 0; i < n; i++ {
		for _, j := range requires[i] {
			addEdge(i+1, n+j, INF)
		}
	}

	// Edmonds-Karp
	bfs := func() bool {
		parent := make([]int, nodeN)
		for i := range parent {
			parent[i] = -1
		}
		parent[S] = S
		queue := []int{S}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, idx := range graph[u] {
				e := &edges[idx]
				if parent[e.to] == -1 && e.cap-e.flow > 0 {
					parent[e.to] = idx
					if e.to == T {
						// 경로 상 최소 잔여 용량
						f := INF
						for v := T; v != S; {
							idx2 := parent[v]
							if edges[idx2].cap-edges[idx2].flow < f {
								f = edges[idx2].cap - edges[idx2].flow
							}
							v = edges[idx2^1].to
						}
						for v := T; v != S; {
							idx2 := parent[v]
							edges[idx2].flow += f
							edges[idx2^1].flow -= f
							v = edges[idx2^1].to
						}
						return true
					}
					queue = append(queue, e.to)
				}
			}
		}
		return false
	}

	total := 0
	for bfs() {
		// 유량은 bfs 내부에서 전송됨, 여기서 누적
	}
	// 총 유량 계산
	for _, idx := range graph[S] {
		total += edges[idx].flow
	}

	result := totalProfit - total
	if result < 0 {
		result = 0
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	profit := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &profit[i])
	}

	salary := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &salary[i])
	}

	requires := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(reader, &k)
		requires[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &requires[i][j])
		}
	}

	fmt.Fprintln(writer, maxProfit(n, m, profit, salary, requires))
}
