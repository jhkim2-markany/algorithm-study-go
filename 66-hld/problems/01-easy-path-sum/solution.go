package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathSum은 HLD(Heavy-Light Decomposition)와 세그먼트 트리를 이용하여
// 트리에서 두 노드 사이의 경로 합 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - val: 각 노드의 값 (1-indexed)
//   - edges: 간선 목록 (u, v 쌍)
//   - queries: 질의 목록 (u, v 쌍)
//
// [반환값]
//   - []int: 각 질의에 대한 경로 합 결과
func pathSum(n int, val []int, edges [][2]int, queries [][2]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	val := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &val[i])
	}

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	queries := make([][2]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	results := pathSum(n, val, edges, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
