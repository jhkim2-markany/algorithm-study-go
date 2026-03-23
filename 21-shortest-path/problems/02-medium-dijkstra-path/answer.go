package main

import (
	"bufio"
	"container/heap"
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
//
// [알고리즘 힌트]
//
//	최소 힙 기반 우선순위 큐를 사용하여 Dijkstra 알고리즘을 구현한다.
//	거리 배열을 MaxInt64로 초기화하고, 시작 정점의 거리를 0으로 설정한다.
//	큐에서 최소 거리 정점을 꺼내 인접 정점에 대해 완화(relaxation)를 수행한다.
//	이미 확정된 거리보다 큰 값이 큐에서 나오면 건너뛴다.
//	경로 복원을 위해 각 정점의 이전 정점을 prev 배열에 기록한다.
func dijkstra(graph [][]Edge, n, start int) ([]int, []int) {
	dist := make([]int, n+1)
	prev := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt64
		prev[i] = -1
	}
	dist[start] = 0

	pq := &PQ{{0, start}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		u := cur.node

		if cur.dist > dist[u] {
			continue
		}

		for _, e := range graph[u] {
			newDist := dist[u] + e.weight
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				prev[e.to] = u
				heap.Push(pq, Item{newDist, e.to})
			}
		}
	}
	return dist, prev
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
