package main

import (
	"bufio"
	"fmt"
	"os"
)

// bfs는 시작 노드에서 모든 노드까지의 최단 거리를 반환한다.
// 각 간선의 가중치는 6이다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 원소는 [2]int{u, v})
//   - s: 시작 노드 (1-indexed)
//
// [반환값]
//   - []int: 시작 노드를 제외한 각 노드까지의 최단 거리 (-1은 도달 불가)
//
// [알고리즘 힌트]
//
//	인접 리스트를 구성한 뒤 BFS로 최단 거리를 계산한다.
//	간선 가중치가 6으로 동일하므로 BFS 레벨 × 6이 최단 거리이다.
func bfs(n int, edges [][2]int, s int) []int {
	// 인접 리스트 구성
	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 거리 배열 초기화 (-1: 미방문)
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = -1
	}
	dist[s] = 0

	// BFS 수행
	queue := []int{s}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range adj[cur] {
			if dist[next] == -1 {
				dist[next] = dist[cur] + 6
				queue = append(queue, next)
			}
		}
	}

	// 시작 노드를 제외한 결과 구성
	result := make([]int, 0, n-1)
	for i := 1; i <= n; i++ {
		if i != s {
			result = append(result, dist[i])
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		edges := make([][2]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &edges[i][0], &edges[i][1])
		}

		var s int
		fmt.Fscan(reader, &s)

		result := bfs(n, edges, s)
		for i, v := range result {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
