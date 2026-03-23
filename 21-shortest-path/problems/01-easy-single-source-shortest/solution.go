package main

import (
	"bufio"
	"fmt"
	"os"
)

// bfsShortestPath는 비가중치 그래프에서 시작 정점으로부터 모든 정점까지의 최단 거리를 구한다.
//
// [매개변수]
//   - graph: 인접 리스트로 표현된 그래프 (1-indexed)
//   - n: 정점의 수
//   - start: 시작 정점 번호
//
// [반환값]
//   - []int: 각 정점까지의 최단 거리 배열 (도달 불가능하면 -1)
func bfsShortestPath(graph [][]int, n, start int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트로 그래프 구성
	graph := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// 핵심 함수 호출
	dist := bfsShortestPath(graph, n, 1)

	// 결과 출력
	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, dist[i])
	}
}
