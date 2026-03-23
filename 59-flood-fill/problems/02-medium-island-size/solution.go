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

	// 1단계: 각 섬에 고유 번호를 부여하고 크기를 구한다 (플러드 필)
	islandID := make([][]int, n) // 각 칸이 속한 섬 번호 (0이면 바다)
	for i := 0; i < n; i++ {
		islandID[i] = make([]int, m)
	}

	islandSize := map[int]int{} // 섬 번호 → 크기
	id := 2                     // 섬 번호는 2부터 시작 (0=바다, 1=원래 땅)

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 && islandID[r][c] == 0 {
				// 새로운 섬 발견 → BFS로 번호 부여
				size := bfs(grid, islandID, r, c, n, m, id)
				islandSize[id] = size
				id++
			}
		}
	}

	// 2단계: 각 바다 칸(0)에 대해 인접한 섬들을 합친 크기를 계산
	best := 0

	// 바다가 없는 경우를 대비해 기존 최대 섬 크기도 확인
	for _, sz := range islandSize {
		if sz > best {
			best = sz
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 0 {
				// 인접한 섬 번호를 중복 없이 수집
				seen := map[int]bool{}
				total := 1 // 현재 바다 칸을 땅으로 바꾸므로 +1

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

	// 출력: 가장 큰 섬의 크기
	fmt.Fprintln(writer, best)
}

// bfs는 (sr, sc)에서 시작하여 연결된 땅에 섬 번호를 부여하고 크기를 반환한다
func bfs(grid [][]int, islandID [][]int, sr, sc, n, m, id int) int {
	queue := [][2]int{{sr, sc}}
	islandID[sr][sc] = id
	size := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		r, c := cur[0], cur[1]
		size++

		// 4방향 인접 칸 탐색
		for d := 0; d < 4; d++ {
			nr, nc := r+dr[d], c+dc[d]
			if nr >= 0 && nr < n && nc >= 0 && nc < m &&
				grid[nr][nc] == 1 && islandID[nr][nc] == 0 {
				islandID[nr][nc] = id // 섬 번호 부여 (방문 처리)
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
	return size
}
