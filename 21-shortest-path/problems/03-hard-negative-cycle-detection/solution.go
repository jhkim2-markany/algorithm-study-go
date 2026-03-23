package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 간선을 나타낸다 (출발, 도착, 가중치)
type Edge struct {
	from, to, weight int
}

// hasNegativeCycle은 그래프에서 음수 사이클의 존재 여부를 판별한다.
//
// [매개변수]
//   - n: 정점의 수
//   - edges: 간선 목록 (출발, 도착, 가중치)
//
// [반환값]
//   - bool: 음수 사이클이 존재하면 true, 아니면 false
func hasNegativeCycle(n int, edges []Edge) bool {
	// 여기에 코드를 작성하세요
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		edges := make([]Edge, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &edges[i].from, &edges[i].to, &edges[i].weight)
		}

		// 핵심 함수 호출
		if hasNegativeCycle(n, edges) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
