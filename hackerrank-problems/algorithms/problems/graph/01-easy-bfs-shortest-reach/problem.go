package main

import (
	"bufio"
	"fmt"
	"os"
)

// bfs는 시작 노드에서 모든 노드까지의 최단 거리를 반환한다.
// 각 간선의 가중치는 6이다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 원소는 [2]int{u, v})
//   - s: 시작 노드 (1-indexed)
//
// [반환값]
//   - []int: 시작 노드를 제외한 각 노드까지의 최단 거리 (-1은 도달 불가)
func bfs(n int, edges [][2]int, s int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		edges := make([][2]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &edges[i][0], &edges[i][1])
		}

		var s int
		fmt.Fscan(reader, &s)

		result := bfs(n, edges, s)
		for i, v := range result {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
