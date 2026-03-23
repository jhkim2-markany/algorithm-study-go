package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MOD     = 1000000007
	NEG_INF = -int(1e18)
)

// Edge는 가중치 간선을 나타낸다.
type Edge struct {
	to, weight int
}

// DAGResult는 DAG 경로 분석 결과를 담는다.
type DAGResult struct {
	totalPaths int
	maxDist    int
	maxPaths   int
}

// dagPathAnalysis는 DAG에서 경로 수, 최장 거리, 최장 경로 수를 구한다.
//
// [매개변수]
//   - n: 정점의 수 (0-indexed)
//   - graph: 인접 리스트 (graph[u] = u에서 나가는 간선 목록)
//   - s: 시작 정점 번호
//   - t: 도착 정점 번호
//
// [반환값]
//   - DAGResult: 전체 경로 수, 최장 거리, 최장 경로 수
//
// [알고리즘 힌트]
//
//	위상 정렬 후 DP로 세 가지 값을 동시에 계산한다.
//	totalPaths[v] += totalPaths[u] (경로 수).
//	maxDist[v] = max(maxDist[v], maxDist[u]+w) (최장 거리).
//	maxPaths[v]: 최장 거리 갱신 시 경로 수도 갱신.
//	시간복잡도: O(V + E)
func dagPathAnalysis(n int, graph [][]Edge, s, t int) DAGResult {
	inDegree := make([]int, n)
	for u := 0; u < n; u++ {
		for _, e := range graph[u] {
			inDegree[e.to]++
		}
	}

	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := []int{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, e := range graph[u] {
			inDegree[e.to]--
			if inDegree[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	totalPaths := make([]int, n)
	totalPaths[s] = 1

	maxDist := make([]int, n)
	for i := range maxDist {
		maxDist[i] = NEG_INF
	}
	maxDist[s] = 0

	maxPaths := make([]int, n)
	maxPaths[s] = 1

	for _, u := range order {
		for _, e := range graph[u] {
			v, w := e.to, e.weight

			if totalPaths[u] > 0 {
				totalPaths[v] = (totalPaths[v] + totalPaths[u]) % MOD
			}

			if maxDist[u] != NEG_INF {
				newDist := maxDist[u] + w
				if newDist > maxDist[v] {
					maxDist[v] = newDist
					maxPaths[v] = maxPaths[u]
				} else if newDist == maxDist[v] {
					maxPaths[v] = (maxPaths[v] + maxPaths[u]) % MOD
				}
			}
		}
	}

	if totalPaths[t] == 0 {
		return DAGResult{0, 0, 0}
	}
	return DAGResult{totalPaths[t], maxDist[t], maxPaths[t]}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	graph := make([][]Edge, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		graph[u] = append(graph[u], Edge{v, w})
	}

	result := dagPathAnalysis(n, graph, s, t)
	if result.totalPaths == 0 {
		fmt.Fprintln(writer, 0)
		fmt.Fprintln(writer, 0)
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, result.totalPaths)
		fmt.Fprintln(writer, result.maxDist)
		fmt.Fprintln(writer, result.maxPaths)
	}
}
