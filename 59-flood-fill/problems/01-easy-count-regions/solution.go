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
func countRegions(grid [][]int, n, m int) int {
	// 여기에 코드를 작성하세요
	return 0
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
