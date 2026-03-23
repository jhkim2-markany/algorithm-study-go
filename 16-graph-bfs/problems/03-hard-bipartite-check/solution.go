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
func isBipartite(adj [][]int, n int) bool {
	// 여기에 코드를 작성하세요
	return false
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
