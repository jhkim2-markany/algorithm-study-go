package main

import (
	"fmt"
	"math"
)

// Bellman-Ford 알고리즘 - 음수 가중치를 허용하는 단일 출발점 최단 경로
// 시간 복잡도: O(V × E)
// 공간 복잡도: O(V)

// Edge는 간선을 나타낸다 (출발, 도착, 가중치)
type Edge struct {
	from, to, weight int
}

// bellmanFord 함수는 출발 정점에서 모든 정점까지의 최단 거리를 반환한다
// 음수 사이클이 존재하면 두 번째 반환값이 true이다
func bellmanFord(n int, edges []Edge, start int) ([]int, bool) {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0

	// (V-1)번 반복하여 최단 거리를 확정한다
	for i := 0; i < n-1; i++ {
		updated := false
		for _, e := range edges {
			// 출발 정점에 도달할 수 없으면 건너뛴다
			if dist[e.from] == math.MaxInt64 {
				continue
			}
			// 완화(relaxation): 더 짧은 경로가 발견되면 갱신
			newDist := dist[e.from] + e.weight
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				updated = true
			}
		}
		// 갱신이 없으면 조기 종료
		if !updated {
			break
		}
	}

	// V번째 반복에서 갱신이 발생하면 음수 사이클 존재
	hasNegCycle := false
	for _, e := range edges {
		if dist[e.from] == math.MaxInt64 {
			continue
		}
		if dist[e.from]+e.weight < dist[e.to] {
			hasNegCycle = true
			break
		}
	}

	return dist, hasNegCycle
}

func main() {
	// 그래프 생성 (정점 5개, 0~4)
	// 음수 가중치 간선 포함
	edges := []Edge{
		{0, 1, 4},
		{0, 2, 5},
		{1, 2, -3},
		{2, 3, 4},
		{3, 4, 2},
		{1, 4, 6},
	}

	// 정점 0에서 출발하는 최단 거리 계산
	dist, hasNegCycle := bellmanFord(5, edges, 0)

	fmt.Println("=== Bellman-Ford 최단 경로 ===")
	if hasNegCycle {
		fmt.Println("음수 사이클이 존재합니다!")
	} else {
		for i, d := range dist {
			if d == math.MaxInt64 {
				fmt.Printf("0 → %d: 도달 불가\n", i)
			} else {
				fmt.Printf("0 → %d: %d\n", i, d)
			}
		}
	}
}
