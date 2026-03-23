package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 격자 크기
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 입력: 격자 정보
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	// 방문 배열
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	// 모든 칸을 순회하며 미방문 땅(1)을 발견하면 BFS로 영역 탐색
	count := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 && !visited[r][c] {
				// 새로운 섬 발견 → BFS로 연결된 영역 전체 방문
				bfs(grid, visited, r, c, n, m)
				count++
			}
		}
	}

	// 출력: 섬의 개수
	fmt.Fprintln(writer, count)
}

// bfs는 (sr, sc)에서 시작하여 연결된 땅(1)을 모두 방문 처리한다
func bfs(grid [][]int, visited [][]bool, sr, sc, n, m int) {
	queue := [][2]int{{sr, sc}}
	visited[sr][sc] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		r, c := cur[0], cur[1]

		// 4방향 인접 칸 탐색
		for d := 0; d < 4; d++ {
			nr, nc := r+dr[d], c+dc[d]
			if nr >= 0 && nr < n && nc >= 0 && nc < m &&
				grid[nr][nc] == 1 && !visited[nr][nc] {
				visited[nr][nc] = true // 큐에 넣을 때 방문 처리
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
}
