package main

import (
	"bufio"
	"fmt"
	"os"
)

// mazeShortestPath는 미로에서 (0,0)부터 (n-1,m-1)까지의 최단 거리를 반환한다.
//
// [매개변수]
//   - maze: N×M 크기의 미로 ('1'은 이동 가능, '0'은 벽)
//   - n: 행 수
//   - m: 열 수
//
// [반환값]
//   - int: 시작점에서 도착점까지의 최단 거리 (시작 칸 포함)
//
// [알고리즘 힌트]
//
//	BFS를 사용하여 시작점 (0,0)에서 도착점 (n-1,m-1)까지 탐색한다.
//	상하좌우 4방향으로 이동하며, 범위/벽/방문 여부를 확인한다.
//	거리 배열을 -1로 초기화하고, 시작점은 1로 설정한다 (시작 칸 포함).
//	BFS는 가중치 없는 그래프에서 최단 거리를 보장한다.
func mazeShortestPath(maze []string, n, m int) int {
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	type point struct{ x, y int }
	queue := []point{{0, 0}}
	dist[0][0] = 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx[d], cur.y+dy[d]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if maze[nx][ny] == '0' || dist[nx][ny] != -1 {
				continue
			}
			dist[nx][ny] = dist[cur.x][cur.y] + 1
			queue = append(queue, point{nx, ny})
		}
	}

	return dist[n-1][m-1]
}

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

	// 핵심 함수 호출
	result := mazeShortestPath(maze, n, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
