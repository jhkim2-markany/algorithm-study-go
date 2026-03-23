package main

import "fmt"

// 최소 비용 최대 유량 (MCMF) - SPFA 기반 구현
// 잔여 그래프에서 SPFA로 최소 비용 증가 경로를 찾아 유량을 보낸다.
// 시간 복잡도: O(V × E × F)
// 공간 복잡도: O(V + E)

const INF = 1<<63 - 1

// Edge는 유량 네트워크의 간선을 나타낸다
type Edge struct {
	to, cap, cost, flow int
}

var (
	edges []Edge  // 모든 간선 목록 (정방향, 역방향 쌍으로 저장)
	graph [][]int // graph[u] = u에서 나가는 간선의 인덱스 목록
	n     int     // 정점 수
)

// addEdge는 u→v 간선(용량 cap, 비용 cost)과 역방향 간선을 추가한다
func addEdge(u, v, cap, cost int) {
	graph[u] = append(graph[u], len(edges))
	edges = append(edges, Edge{v, cap, cost, 0})
	graph[v] = append(graph[v], len(edges))
	edges = append(edges, Edge{u, 0, -cost, 0})
}

// mcmf는 소스 s에서 싱크 t로 최소 비용 최대 유량을 구한다
// 반환값: (최대 유량, 최소 비용)
func mcmf(s, t int) (int, int) {
	totalFlow := 0
	totalCost := 0

	for {
		// SPFA로 s→t 최소 비용 경로를 찾는다
		dist := make([]int, n)
		for i := range dist {
			dist[i] = INF
		}
		dist[s] = 0

		inQueue := make([]bool, n)
		inQueue[s] = true

		prevEdge := make([]int, n) // 경로 역추적용: 해당 노드로 들어온 간선 인덱스
		for i := range prevEdge {
			prevEdge[i] = -1
		}

		queue := []int{s}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			inQueue[u] = false

			// u에서 나가는 모든 간선을 확인한다
			for _, idx := range graph[u] {
				e := &edges[idx]
				// 잔여 용량이 있고 비용이 더 저렴한 경로가 있으면 갱신
				if e.cap-e.flow > 0 && dist[u]+e.cost < dist[e.to] {
					dist[e.to] = dist[u] + e.cost
					prevEdge[e.to] = idx
					if !inQueue[e.to] {
						inQueue[e.to] = true
						queue = append(queue, e.to)
					}
				}
			}
		}

		// 싱크에 도달할 수 없으면 종료
		if dist[t] == INF {
			break
		}

		// 경로 상 최소 잔여 용량을 구한다
		f := INF
		for v := t; v != s; {
			idx := prevEdge[v]
			e := &edges[idx]
			if e.cap-e.flow < f {
				f = e.cap - e.flow
			}
			v = edges[idx^1].to // 역방향 간선의 to가 이전 노드
		}

		// 경로를 따라 유량을 보낸다
		for v := t; v != s; {
			idx := prevEdge[v]
			edges[idx].flow += f
			edges[idx^1].flow -= f
			v = edges[idx^1].to
		}

		totalFlow += f
		totalCost += f * dist[t]
	}

	return totalFlow, totalCost
}

func main() {
	// 예시: 2명의 작업자를 2개의 작업에 최소 비용으로 배정
	// 작업자 A(노드1), B(노드2) → 작업 X(노드3), Y(노드4)
	// 비용: A→X=3, A→Y=8, B→X=5, B→Y=2
	// 소스 S=0, 싱크 T=5

	n = 6
	edges = make([]Edge, 0)
	graph = make([][]int, n)

	S, T := 0, 5
	A, B, X, Y := 1, 2, 3, 4

	// 소스 → 작업자 (용량 1, 비용 0)
	addEdge(S, A, 1, 0)
	addEdge(S, B, 1, 0)

	// 작업자 → 작업 (용량 1, 비용 = 배정 비용)
	addEdge(A, X, 1, 3)
	addEdge(A, Y, 1, 8)
	addEdge(B, X, 1, 5)
	addEdge(B, Y, 1, 2)

	// 작업 → 싱크 (용량 1, 비용 0)
	addEdge(X, T, 1, 0)
	addEdge(Y, T, 1, 0)

	flow, cost := mcmf(S, T)

	fmt.Println("=== 최소 비용 최대 유량 (MCMF) ===")
	fmt.Printf("최대 유량: %d\n", flow)
	fmt.Printf("최소 비용: %d\n", cost)
	// 출력:
	// 최대 유량: 2
	// 최소 비용: 5
	// (A→X 비용3, B→Y 비용2, 합계 5)
}
