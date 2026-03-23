package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to, weight int
}

// treeDiameter는 가중치 트리의 지름(가장 먼 두 노드 사이의 거리)을 반환한다.
//
// [매개변수]
//   - n: 노드의 수
//   - adj: 가중치 간선의 인접 리스트
//
// [반환값]
//   - int: 트리의 지름
func treeDiameter(n int, adj [][]edge) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	adj := make([][]edge, n+1)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}

	fmt.Fprintln(writer, treeDiameter(n, adj))
}
