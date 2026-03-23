package main

import (
	"bufio"
	"fmt"
	"os"
)

// countConnectedComponents는 무방향 그래프의 연결 요소 개수를 반환한다.
//
// [매개변수]
//   - adj: 인접 리스트 (1-indexed)
//   - n: 정점 수
//
// [반환값]
//   - int: 연결 요소의 개수
func countConnectedComponents(adj [][]int, n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정점 수와 간선 수 입력
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
	result := countConnectedComponents(adj, n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
