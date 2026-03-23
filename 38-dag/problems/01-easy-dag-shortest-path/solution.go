package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = int(1e18)

// 간선 구조체
type Edge struct {
	to, weight int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 정점 수, 간선 수, 시작 정점
	var n, m, s int
	fmt.Fscan(reader, &n, &m, &s)

	// 인접 리스트 구성
	graph := make([][]Edge, n)
	inDegree := make([]int, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		graph[u] = append(graph[u], Edge{v, w})
		inDegree[v]++
	}

	// 위상 정렬 (Kahn's Algorithm)
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

	// 최단 거리 초기화
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[s] = 0

	// 위상 정렬 순서대로 간선 완화
	for _, u := range order {
		if dist[u] == INF {
			continue
		}
		for _, e := range graph[u] {
			// 최단 거리 갱신
			if dist[u]+e.weight < dist[e.to] {
				dist[e.to] = dist[u] + e.weight
			}
		}
	}

	// 결과 출력
	for i := 0; i < n; i++ {
		if dist[i] == INF {
			fmt.Fprintln(writer, "INF")
		} else {
			fmt.Fprintln(writer, dist[i])
		}
	}
}
