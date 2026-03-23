package main

import (
	"bufio"
	"fmt"
	"os"
)

// isBipartite는 무방향 그래프가 이분 그래프인지 판별한다.
//
// [매개변수]
//   - adj: 인접 리스트 (1-indexed)
//   - n: 정점 수
//
// [반환값]
//   - bool: 이분 그래프이면 true, 아니면 false
//
// [알고리즘 힌트]
//
//	BFS로 2-색칠을 시도한다.
//	미방문 정점에 색 1을 칠하고 BFS를 시작한다.
//	인접한 미방문 정점에 반대 색을 칠하고,
//	인접한 정점이 같은 색이면 이분 그래프가 아니다.
//	비연결 그래프를 처리하기 위해 모든 정점에서 BFS를 시도한다.
func isBipartite(adj [][]int, n int) bool {
	color := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if color[i] != 0 {
			continue
		}

		queue := []int{i}
		color[i] = 1

		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]

			for _, next := range adj[v] {
				if color[next] == 0 {
					if color[v] == 1 {
						color[next] = 2
					} else {
						color[next] = 1
					}
					queue = append(queue, next)
				} else if color[next] == color[v] {
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

		// 핵심 함수 호출
		if isBipartite(adj, n) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
