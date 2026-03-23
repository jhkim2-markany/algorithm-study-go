package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// solveEquations는 가우스 소거법으로 N원 연립일차방정식의 유일해를 구한다.
//
// [매개변수]
//   - a: N×(N+1) 확대 행렬 [A|b]
//   - n: 미지수의 수
//
// [반환값]
//   - []float64: 해 벡터 x
//
// [알고리즘 힌트]
//   1. 전진 소거: 각 열에서 절댓값이 가장 큰 행을 피벗으로 선택(부분 피벗팅)
//   2. 피벗 행을 교환하고, 피벗 아래 행을 소거한다
//   3. 후진 대입: 마지막 행부터 역순으로 해를 구한다
//   4. x[i] = (a[i][n] - Σ(a[i][j]*x[j])) / a[i][i]
func solveEquations(a [][]float64, n int) []float64 {
	for col := 0; col < n; col++ {
		pivotRow := col
		for row := col + 1; row < n; row++ {
			if math.Abs(a[row][col]) > math.Abs(a[pivotRow][col]) {
				pivotRow = row
			}
		}
		a[col], a[pivotRow] = a[pivotRow], a[col]

		for row := col + 1; row < n; row++ {
			factor := a[row][col] / a[col][col]
			for j := col; j <= n; j++ {
				a[row][j] -= factor * a[col][j]
			}
		}
	}

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

	var n int
	fmt.Fscan(reader, &n)

	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	x := solveEquations(a, n)

	const eps = 1e-9
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
