package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXLOG = 17

var (
	adj   [][]int
	depth []int
	up    [MAXLOG][]int
)

// buildLCA는 트리의 깊이와 희소 테이블을 구축한다.
//
// [매개변수]
//   - adj: 인접 리스트 (1-indexed)
//   - n: 노드 수
//   - root: 루트 노드 번호
func buildLCA(adj [][]int, n, root int) {
	// 여기에 코드를 작성하세요
}

// queryLCA는 두 노드의 최소 공통 조상을 반환한다.
//
// [매개변수]
//   - u: 첫 번째 노드 번호
//   - v: 두 번째 노드 번호
//
// [반환값]
//   - int: 두 노드의 최소 공통 조상 노드 번호
func queryLCA(u, v int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 인접 리스트 초기화
	adj = make([][]int, n+1)
	depth = make([]int, n+1)
	for k := 0; k < MAXLOG; k++ {
		up[k] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// LCA 전처리
	buildLCA(adj, n, 1)

	// 쿼리 처리
	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintln(writer, queryLCA(a, b))
	}
}
