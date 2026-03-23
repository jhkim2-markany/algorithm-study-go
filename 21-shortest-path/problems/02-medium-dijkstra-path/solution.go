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

	// Dijkstra 알고리즘으로 최단 거리 및 경로 추적
	dist := make([]int, n+1)
	prev := make([]int, n+1) // 경로 복원을 위한 이전 정점 배열
	for i := range dist {
		dist[i] = math.MaxInt64
		prev[i] = -1
	}
	dist[1] = 0

	pq := &PQ{{0, 1}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		u := cur.node

		// 이미 더 짧은 경로가 확정되었으면 건너뛴다
		if cur.dist > dist[u] {
			continue
		}

		// 인접 정점에 대해 완화(relaxation) 수행
		for _, e := range graph[u] {
			newDist := dist[u] + e.weight
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				prev[e.to] = u // 이전 정점 기록
				heap.Push(pq, Item{newDist, e.to})
			}
		}
	}

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
