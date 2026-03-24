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
//
// [알고리즘 힌트]
//
//	각 1인 셀에서 DFS를 수행하여 8방향으로 연결된 영역 크기를 구한다.
//	방문한 셀은 0으로 바꿔 재방문을 방지한다.
func connectedCell(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])

	// 8방향 이동 벡터
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// DFS로 영역 크기를 계산하는 함수
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		// 범위 밖이거나 0이면 종료
		if r < 0 || r >= n || c < 0 || c >= m || matrix[r][c] == 0 {
			return 0
		}
		// 방문 처리
		matrix[r][c] = 0
		size := 1

		// 8방향 탐색
		for d := 0; d < 8; d++ {
			size += dfs(r+dx[d], c+dy[d])
		}
		return size
	}

	// 모든 셀을 순회하며 최대 영역 크기 탐색
	maxSize := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				size := dfs(i, j)
				if size > maxSize {
					maxSize = size
				}
			}
		}
	}

	return maxSize
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
