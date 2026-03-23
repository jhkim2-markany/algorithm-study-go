package main

import (
	"bufio"
	"fmt"
	"os"
)

// isBipartite는 그래프가 이분 그래프인지 판별한다.
//
// [매개변수]
//   - adj: 인접 리스트로 표현된 무방향 그래프 (1-indexed)
//   - n: 정점의 수
//
// [반환값]
//   - bool: 이분 그래프이면 true, 아니면 false
//
// [알고리즘 힌트]
//
//	BFS 색칠법을 사용한다. 각 정점에 0 또는 1의 색을 칠하되,
//	인접한 정점에는 반대 색을 칠한다. 인접한 두 정점이 같은 색이면
//	이분 그래프가 아니다. 비연결 그래프를 처리하기 위해
//	모든 연결 요소에 대해 BFS를 수행한다.
func isBipartite(adj [][]int, n int) bool {
	color := make([]int, n+1)
	for i := range color {
		color[i] = -1
	}

	for start := 1; start <= n; start++ {
		if color[start] != -1 {
			continue
		}

		color[start] = 0
		queue := []int{start}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for _, next := range adj[cur] {
				if color[next] == -1 {
					color[next] = 1 - color[cur]
					queue = append(queue, next)
				} else if color[next] == color[cur] {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화
	adj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력 (무방향)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 핵심 함수 호출
	if isBipartite(adj, n) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
