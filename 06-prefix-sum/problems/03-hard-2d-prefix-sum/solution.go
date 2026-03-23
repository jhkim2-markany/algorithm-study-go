package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	// 2차원 누적합 배열 구축 (1-indexed)
	// prefix[i][j] = (1,1)부터 (i,j)까지의 부분 행렬 합
	prefix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prefix[i] = make([]int, m+1)
	}

	// 포함-배제 원리로 누적합 계산
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	// 각 쿼리에 대해 부분 행렬 합 계산
	for i := 0; i < q; i++ {
		var r1, c1, r2, c2 int
		fmt.Fscan(reader, &r1, &c1, &r2, &c2)

		// 포함-배제 원리로 부분 행렬 합 계산
		// sum = prefix[r2][c2] - prefix[r1-1][c2] - prefix[r2][c1-1] + prefix[r1-1][c1-1]
		sum := prefix[r2][c2] - prefix[r1-1][c2] - prefix[r2][c1-1] + prefix[r1-1][c1-1]
		fmt.Fprintln(writer, sum)
	}
}
