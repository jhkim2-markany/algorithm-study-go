package main

import (
	"bufio"
	"fmt"
	"os"
)

// connectedCell은 이진 행렬에서 가장 큰 영역의 셀 수를 반환한다.
//
// [매개변수]
//   - matrix: N×M 이진 행렬
//
// [반환값]
//   - int: 가장 큰 영역의 셀 수
func connectedCell(matrix [][]int) int {
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

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	fmt.Fprintln(writer, connectedCell(matrix))
}
