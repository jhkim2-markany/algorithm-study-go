package main

import (
	"bufio"
	"fmt"
	"os"
)

// prefixSum2D는 2차원 행렬과 쿼리 목록을 받아 각 부분 행렬의 합을 반환한다.
//
// [매개변수]
//   - matrix: N×M 크기의 2차원 정수 배열 (0-indexed)
//   - queries: 쿼리 배열, 각 쿼리는 [r1, c1, r2, c2] (1-indexed)
//
// [반환값]
//   - []int: 각 쿼리에 대한 부분 행렬 합 결과 배열
func prefixSum2D(matrix [][]int, queries [][4]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 행렬 크기 N×M과 쿼리 수 Q 입력
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)

	// 행렬 입력
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	// 쿼리 입력
	queries := make([][4]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2], &queries[i][3])
	}

	// 핵심 함수 호출
	results := prefixSum2D(matrix, queries)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
