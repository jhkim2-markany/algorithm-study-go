package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트로 그래프 구성
	graph := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// BFS로 최단 거리 계산 (비가중치 그래프)
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = -1
	}
	dist[1] = 0

	queue := []int{1}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// 인접 정점 탐색
		for _, v := range graph[u] {
			// 아직 방문하지 않은 정점이면 거리 갱신
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				queue = append(queue, v)
			}
		}
	}

	// 결과 출력
	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, dist[i])
	}
}
