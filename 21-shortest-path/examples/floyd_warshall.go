package main

import (
	"fmt"
	"math"
)

// Floyd-Warshall 알고리즘 - 모든 정점 쌍 사이의 최단 경로
// 시간 복잡도: O(V³)
// 공간 복잡도: O(V²)

const INF = math.MaxInt64 / 2 // 오버플로 방지를 위해 절반값 사용

// floydWarshall 함수는 모든 정점 쌍의 최단 거리 행렬을 반환한다
func floydWarshall(n int, dist [][]int) [][]int {
	// 중간 정점 k를 하나씩 추가하며 최단 거리를 갱신한다
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// k를 경유하는 경로가 더 짧으면 갱신
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist
}

func main() {
	// 그래프 생성 (정점 4개, 0~3)
	n := 4
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}

	// 간선 추가 (방향 그래프)
	dist[0][1] = 3
	dist[0][3] = 7
	dist[1][0] = 8
	dist[1][2] = 2
	dist[2][0] = 5
	dist[2][3] = 1
	dist[3][0] = 2

	// Floyd-Warshall 실행
	dist = floydWarshall(n, dist)

	// 결과 출력
	fmt.Println("=== Floyd-Warshall 모든 쌍 최단 경로 ===")
	fmt.Print("    ")
	for j := 0; j < n; j++ {
		fmt.Printf("%4d", j)
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Printf("%d : ", i)
		for j := 0; j < n; j++ {
			if dist[i][j] >= INF {
				fmt.Print(" INF")
			} else {
				fmt.Printf("%4d", dist[i][j])
			}
		}
		fmt.Println()
	}
}
