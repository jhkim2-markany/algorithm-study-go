package main

import (
	"bufio"
	"fmt"
	"os"
)

// minimumMoves는 격자에서 시작 위치부터 목표 위치까지의 최소 이동 횟수를 반환한다.
//
// [매개변수]
//   - grid: N × N 격자 (각 행은 문자열)
//   - startX: 시작 행
//   - startY: 시작 열
//   - goalX: 목표 행
//   - goalY: 목표 열
//
// [반환값]
//   - int: 최소 이동 횟수
//
// [알고리즘 힌트]
//
//	BFS를 사용하여 최소 이동 횟수를 구한다.
//	각 위치에서 상하좌우로 벽이나 경계까지 직선 이동하며
//	방문하지 않은 칸을 큐에 추가한다.
func minimumMoves(grid []string, startX, startY, goalX, goalY int) int {
	n := len(grid)

	// 방향 벡터: 상, 하, 좌, 우
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	// 방문 배열 및 거리 배열 초기화
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	// BFS 큐 초기화
	type point struct{ x, y int }
	queue := []point{{startX, startY}}
	dist[startX][startY] = 0

	// BFS 탐색
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// 목표 도달 확인
		if cur.x == goalX && cur.y == goalY {
			return dist[cur.x][cur.y]
		}

		// 네 방향으로 직선 이동
		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx[d], cur.y+dy[d]

			// 벽이나 경계까지 직선으로 이동
			for nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == '.' {
				// 방문하지 않은 칸이면 큐에 추가
				if dist[nx][ny] == -1 {
					dist[nx][ny] = dist[cur.x][cur.y] + 1
					queue = append(queue, point{nx, ny})
				}
				// 같은 방향으로 계속 이동
				nx += dx[d]
				ny += dy[d]
			}
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 격자 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 격자 입력
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	// 시작/목표 위치 입력
	var startX, startY, goalX, goalY int
	fmt.Fscan(reader, &startX, &startY, &goalX, &goalY)

	// 핵심 함수 호출
	result := minimumMoves(grid, startX, startY, goalX, goalY)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
