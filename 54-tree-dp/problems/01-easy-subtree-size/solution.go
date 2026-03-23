package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	adj [][]int // 인접 리스트
	sz  []int   // 서브트리 크기
)

// dfs는 후위 순회로 서브트리 크기를 계산한다
func dfs(v, parent int) {
	sz[v] = 1
	for _, u := range adj[v] {
		if u == parent {
			continue // 부모 방향 역행 방지
		}
		dfs(u, v)
		sz[v] += sz[u] // 자식 서브트리 크기 합산
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수
	var n int
	fmt.Fscan(reader, &n)

	adj = make([][]int, n+1)
	sz = make([]int, n+1)

	// 입력: 간선 정보
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 루트(1번)에서 DFS 수행
	dfs(1, 0)

	// 출력: 각 노드의 서브트리 크기
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, sz[i])
	}
	fmt.Fprintln(writer)
}
