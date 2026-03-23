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
func isBipartite(adj [][]int, n int) bool {
	// 여기에 코드를 작성하세요
	return false
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
