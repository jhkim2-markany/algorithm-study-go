package main

import (
	"bufio"
	"fmt"
	"os"
)

// 상하좌우 이동을 위한 방향 배열
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 미로 크기 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 미로 입력
	maze := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &maze[i])
	}

	// BFS를 위한 거리 배열 초기화 (-1은 미방문)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	// 시작점 (0, 0)에서 BFS 시작
	type point struct{ x, y int }
	queue := []point{{0, 0}}
	dist[0][0] = 1 // 시작 칸도 포함하므로 거리 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// 4방향 탐색
		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx[d], cur.y+dy[d]

			// 범위 확인, 벽 확인, 방문 확인
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if maze[nx][ny] == '0' || dist[nx][ny] != -1 {
				continue
			}

			// 이동 가능한 칸이면 거리 갱신 후 큐에 추가
			dist[nx][ny] = dist[cur.x][cur.y] + 1
			queue = append(queue, point{nx, ny})
		}
	}

	// 도착점까지의 최단 거리 출력
	fmt.Fprintln(writer, dist[n-1][m-1])
}
