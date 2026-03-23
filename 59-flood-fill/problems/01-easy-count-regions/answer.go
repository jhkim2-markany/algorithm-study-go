package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

// countRegions는 0과 1로 이루어진 격자에서 1로 연결된 영역(섬)의
// 개수를 BFS를 이용하여 구한다. 상하좌우로 인접한 1끼리 같은 영역이다.
//
// [매개변수]
//   - grid: n×m 크기의 격자 (0: 바다, 1: 땅)
//   - n: 격자의 행 수
//   - m: 격자의 열 수
//
// [반환값]
//   - int: 1로 이루어진 연결 영역의 개수
//
// [알고리즘 힌트]
//   1. 방문 배열을 만들어 각 칸의 방문 여부를 관리한다
//   2. 모든 칸을 순회하며 미방문 땅(1)을 발견하면 BFS로 연결된 영역 전체를 방문 처리한다
//   3. BFS를 시작할 때마다 영역 카운트를 1 증가시킨다
//   4. BFS에서는 큐에 넣을 때 방문 처리하여 중복 방문을 방지한다
func countRegions(grid [][]int, n, m int) int {
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	count := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 && !visited[r][c] {
				// BFS로 연결된 영역 전체 방문
				queue := [][2]int{{r, c}}
				visited[r][c] = true
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					cr, cc := cur[0], cur[1]
					for d := 0; d < 4; d++ {
						nr, nc := cr+dr[d], cc+dc[d]
						if nr >= 0 && nr < n && nc >= 0 && nc < m &&
							grid[nr][nc] == 1 && !visited[nr][nc] {
							visited[nr][nc] = true
							queue = append(queue, [2]int{nr, nc})
						}
					}
				}
				count++
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	fmt.Fprintln(writer, countRegions(grid, n, m))
}
