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

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		// 인접 리스트 초기화
		adj := make([][]int, n+1)
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

		// 색칠 배열: 0은 미방문, 1과 2는 두 그룹을 나타낸다
		color := make([]int, n+1)
		bipartite := true

		// 모든 정점에 대해 BFS 수행 (비연결 그래프 처리)
		for i := 1; i <= n && bipartite; i++ {
			if color[i] != 0 {
				continue
			}

			// BFS로 2-색칠 시도
			queue := []int{i}
			color[i] = 1

			for len(queue) > 0 && bipartite {
				v := queue[0]
				queue = queue[1:]

				for _, next := range adj[v] {
					if color[next] == 0 {
						// 미방문 정점에 반대 색을 칠한다
						if color[v] == 1 {
							color[next] = 2
						} else {
							color[next] = 1
						}
						queue = append(queue, next)
					} else if color[next] == color[v] {
						// 인접한 정점이 같은 색이면 이분 그래프가 아니다
						bipartite = false
					}
				}
			}
		}

		if bipartite {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
