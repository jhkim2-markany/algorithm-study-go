package main

import (
	"bufio"
	"fmt"
	"os"
)

// 프로젝트 선택 문제 - 최소 컷(최대 유량) 풀이
// 프로젝트와 직원의 의존 관계에서 최대 순이익을 구한다.
//
// 모델링 (최대 가중 폐합, Maximum Weight Closure):
//   이익이 있는 노드(프로젝트)는 소스에 연결, 비용이 있는 노드(직원)는 싱크에 연결
//   최대 순이익 = 모든 프로젝트 이익의 합 - 최소 컷(최대 유량)

const INF = 1<<63 - 1

// Edge는 유량 네트워크의 간선을 나타낸다
type Edge struct {
	to, cap, flow int
}

var (
	edges []Edge
	graph [][]int
	nodeN int
)

// addEdge는 u→v 간선과 역방향 간선을 쌍으로 추가한다
func addEdge(u, v, cap int) {
	graph[u] = append(graph[u], len(edges))
	edges = append(edges, Edge{v, cap, 0})
	graph[v] = append(graph[v], len(edges))
	edges = append(edges, Edge{u, 0, 0})
}

// bfs는 소스에서 싱크까지의 증가 경로를 BFS로 찾는다 (Edmonds-Karp)
func bfs(s, t int, parent []int) bool {
	for i := range parent {
		parent[i] = -1
	}
	parent[s] = s

	queue := []int{s}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, idx := range graph[u] {
			e := &edges[idx]
			if parent[e.to] == -1 && e.cap-e.flow > 0 {
				parent[e.to] = idx
				if e.to == t {
					return true
				}
				queue = append(queue, e.to)
			}
		}
	}
	return false
}

// maxFlow는 소스 s에서 싱크 t로의 최대 유량을 구한다 (Edmonds-Karp)
func maxFlow(s, t int) int {
	total := 0
	parent := make([]int, nodeN)

	for bfs(s, t, parent) {
		// 경로 상 최소 잔여 용량
		f := INF
		for v := t; v != s; {
			idx := parent[v]
			if edges[idx].cap-edges[idx].flow < f {
				f = edges[idx].cap - edges[idx].flow
			}
			v = edges[idx^1].to
		}

		// 유량 전송
		for v := t; v != s; {
			idx := parent[v]
			edges[idx].flow += f
			edges[idx^1].flow -= f
			v = edges[idx^1].to
		}

		total += f
	}
	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 프로젝트 수 N, 직원 수 M
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 입력: 각 프로젝트의 이익
	profit := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &profit[i])
	}

	// 입력: 각 직원의 급여
	salary := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &salary[i])
	}

	// 입력: 각 프로젝트가 요구하는 직원 목록
	requires := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(reader, &k)
		requires[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &requires[i][j])
		}
	}

	// 네트워크 구성 (최대 가중 폐합 모델)
	// 노드: S(0), 프로젝트(1~N), 직원(N+1~N+M), T(N+M+1)
	S := 0
	T := n + m + 1
	nodeN = n + m + 2
	edges = make([]Edge, 0)
	graph = make([][]int, nodeN)

	totalProfit := 0

	// S → 프로젝트 i (용량 = profit[i])
	// 이 간선이 컷에 포함되면 = 프로젝트를 포기한다는 의미
	for i := 0; i < n; i++ {
		addEdge(S, i+1, profit[i])
		totalProfit += profit[i]
	}

	// 직원 j → T (용량 = salary[j])
	// 이 간선이 컷에 포함되면 = 직원을 고용한다는 의미
	for j := 0; j < m; j++ {
		addEdge(n+1+j, T, salary[j])
	}

	// 프로젝트 i → 직원 j (용량 = INF)
	// 프로젝트를 선택하면 필요한 직원을 반드시 고용해야 함
	for i := 0; i < n; i++ {
		for _, j := range requires[i] {
			addEdge(i+1, n+j, INF) // 직원 번호는 1-indexed
		}
	}

	// 최대 순이익 = 전체 이익 - 최소 컷(최대 유량)
	minCut := maxFlow(S, T)
	result := totalProfit - minCut

	// 순이익이 음수이면 아무것도 선택하지 않는 것이 최적
	if result < 0 {
		result = 0
	}

	// 출력: 최대 순이익
	fmt.Fprintln(writer, result)
}
