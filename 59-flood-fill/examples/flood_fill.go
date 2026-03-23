package main

import "fmt"

// 플러드 필 (Flood Fill) - 격자 영역 탐색과 채우기
// BFS 기반으로 연결된 영역을 탐색한다.
// 시간 복잡도: O(N × M)
// 공간 복잡도: O(N × M)

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

// floodFillBFS는 (sr, sc)에서 시작하여 같은 값을 가진 연결 영역을 newColor로 채운다
func floodFillBFS(grid [][]int, sr, sc, newColor int) {
	n := len(grid)
	m := len(grid[0])
	oldColor := grid[sr][sc]

	// 이미 같은 색이면 아무것도 하지 않는다
	if oldColor == newColor {
		return
	}

	// BFS 큐 초기화
	queue := [][2]int{{sr, sc}}
	grid[sr][sc] = newColor // 시작 칸 색 변경

	for len(queue) > 0 {
		// 큐에서 꺼내기
		cur := queue[0]
		queue = queue[1:]
		r, c := cur[0], cur[1]

		// 4방향 인접 칸 탐색
		for d := 0; d < 4; d++ {
			nr, nc := r+dr[d], c+dc[d]
			// 범위 체크 및 같은 색인지 확인
			if nr >= 0 && nr < n && nc >= 0 && nc < m && grid[nr][nc] == oldColor {
				grid[nr][nc] = newColor // 색 변경 (방문 처리)
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
}

// countRegions는 격자에서 값이 target인 연결 영역의 개수를 센다
func countRegions(grid [][]int, target int) int {
	n := len(grid)
	m := len(grid[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	count := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == target && !visited[r][c] {
				// 새로운 영역 발견 → BFS로 영역 전체 방문
				bfs(grid, visited, r, c, target)
				count++
			}
		}
	}
	return count
}

// bfs는 (sr, sc)에서 시작하여 target 값을 가진 연결 영역을 방문 처리한다
func bfs(grid [][]int, visited [][]bool, sr, sc, target int) {
	n := len(grid)
	m := len(grid[0])
	queue := [][2]int{{sr, sc}}
	visited[sr][sc] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		r, c := cur[0], cur[1]

		for d := 0; d < 4; d++ {
			nr, nc := r+dr[d], c+dc[d]
			if nr >= 0 && nr < n && nc >= 0 && nc < m &&
				grid[nr][nc] == target && !visited[nr][nc] {
				visited[nr][nc] = true
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
}

func main() {
	// 예시 1: 색 채우기 (Flood Fill)
	fmt.Println("=== 플러드 필: 색 채우기 ===")
	grid1 := [][]int{
		{1, 1, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{1, 1, 0, 0, 0},
	}

	fmt.Println("변경 전:")
	for _, row := range grid1 {
		fmt.Println(row)
	}

	// (0,0)에서 시작하여 색 1을 색 3으로 변경
	floodFillBFS(grid1, 0, 0, 3)

	fmt.Println("(0,0)에서 색 1→3으로 변경 후:")
	for _, row := range grid1 {
		fmt.Println(row)
	}
	// 출력:
	// [3 3 0 0 1]
	// [3 3 0 1 1]
	// [0 0 1 0 0]
	// [1 1 0 0 0]

	// 예시 2: 영역 개수 세기
	fmt.Println("\n=== 플러드 필: 영역 개수 세기 ===")
	grid2 := [][]int{
		{1, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1},
	}

	fmt.Println("격자:")
	for _, row := range grid2 {
		fmt.Println(row)
	}

	regions := countRegions(grid2, 1)
	fmt.Printf("값 1로 이루어진 영역 수: %d\n", regions)
	// 출력: 값 1로 이루어진 영역 수: 5
}
