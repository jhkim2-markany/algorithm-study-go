package main

import (
	"bufio"
	"fmt"
	"os"
)

// coloredPathQuery는 센트로이드 분할로 거리별 같은 색상 쌍의 수를 전처리하고,
// 각 쿼리에 대해 거리 k에서의 같은 색상 쌍 수를 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - color: 각 노드의 색상 (1-indexed)
//   - edges: 간선 목록 (u, v 쌍)
//   - queries: 거리 쿼리 목록
//
// [반환값]
//   - []int64: 각 쿼리에 대한 결과 (거리 k에서 같은 색상인 쌍의 수)
func coloredPathQuery(n int, color []int, edges [][2]int, queries []int) []int64 {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	color := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &color[i])
	}

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	queries := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i])
	}

	results := coloredPathQuery(n, color, edges, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
