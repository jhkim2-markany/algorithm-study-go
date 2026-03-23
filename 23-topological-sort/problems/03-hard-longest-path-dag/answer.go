package main

import (
	"bufio"
	"fmt"
	"os"
)

const NEG_INF = -1 << 60

// Edge는 가중치 간선을 나타낸다
type Edge struct {
	to, weight int
}

// longestPathDAG는 DAG에서 시작 정점으로부터 각 정점까지의 최장 거리를 구한다.
//
// [매개변수]
//   - adj: 인접 리스트로 표현된 가중치 DAG (1-indexed)
//   - inDegree: 각 정점의 진입 차수 배열
//   - n: 정점의 수
//   - s: 시작 정점 번호
//
// [반환값]
//   - []int: 각 정점까지의 최장 거리 배열 (도달 불가능하면 NEG_INF)
//
// [알고리즘 힌트]
//
//	Kahn 알고리즘으로 위상 정렬을 수행한 후,
//	위상 정렬 순서대로 DP를 갱신한다.
//	dist[s] = 0으로 초기화하고, 나머지는 NEG_INF로 설정한다.
//	각 정점에서 나가는 간선을 통해 dist[u] + weight > dist[v]이면 갱신한다.
//	도달 불가능한 정점은 NEG_INF 상태로 남는다.
func longestPathDAG(adj [][]Edge, inDegree []int, n, s int) []int {
	// 위상 정렬
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := make([]int, 0, n)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		order = append(order, cur)

		for _, e := range adj[cur] {
			inDegree[e.to]--
			if inDegree[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	// 최장 경로 DP
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = NEG_INF
	}
	dist[s] = 0

	for _, u := range order {
		if dist[u] == NEG_INF {
			continue
		}
		for _, e := range adj[u] {
			if dist[u]+e.weight > dist[e.to] {
				dist[e.to] = dist[u] + e.weight
			}
		}
	}

	return dist
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s int
	fmt.Fscan(reader, &n, &m, &s)

	// 인접 리스트와 진입 차수 배열 초기화
	adj := make([][]Edge, n+1)
	inDegree := make([]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []Edge{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], Edge{v, w})
		inDegree[v]++
	}

	// 핵심 함수 호출
	dist := longestPathDAG(adj, inDegree, n, s)

	// 결과 출력
	for i := 1; i <= n; i++ {
		if dist[i] == NEG_INF {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, dist[i])
		}
	}
}
