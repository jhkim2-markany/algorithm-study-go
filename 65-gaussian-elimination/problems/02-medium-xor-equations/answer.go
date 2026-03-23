package main

import (
	"bufio"
	"fmt"
	"os"
)

// solveXorEquations는 GF(2) 위의 XOR 연립방정식을 가우스 소거법으로 풀어
// 해를 구한다. 해가 없으면 nil을 반환한다.
//
// [매개변수]
//   - n: 미지수 수
//   - m: 방정식 수
//   - indices: 각 방정식에 포함된 변수 인덱스 목록 (1-indexed)
//   - b: 각 방정식의 결과값
//
// [반환값]
//   - []int: 해 벡터 (자유 변수는 0), 해가 없으면 nil
//
// [알고리즘 힌트]
//   1. 각 방정식을 비트 배열로 표현한다 (a[i][j] = 0 또는 1)
//   2. 전진 소거: 각 열에서 1인 행을 찾아 피벗으로 선택한다
//   3. 다른 모든 행에서 XOR 연산으로 해당 열을 소거한다
//   4. 소거 후 "0 0 ... 0 | 1" 행이 있으면 해가 없다
//   5. 후진 대입으로 해를 구하고, 자유 변수는 0으로 설정한다
func solveXorEquations(n, m int, indices [][]int, bVals []int) []int {
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n+1)
		for _, idx := range indices[i] {
			a[i][idx-1] = 1
		}
		a[i][n] = bVals[i]
	}

	pivotRow := 0
	pivotCol := make([]int, n)
	for i := range pivotCol {
		pivotCol[i] = -1
	}

	for col := 0; col < n && pivotRow < m; col++ {
		found := -1
		for row := pivotRow; row < m; row++ {
			if a[row][col] == 1 {
				found = row
				break
			}
		}
		if found == -1 {
			continue
		}
		a[pivotRow], a[found] = a[found], a[pivotRow]
		pivotCol[col] = pivotRow

		for row := 0; row < m; row++ {
			if row != pivotRow && a[row][col] == 1 {
				for j := col; j <= n; j++ {
					a[row][j] ^= a[pivotRow][j]
				}
			}
		}
		pivotRow++
	}

	for row := pivotRow; row < m; row++ {
		if a[row][n] == 1 {
			return nil
		}
	}

	x := make([]int, n)
	for col := 0; col < n; col++ {
		if pivotCol[col] != -1 {
			x[col] = a[pivotCol[col]][n]
		}
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	indices := make([][]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		var k int
		fmt.Fscan(reader, &k)
		indices[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &indices[i][j])
		}
		fmt.Fscan(reader, &b[i])
	}

	x := solveXorEquations(n, m, indices, b)
	if x == nil {
		fmt.Fprintln(writer, "IMPOSSIBLE")
	} else {
		for i := 0; i < n; i++ {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, x[i])
		}
		fmt.Fprintln(writer)
	}
}
