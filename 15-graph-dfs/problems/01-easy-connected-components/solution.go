package main

import (
	"bufio"
	"fmt"
	"os"
)

var adj [][]int
var visited []bool

// dfs 함수는 현재 정점에서 연결된 모든 정점을 방문한다
func dfs(v int) {
	visited[v] = true
	for _, next := range adj[v] {
		if !visited[next] {
			dfs(next)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정점 수와 간선 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화
	adj = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력 (양방향)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 모든 정점을 순회하며 연결 요소 개수를 센다
	visited = make([]bool, n+1)
	count := 0
	for i := 1; i <= n; i++ {
		if !visited[i] {
			// 미방문 정점에서 DFS 시작 → 새로운 연결 요소 발견
			dfs(i)
			count++
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, count)
}
