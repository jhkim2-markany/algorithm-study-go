package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 유량 네트워크의 간선을 나타낸다.
type Edge struct {
	to, cap, flow, rev int
}

// maxFlow는 소스에서 싱크까지의 최대 유량을 구한다.
//
// [매개변수]
//   - n: 정점의 수 (1-indexed)
//   - edges: 간선 목록 (각 원소는 [u, v, cap])
//   - source: 소스 정점 번호
//   - sink: 싱크 정점 번호
//
// [반환값]
//   - int: 소스에서 싱크까지의 최대 유량
func maxFlow(n int, edges [][3]int, source, sink int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	fmt.Fprintln(writer, maxFlow(n, edges, 1, n))
}
