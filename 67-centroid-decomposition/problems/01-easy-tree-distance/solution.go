package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// countPairsWithinK는 센트로이드 분할을 이용하여 트리에서
// 거리가 K 이하인 노드 쌍의 수를 구한다.
//
// [매개변수]
//   - n: 노드 수
//   - k: 거리 제한
//   - edges: 간선 목록 (u, v 쌍)
//
// [반환값]
//   - int64: 거리가 K 이하인 노드 쌍의 수
func countPairsWithinK(n, k int, edges [][2]int) int64 {
	// 여기에 코드를 작성하세요
	_ = sort.Ints
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	fmt.Fprintln(writer, countPairsWithinK(n, k, edges))
}
