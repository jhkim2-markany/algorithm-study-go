package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const eps = 1e-9

// gaussianElimination은 확대 행렬 a[n][n+1]을 받아 유일해를 구한다.
func gaussianElimination(a [][]float64, n int) []float64 {
	// 전진 소거 (Forward Elimination)
	for col := 0; col < n; col++ {
		// 부분 피벗팅: 절댓값이 가장 큰 행을 피벗으로 선택
		pivotRow := col
		for row := col + 1; row < n; row++ {
			if math.Abs(a[row][col]) > math.Abs(a[pivotRow][col]) {
				pivotRow = row
			}
		}

		// 피벗 행 교환
		a[col], a[pivotRow] = a[pivotRow], a[col]

		// 피벗 아래 행 소거
		for row := col + 1; row < n; row++ {
			factor := a[row][col] / a[col][col]
			for j := col; j <= n; j++ {
				a[row][j] -= factor * a[col][j]
			}
		}
	}

	// 후진 대입 (Back Substitution)
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		x[i] = a[i][n]
		for j := i + 1; j < n; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}

	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 미지수의 수
	var n int
	fmt.Fscan(reader, &n)

	// 입력: 확대 행렬 [A|b]
	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	// 가우스 소거법으로 해 구하기
	x := gaussianElimination(a, n)

	// 출력: 해 벡터
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		if math.Abs(x[i]) < eps {
			fmt.Fprintf(writer, "%.6f", 0.0)
		} else {
			fmt.Fprintf(writer, "%.6f", x[i])
		}
	}
	fmt.Fprintln(writer)
}
