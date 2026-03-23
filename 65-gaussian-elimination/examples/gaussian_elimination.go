package main

import (
	"fmt"
	"math"
)

// 가우스 소거법 (Gaussian Elimination) - 연립일차방정식을 행렬 행 연산으로 풀기
// 시간 복잡도: O(N³)
// 공간 복잡도: O(N²)

const eps = 1e-9

// gaussianElimination은 확대 행렬 a[N][N+1]을 받아 연립방정식의 해를 구한다.
// 반환값: 해 벡터, 해의 존재 여부 (0: 유일해, 1: 무한해, 2: 해 없음)
func gaussianElimination(a [][]float64, n int) ([]float64, int) {
	// 전진 소거 (Forward Elimination)
	for col := 0; col < n; col++ {
		// 부분 피벗팅: 현재 열에서 절댓값이 가장 큰 행을 찾는다
		pivotRow := col
		for row := col + 1; row < n; row++ {
			if math.Abs(a[row][col]) > math.Abs(a[pivotRow][col]) {
				pivotRow = row
			}
		}

		// 피벗 행과 현재 행을 교환한다
		a[col], a[pivotRow] = a[pivotRow], a[col]

		// 피벗이 0이면 유일해가 존재하지 않는다
		if math.Abs(a[col][col]) < eps {
			continue
		}

		// 피벗 아래의 모든 행을 소거한다
		for row := col + 1; row < n; row++ {
			factor := a[row][col] / a[col][col]
			for j := col; j <= n; j++ {
				a[row][j] -= factor * a[col][j]
			}
		}
	}

	// 해의 존재 여부 판별
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		if math.Abs(a[i][i]) < eps {
			if math.Abs(a[i][n]) > eps {
				return nil, 2 // 해 없음 (0 = 상수 ≠ 0)
			}
			return nil, 1 // 무한해 (자유 변수 존재)
		}

		// 후진 대입 (Back Substitution)
		x[i] = a[i][n]
		for j := i + 1; j < n; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}

	return x, 0 // 유일해
}

func main() {
	// 예제: 연립방정식
	//  2x +  y -  z =  8
	// -3x -  y + 2z = -11
	// -2x +  y + 2z = -3
	// 기대 해: x=2, y=3, z=-1

	var n int
	fmt.Scan(&n)

	// 확대 행렬 [A|b] 입력
	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	// 가우스 소거법 수행
	x, status := gaussianElimination(a, n)

	switch status {
	case 0:
		// 유일해 출력
		for i := 0; i < n; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			// 매우 작은 값은 0으로 처리
			if math.Abs(x[i]) < eps {
				fmt.Printf("%.6f", 0.0)
			} else {
				fmt.Printf("%.6f", x[i])
			}
		}
		fmt.Println()
	case 1:
		fmt.Println("MULTIPLE")
	case 2:
		fmt.Println("INCONSISTENT")
	}
}
