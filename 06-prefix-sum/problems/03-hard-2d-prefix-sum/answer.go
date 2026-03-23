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
//
// [알고리즘 힌트]
//
//	2차원 누적합 배열을 구축한다 (1-indexed).
//	prefix[i][j] = (1,1)부터 (i,j)까지의 부분 행렬 합
//	구축: prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i-1][j-1]
//	쿼리: sum(r1,c1,r2,c2) = prefix[r2][c2] - prefix[r1-1][c2] - prefix[r2][c1-1] + prefix[r1-1][c1-1]
//	(포함-배제 원리)
//
//	시간복잡도: 전처리 O(N×M), 쿼리당 O(1)
func prefixSum2D(matrix [][]int, queries [][4]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	prefix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prefix[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	results := make([]int, len(queries))
	for i, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		results[i] = prefix[r2][c2] - prefix[r1-1][c2] - prefix[r2][c1-1] + prefix[r1-1][c1-1]
	}
	return results
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
