package main

import (
	"bufio"
	"fmt"
	"os"
)

// buildSpiralMatrix는 N×N 크기의 행렬을 나선형(spiral) 순서로
// 1부터 N²까지 채워 반환한다.
//
// [매개변수]
//   - n: 행렬의 크기 (N×N)
//
// [반환값]
//   - [][]int: 나선형 순서로 채워진 N×N 행렬
//
// [알고리즘 힌트]
//
//	방향 배열을 사용하여 오른쪽→아래→왼쪽→위 순서로 이동하며 값을 채운다.
//	현재 위치에서 다음 위치가 범위를 벗어나거나 이미 채워진 칸이면
//	방향을 전환(시계 방향 90도)한다.
//
//	예시 (3×3):
//	  1 2 3
//	  8 9 4
//	  7 6 5
func buildSpiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}

	x, y := 0, 0
	dir := 0
	for num := 1; num <= n*n; num++ {
		matrix[x][y] = num
		nx, ny := x+dx[dir], y+dy[dir]
		if nx < 0 || nx >= n || ny < 0 || ny >= n || matrix[nx][ny] != 0 {
			dir = (dir + 1) % 4
			nx, ny = x+dx[dir], y+dy[dir]
		}
		x, y = nx, ny
	}

	return matrix
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 행렬 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	matrix := buildSpiralMatrix(n)

	// 결과 출력
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, matrix[i][j])
		}
		fmt.Fprintln(writer)
	}
}
