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
func maxRegion(grid [][]int) int {
	// 여기에 코드를 작성하세요
	return 0
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
