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
//
// [알고리즘 힌트]
//   - 루트(1번)에서 후위 순회(DFS)를 수행한다
//   - 각 노드의 서브트리 크기 = 1(자기 자신) + 모든 자식 서브트리 크기의 합
//   - 부모 방향으로의 역행을 방지하기 위해 parent 매개변수를 사용한다
func subtreeSize(n int, adj [][]int) []int {
	sz := make([]int, n+1)

	var dfs func(v, parent int)
	dfs = func(v, parent int) {
		sz[v] = 1
		for _, u := range adj[v] {
			if u == parent {
				continue
			}
			dfs(u, v)
			sz[v] += sz[u]
		}
	}

	dfs(1, 0)
	return sz
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
