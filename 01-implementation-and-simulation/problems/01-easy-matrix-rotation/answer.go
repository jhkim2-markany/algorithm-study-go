package main

import (
	"bufio"
	"fmt"
	"os"
)

// rotateMatrix는 N×N 행렬을 시계 방향으로 90도 회전한 결과를 반환한다.
//
// [매개변수]
//   - matrix: N×N 크기의 2차원 정수 배열
//   - n: 행렬의 크기 (행과 열의 수)
//
// [반환값]
//   - [][]int: 시계 방향으로 90도 회전된 N×N 행렬
//
// [알고리즘 힌트]
//
//	시계 방향 90도 회전 시, 원래 행렬의 (i, j) 위치에 있는 값은
//	새 행렬의 (j, n-1-i) 위치로 이동한다.
//	즉, rotated[i][j] = matrix[n-1-j][i]
//
//	예시 (3×3):
//	  원래:       회전 후:
//	  1 2 3       7 4 1
//	  4 5 6  →    8 5 2
//	  7 8 9       9 6 3
func rotateMatrix(matrix [][]int, n int) [][]int {
	rotated := make([][]int, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = matrix[n-1-j][i]
		}
	}
	return rotated
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 행렬 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 행렬 입력
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	// 핵심 함수 호출
	rotated := rotateMatrix(matrix, n)

	// 결과 출력
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, rotated[i][j])
		}
		fmt.Fprintln(writer)
	}
}
