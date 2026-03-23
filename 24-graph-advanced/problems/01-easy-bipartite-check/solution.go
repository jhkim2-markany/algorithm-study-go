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

	// 색칠 배열: -1은 미색칠, 0과 1은 두 그룹
	color := make([]int, n+1)
	for i := range color {
		color[i] = -1
	}

	bipartite := true

	// 모든 연결 요소에 대해 BFS 색칠 수행
	for start := 1; start <= n; start++ {
		if color[start] != -1 {
			continue
		}

		// BFS로 시작 정점부터 색칠
		color[start] = 0
		queue := []int{start}

		for len(queue) > 0 && bipartite {
			cur := queue[0]
			queue = queue[1:]

			for _, next := range adj[cur] {
				if color[next] == -1 {
					// 인접 정점에 반대 색을 칠한다
					color[next] = 1 - color[cur]
					queue = append(queue, next)
				} else if color[next] == color[cur] {
					// 같은 색이면 이분 그래프가 아니다
					bipartite = false
					break
				}
			}
		}

		if !bipartite {
			break
		}
	}

	if bipartite {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
