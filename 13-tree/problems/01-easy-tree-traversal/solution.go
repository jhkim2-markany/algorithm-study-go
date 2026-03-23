package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var adj [][]int
var result []int

// preorder 함수는 전위 순회를 수행한다 (현재 노드 → 자식 노드)
func preorder(cur, par int) {
	// 현재 노드를 먼저 방문
	result = append(result, cur)
	// 자식 노드를 번호 순서대로 방문
	for _, next := range adj[cur] {
		if next != par {
			preorder(next, cur)
		}
	}
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

	// 루트(1번)에서 전위 순회 시작
	result = []int{}
	preorder(1, 0)

	// 결과 출력
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
