package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// preorderTraversal은 루트가 root인 트리를 전위 순회한 결과를 반환한다.
//
// [매개변수]
//   - adj: 인접 리스트 (1-indexed)
//   - root: 시작 노드 번호
//   - n: 노드 수
//
// [반환값]
//   - []int: 전위 순회 결과 (노드 번호 순서)
//
// [알고리즘 힌트]
//
//	전위 순회는 현재 노드를 먼저 방문한 뒤 자식 노드를 재귀적으로 방문한다.
//	부모 노드를 추적하여 역방향 탐색을 방지한다.
//	자식 노드를 번호 순서대로 방문하기 위해 인접 리스트를 미리 정렬해둔다.
func preorderTraversal(adj [][]int, root, n int) []int {
	var result []int
	var dfs func(cur, par int)
	dfs = func(cur, par int) {
		result = append(result, cur)
		for _, next := range adj[cur] {
			if next != par {
				dfs(next, cur)
			}
		}
	}
	dfs(root, 0)
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 인접 리스트 초기화
	adj := make([][]int, n+1)
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

	// 자식 노드를 번호 순서대로 방문하기 위해 정렬
	for i := 1; i <= n; i++ {
		sort.Ints(adj[i])
	}

	// 핵심 함수 호출
	result := preorderTraversal(adj, 1, n)

	// 결과 출력
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
