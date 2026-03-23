package main

import (
	"bufio"
	"fmt"
	"os"
)

// XOR 가우스 소거법 (GF(2) Gaussian Elimination)
// 각 방정식을 비트마스크로 표현하여 효율적으로 소거한다.

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 미지수 수 N, 방정식 수 M
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 각 방정식을 비트 배열로 저장 (a[i][j] = 0 또는 1, a[i][n] = 결과값 b)
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n+1)
		var k int
		fmt.Fscan(reader, &k)
		for j := 0; j < k; j++ {
			var idx int
			fmt.Fscan(reader, &idx)
			a[i][idx-1] = 1 // 1-indexed → 0-indexed
		}
		fmt.Fscan(reader, &a[i][n]) // 결과값 b
	}

	// XOR 가우스 소거법 (전진 소거)
	pivotRow := 0
	pivotCol := make([]int, n) // 각 열의 피벗 행 위치 (-1이면 자유 변수)
	for i := range pivotCol {
		pivotCol[i] = -1
	}

	for col := 0; col < n && pivotRow < m; col++ {
		// 현재 열에서 1인 행을 찾는다
		found := -1
		for row := pivotRow; row < m; row++ {
			if a[row][col] == 1 {
				found = row
				break
			}
		}
		if found == -1 {
			continue // 이 열에는 피벗이 없다 (자유 변수)
		}

		// 피벗 행과 현재 행을 교환
		a[pivotRow], a[found] = a[found], a[pivotRow]
		pivotCol[col] = pivotRow

		// 다른 모든 행에서 현재 열을 소거 (XOR 연산)
		for row := 0; row < m; row++ {
			if row != pivotRow && a[row][col] == 1 {
				for j := col; j <= n; j++ {
					a[row][j] ^= a[pivotRow][j]
				}
			}
		}
		pivotRow++
	}

	// 해의 존재 여부 판별
	// 소거 후 "0 0 ... 0 | 1" 형태의 행이 있으면 해가 없다
	for row := pivotRow; row < m; row++ {
		if a[row][n] == 1 {
			fmt.Fprintln(writer, "IMPOSSIBLE")
			return
		}
	}

	// 후진 대입으로 해 구하기 (자유 변수는 0으로 설정)
	x := make([]int, n)
	for col := 0; col < n; col++ {
		if pivotCol[col] != -1 {
			x[col] = a[pivotCol[col]][n]
		}
		// 자유 변수는 기본값 0
	}

	// 출력
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x[i])
	}
	fmt.Fprintln(writer)
}
