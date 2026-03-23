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
func solveEquations(a [][]float64, n int) []float64 {
	// 여기에 코드를 작성하세요
	_ = math.Abs
	return nil
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
