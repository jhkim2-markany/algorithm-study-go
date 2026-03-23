package main

import (
	"bufio"
	"fmt"
	"os"
)

// hasCycleInDirectedGraph는 방향 그래프에 사이클이 존재하는지 판별한다.
//
// [매개변수]
//   - adj: 인접 리스트 (1-indexed, 방향 그래프)
//   - n: 정점 수
//
// [반환값]
//   - bool: 사이클이 존재하면 true, 아니면 false
//
// [알고리즘 힌트]
//
//	3색 DFS를 사용한다.
//	WHITE(0): 미방문, GRAY(1): 현재 DFS 경로에 포함(탐색 중), BLACK(2): 탐색 완료.
//	탐색 중(GRAY) 정점을 다시 만나면 역방향 간선이므로 사이클이 존재한다.
//	모든 정점에서 DFS를 시작하여 비연결 그래프도 처리한다.
func hasCycleInDirectedGraph(adj [][]int, n int) bool {
	const (
		WHITE = 0
		GRAY  = 1
		BLACK = 2
	)

	color := make([]int, n+1)
	hasCycle := false

	var dfs func(v int)
	dfs = func(v int) {
		color[v] = GRAY
		for _, next := range adj[v] {
			if hasCycle {
				return
			}
			if color[next] == GRAY {
				hasCycle = true
				return
			}
			if color[next] == WHITE {
				dfs(next)
			}
		}
		color[v] = BLACK
	}

	for i := 1; i <= n; i++ {
		if color[i] == WHITE {
			dfs(i)
			if hasCycle {
				break
			}
		}
	}

	return hasCycle
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정점 수와 간선 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화 (방향 그래프)
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력 (단방향)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
	}

	// 핵심 함수 호출
	if hasCycleInDirectedGraph(adj, n) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
