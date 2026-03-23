package main

import "fmt"

// 0-1 BFS - 가중치 0/1 그래프에서 덱 기반 최단 경로
// 간선 가중치가 0 또는 1인 그래프에서 최단 거리를 구한다.
// 시간 복잡도: O(V + E)
// 공간 복잡도: O(V + E)

const INF = 1<<31 - 1

// 간선 구조체: 도착 노드와 가중치(0 또는 1)
type Edge struct {
	to, weight int
}

// zeroOneBFS는 시작 노드에서 모든 노드까지의 최단 거리를 구한다
func zeroOneBFS(adj [][]Edge, start, n int) []int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = INF // 모든 거리를 무한대로 초기화
	}
	dist[start] = 0

	// 덱을 슬라이스로 구현
	deque := []int{start}

	for len(deque) > 0 {
		// 덱 앞에서 꺼내기
		v := deque[0]
		deque = deque[1:]

		for _, e := range adj[v] {
			newDist := dist[v] + e.weight
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				if e.weight == 0 {
					// 가중치 0: 덱 앞에 추가
					deque = append([]int{e.to}, deque...)
				} else {
					// 가중치 1: 덱 뒤에 추가
					deque = append(deque, e.to)
				}
			}
		}
	}
	return dist
}

func main() {
	// 예시 그래프:
	//   1 --0-- 2 --1-- 3
	//   |               |
	//   1               0
	//   |               |
	//   4 --0-- 5 --1-- 6
	//
	// 간선: (1,2,0), (2,3,1), (1,4,1), (3,6,0), (4,5,0), (5,6,1)

	n := 6
	adj := make([][]Edge, n+1)

	// 간선 추가 (양방향)
	edges := [][3]int{
		{1, 2, 0}, {2, 3, 1}, {1, 4, 1},
		{3, 6, 0}, {4, 5, 0}, {5, 6, 1},
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], Edge{e[0], e[2]})
	}

	// 노드 1에서 출발하는 최단 거리 계산
	dist := zeroOneBFS(adj, 1, n)

	fmt.Println("0-1 BFS 최단 거리 (시작: 노드 1):")
	for i := 1; i <= n; i++ {
		fmt.Printf("  노드 %d: %d\n", i, dist[i])
	}
	// 출력:
	// 0-1 BFS 최단 거리 (시작: 노드 1):
	//   노드 1: 0
	//   노드 2: 0
	//   노드 3: 1
	//   노드 4: 1
	//   노드 5: 1
	//   노드 6: 1
}
