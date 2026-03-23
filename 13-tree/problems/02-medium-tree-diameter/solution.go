package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 가중치가 있는 간선을 나타낸다.
type Edge struct {
	to, weight int
}

// treeDiameter는 가중치가 있는 트리의 지름(가장 먼 두 노드 사이의 거리)을 반환한다.
//
// [매개변수]
//   - adj: 가중치 간선의 인접 리스트 (1-indexed)
//   - n: 노드 수
//
// [반환값]
//   - int: 트리의 지름
func treeDiameter(adj [][]Edge, n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 인접 리스트 초기화
	adj := make([][]Edge, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []Edge{}
	}

	// 간선 입력 (가중치 포함)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], Edge{v, w})
		adj[v] = append(adj[v], Edge{u, w})
	}

	// 핵심 함수 호출
	result := treeDiameter(adj, n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
