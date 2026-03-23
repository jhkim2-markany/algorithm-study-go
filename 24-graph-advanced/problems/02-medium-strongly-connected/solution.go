package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m    int
	adj     [][]int
	radj    [][]int
	visited []bool
	order   []int
)

// dfs1 함수는 원본 그래프에서 DFS를 수행하며 완료 순서를 기록한다
func dfs1(u int) {
	visited[u] = true
	for _, v := range adj[u] {
		if !visited[v] {
			dfs1(v)
		}
	}
	order = append(order, u)
}

// dfs2 함수는 역방향 그래프에서 DFS를 수행하여 SCC를 구한다
func dfs2(u int) {
	visited[u] = true
	for _, v := range radj[u] {
		if !visited[v] {
			dfs2(v)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화 (원본 + 역방향)
	adj = make([][]int, n+1)
	radj = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []int{}
		radj[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		radj[v] = append(radj[v], u) // 역방향 간선
	}

	// 1단계: 원본 그래프에서 DFS, 완료 순서 기록
	visited = make([]bool, n+1)
	order = []int{}
	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs1(i)
		}
	}

	// 2단계: 역방향 그래프에서 완료 순서의 역순으로 DFS
	for i := range visited {
		visited[i] = false
	}

	sccCount := 0
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		if !visited[u] {
			dfs2(u)
			sccCount++ // 하나의 SCC 완성
		}
	}

	fmt.Fprintln(writer, sccCount)
}
