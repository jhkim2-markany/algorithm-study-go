package main

import (
	"bufio"
	"fmt"
	"os"
)

// subtreeSize는 루트에서 DFS를 수행하여 각 노드의 서브트리 크기를 계산한다.
//
// [매개변수]
//   - n: 노드의 수
//   - adj: 인접 리스트
//
// [반환값]
//   - []int: 각 노드의 서브트리 크기 배열 (1-indexed)
func subtreeSize(n int, adj [][]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	adj := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := subtreeSize(n, adj)

	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, sz[i])
	}
	fmt.Fprintln(writer)
}
