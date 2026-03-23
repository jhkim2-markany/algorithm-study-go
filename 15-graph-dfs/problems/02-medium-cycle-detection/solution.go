package main

import (
	"bufio"
	"fmt"
	"os"
)

// 상태 상수: 미방문, 탐색 중, 탐색 완료
const (
	WHITE = 0 // 미방문
	GRAY  = 1 // 현재 DFS 경로에 포함 (탐색 중)
	BLACK = 2 // 탐색 완료
)

var adj [][]int
var color []int
var hasCycle bool

// dfs 함수는 방향 그래프에서 사이클을 탐색한다
func dfs(v int) {
	// 현재 정점을 탐색 중(GRAY)으로 표시
	color[v] = GRAY

	for _, next := range adj[v] {
		if hasCycle {
			return
		}
		if color[next] == GRAY {
			// 탐색 중인 정점을 다시 만남 → 역방향 간선 → 사이클 존재
			hasCycle = true
			return
		}
		if color[next] == WHITE {
			// 미방문 정점이면 재귀 탐색
			dfs(next)
		}
	}

	// 모든 인접 정점 탐색 완료 → BLACK으로 표시
	color[v] = BLACK
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정점 수와 간선 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화 (방향 그래프)
	adj = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력 (단방향)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
	}

	// 3색 DFS로 사이클 판별
	color = make([]int, n+1)
	hasCycle = false

	for i := 1; i <= n; i++ {
		if color[i] == WHITE {
			dfs(i)
			if hasCycle {
				break
			}
		}
	}

	// 결과 출력
	if hasCycle {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
