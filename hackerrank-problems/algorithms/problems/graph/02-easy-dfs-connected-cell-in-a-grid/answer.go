package main

import (
	"bufio"
	"fmt"
	"os"
)

// maxRegion은 그리드에서 가장 큰 연결 영역의 크기를 반환한다.
//
// [매개변수]
//   - grid: N×M 이진 행렬
//
// [반환값]
//   - int: 가장 큰 연결 영역의 셀 수
//
// [알고리즘 힌트]
//
//	각 셀에서 DFS를 수행하여 8방향으로 연결된 1의 개수를 센다.
//	방문 체크를 위해 원본 그리드의 값을 0으로 변경한다.
func maxRegion(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	// 8방향 이동 벡터
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// DFS로 연결 영역 크기 계산
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		// 범위 밖이거나 0이면 종료
		if r < 0 || r >= n || c < 0 || c >= m || grid[r][c] == 0 {
			return 0
		}
		// 방문 처리
		grid[r][c] = 0
		count := 1
		// 8방향 탐색
		for d := 0; d < 8; d++ {
			count += dfs(r+dx[d], c+dy[d])
		}
		return count
	}

	// 모든 셀을 순회하며 최대 영역 크기 탐색
	best := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				size := dfs(i, j)
				if size > best {
					best = size
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
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	result := maxRegion(grid)
	fmt.Fprintln(writer, result)
}
