package main

import (
	"bufio"
	"fmt"
	"os"
)

// countSCC는 코사라주 알고리즘으로 강한 연결 요소의 개수를 구한다.
//
// [매개변수]
//   - adj: 인접 리스트로 표현된 방향 그래프 (1-indexed)
//   - radj: 역방향 인접 리스트 (1-indexed)
//   - n: 정점의 수
//
// [반환값]
//   - int: 강한 연결 요소(SCC)의 개수
func countSCC(adj, radj [][]int, n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화 (원본 + 역방향)
	adj := make([][]int, n+1)
	radj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []int{}
		radj[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		radj[v] = append(radj[v], u)
	}

	// 핵심 함수 호출
	result := countSCC(adj, radj, n)

	fmt.Fprintln(writer, result)
}
