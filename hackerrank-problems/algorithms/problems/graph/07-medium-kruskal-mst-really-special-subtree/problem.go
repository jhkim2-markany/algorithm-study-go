package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Edge는 그래프의 간선을 나타낸다.
type Edge struct {
	U, V, W int
}

// kruskalMST는 크루스칼 알고리즘으로 MST의 가중치 합을 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록
//
// [반환값]
//   - int: MST의 가중치 합
func kruskalMST(n int, edges []Edge) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i].U, &edges[i].V, &edges[i].W)
	}

	var s int
	fmt.Fscan(reader, &s)
	_ = s

	result := kruskalMST(n, edges)
	fmt.Fprintln(writer, result)

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Slice
}
