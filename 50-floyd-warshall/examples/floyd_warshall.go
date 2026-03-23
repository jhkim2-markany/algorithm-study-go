package main

import "fmt"

// 플로이드-워셜 알고리즘 - 모든 쌍 최단 경로
// 시간 복잡도: O(V³)
// 공간 복잡도: O(V²)

const INF = 1<<31 - 1 // 무한대를 나타내는 값

// floydWarshall은 인접 행렬로 표현된 그래프에서 모든 쌍 최단 거리를 구한다.
// dist[i][j]는 정점 i에서 j까지의 간선 가중치이며, 연결이 없으면 INF이다.
// 함수 수행 후 dist[i][j]에 최단 거리가 저장된다.
func floydWarshall(dist [][]int, n int) {
	// 중간 정점 k를 하나씩 추가하며 최단 거리를 갱신한다
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// 오버플로우 방지: i→k 또는 k→j가 도달 불가능하면 건너뛴다
				if dist[i][k] == INF || dist[k][j] == INF {
					continue
				}
				// k를 경유하는 경로가 더 짧으면 갱신한다
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
}

// floydWarshallWithPath는 최단 거리와 함께 경로 복원용 next 행렬을 구한다.
// next[i][j]는 i에서 j로 가는 최단 경로에서 i 다음에 방문할 정점이다.
func floydWarshallWithPath(dist [][]int, next [][]int, n int) {
	// 중간 정점 k를 순회하며 최단 거리와 경로를 갱신한다
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] == INF || dist[k][j] == INF {
					continue
				}
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					// 경로 복원: i에서 j로 갈 때 k를 경유하므로 next[i][j]를 갱신한다
					next[i][j] = next[i][k]
				}
			}
		}
	}
}

// reconstructPath는 next 행렬을 이용하여 정점 from에서 to까지의 경로를 복원한다.
func reconstructPath(next [][]int, from, to int) []int {
	// 경로가 존재하지 않는 경우
	if next[from][to] == -1 {
		return nil
	}
	// 시작 정점부터 도착 정점까지 next를 따라간다
	path := []int{from}
	cur := from
	for cur != to {
		cur = next[cur][to]
		path = append(path, cur)
	}
	return path
}

func main() {
	// 4개 정점, 5개 간선의 방향 그래프 예제
	n := 4

	// 거리 행렬 초기화
	dist := make([][]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
				next[i][j] = j
			} else {
				dist[i][j] = INF
				next[i][j] = -1
			}
		}
	}

	// 간선 추가 (방향 그래프)
	edges := [][3]int{
		{0, 1, 3},  // 0 → 1, 가중치 3
		{0, 3, 7},  // 0 → 3, 가중치 7
		{1, 2, 2},  // 1 → 2, 가중치 2
		{2, 3, 1},  // 2 → 3, 가중치 1
		{3, 1, -2}, // 3 → 1, 가중치 -2 (음수 가중치)
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		dist[u][v] = w
		next[u][v] = v
	}

	// 플로이드-워셜 실행 (경로 복원 포함)
	floydWarshallWithPath(dist, next, n)

	// 결과 출력: 모든 쌍 최단 거리
	fmt.Println("=== 모든 쌍 최단 거리 ===")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dist[i][j] == INF {
				fmt.Printf("  INF")
			} else {
				fmt.Printf("%5d", dist[i][j])
			}
		}
		fmt.Println()
	}

	// 경로 복원 예시: 0에서 3까지의 최단 경로
	fmt.Println("\n=== 경로 복원: 0 → 3 ===")
	path := reconstructPath(next, 0, 3)
	if path != nil {
		fmt.Printf("경로: ")
		for i, v := range path {
			if i > 0 {
				fmt.Printf(" → ")
			}
			fmt.Printf("%d", v)
		}
		fmt.Printf(" (거리: %d)\n", dist[0][3])
	} else {
		fmt.Println("경로 없음")
	}
}
