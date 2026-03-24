package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

// RaptureItem은 우선순위 큐의 원소이다.
type RaptureItem struct {
	node, maxWeight int
}

// RapturePQ는 최소 힙 기반 우선순위 큐이다.
type RapturePQ []RaptureItem

func (pq RapturePQ) Len() int            { return len(pq) }
func (pq RapturePQ) Less(i, j int) bool  { return pq[i].maxWeight < pq[j].maxWeight }
func (pq RapturePQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *RapturePQ) Push(x interface{}) { *pq = append(*pq, x.(RaptureItem)) }
func (pq *RapturePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// jackGoesToRapture는 1번에서 N번까지 경로 상 최대 요금의 최솟값을 반환한다.
//
// [매개변수]
//   - n: 역 수
//   - edges: 노선 목록 (각 원소는 [3]int{u, v, w})
//
// [반환값]
//   - int: 최대 요금의 최솟값 (-1이면 도달 불가)
//
// [알고리즘 힌트]
//
//	변형 다익스트라로 미니맥스 경로를 구한다.
//	dist[v] = min(dist[v], max(dist[u], w(u,v)))로 갱신한다.
func jackGoesToRapture(n int, edges [][3]int) int {
	// 인접 리스트 구성
	type Edge struct {
		to, w int
	}
	adj := make([][]Edge, n+1)
	for i := range adj {
		adj[i] = []Edge{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], Edge{e[0], e[2]})
	}

	// 최대 가중치 배열 초기화
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[1] = 0

	// 변형 다익스트라 수행
	pq := &RapturePQ{{1, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(RaptureItem)
		// 이미 더 좋은 경로가 확정된 경우 건너뛰기
		if cur.maxWeight > dist[cur.node] {
			continue
		}
		// 목적지 도달
		if cur.node == n {
			return dist[n]
		}
		for _, e := range adj[cur.node] {
			// 경로 상 최대 가중치 계산
			newMax := cur.maxWeight
			if e.w > newMax {
				newMax = e.w
			}
			if newMax < dist[e.to] {
				dist[e.to] = newMax
				heap.Push(pq, RaptureItem{e.to, newMax})
			}
		}
	}

	return -1
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

	result := jackGoesToRapture(n, edges)
	if result == -1 {
		fmt.Fprintln(writer, "NO PATH EXISTS")
	} else {
		fmt.Fprintln(writer, result)
	}
}
