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
func hasCycleInDirectedGraph(adj [][]int, n int) bool {
	// 여기에 코드를 작성하세요
	return false
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
