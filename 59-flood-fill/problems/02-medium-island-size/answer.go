package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

// maxIslandSize는 0과 1로 이루어진 격자에서 바다(0) 한 칸을
// 땅(1)으로 바꿨을 때 만들 수 있는 가장 큰 섬의 크기를 구한다.
//
// [매개변수]
//   - grid: n×m 크기의 격자 (0: 바다, 1: 땅)
//   - n: 격자의 행 수
//   - m: 격자의 열 수
//
// [반환값]
//   - int: 바다 한 칸을 땅으로 바꿨을 때 가능한 최대 섬 크기
//
// [알고리즘 힌트]
//   1. BFS로 각 섬에 고유 번호(2부터)를 부여하고 크기를 기록한다
//   2. 각 바다 칸에 대해 상하좌우 인접한 섬 번호를 중복 없이 수집한다
//   3. 인접 섬들의 크기 합 + 1(바다→땅)이 해당 칸을 바꿨을 때의 섬 크기이다
//   4. 모든 바다 칸 중 최대값을 구하되, 바다가 없는 경우도 처리한다
func maxIslandSize(grid [][]int, n, m int) int {
	islandID := make([][]int, n)
	for i := 0; i < n; i++ {
		islandID[i] = make([]int, m)
	}

	islandSize := map[int]int{}
	id := 2

	// BFS로 섬 번호 부여
	bfs := func(sr, sc, id int) int {
		queue := [][2]int{{sr, sc}}
		islandID[sr][sc] = id
		size := 0
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			r, c := cur[0], cur[1]
			size++
			for d := 0; d < 4; d++ {
				nr, nc := r+dr[d], c+dc[d]
				if nr >= 0 && nr < n && nc >= 0 && nc < m &&
					grid[nr][nc] == 1 && islandID[nr][nc] == 0 {
					islandID[nr][nc] = id
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		return size
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 && islandID[r][c] == 0 {
				size := bfs(r, c, id)
				islandSize[id] = size
				id++
			}
		}
	}

	best := 0
	for _, sz := range islandSize {
		if sz > best {
			best = sz
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 0 {
				seen := map[int]bool{}
				total := 1
				for d := 0; d < 4; d++ {
					nr, nc := r+dr[d], c+dc[d]
					if nr >= 0 && nr < n && nc >= 0 && nc < m && islandID[nr][nc] > 0 {
						sid := islandID[nr][nc]
						if !seen[sid] {
							seen[sid] = true
							total += islandSize[sid]
						}
					}
				}
				if total > best {
					best = total
				}
			}
		}
	}
	return best
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

	fmt.Fprintln(writer, maxIslandSize(grid, n, m))
}
