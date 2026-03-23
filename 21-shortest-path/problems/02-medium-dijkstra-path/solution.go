package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Edge는 인접 정점과 가중치를 나타낸다
type Edge struct {
	to, weight int
}

// Item은 우선순위 큐에 저장되는 (거리, 정점) 쌍이다
type Item struct {
	dist, node int
}

// PQ는 최소 힙 기반 우선순위 큐이다
type PQ []Item

func (pq PQ) Len() int            { return len(pq) }
func (pq PQ) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// dijkstra는 가중치 그래프에서 시작 정점으로부터 모든 정점까지의 최단 거리와 경로를 구한다.
//
// [매개변수]
//   - graph: 인접 리스트로 표현된 가중치 그래프 (1-indexed)
//   - n: 정점의 수
//   - start: 시작 정점 번호
//
// [반환값]
//   - []int: 각 정점까지의 최단 거리 배열
//   - []int: 경로 복원을 위한 이전 정점 배열
func dijkstra(graph [][]Edge, n, start int) ([]int, []int) {
	// 여기에 코드를 작성하세요
	return nil, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트로 그래프 구성
	graph := make([][]Edge, n+1)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		graph[u] = append(graph[u], Edge{v, w})
	}

	// 핵심 함수 호출
	dist, prev := dijkstra(graph, n, 1)

	// 도달 불가능한 경우
	if dist[n] == math.MaxInt64 {
		fmt.Fprintln(writer, -1)
		return
	}

	// 최단 거리 출력
	fmt.Fprintln(writer, dist[n])

	// 경로 복원: 도착점에서 출발점까지 역추적
	path := []int{}
	for v := n; v != -1; v = prev[v] {
		path = append(path, v)
	}

	// 경로를 뒤집어서 출발점부터 출력
	fmt.Fprintln(writer, len(path))
	for i := len(path) - 1; i >= 0; i-- {
		if i < len(path)-1 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, path[i])
	}
	fmt.Fprintln(writer)
}
