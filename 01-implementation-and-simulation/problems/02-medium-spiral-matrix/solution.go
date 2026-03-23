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

	// 행렬 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// N×N 행렬 초기화
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	// 방향 배열: 오른쪽, 아래, 왼쪽, 위 순서
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}

	// 나선형 순서로 1부터 N²까지 채우기
	x, y := 0, 0 // 현재 위치
	dir := 0     // 현재 방향 (0: 오른쪽)
	for num := 1; num <= n*n; num++ {
		matrix[x][y] = num

		// 다음 위치 계산
		nx, ny := x+dx[dir], y+dy[dir]

		// 범위를 벗어나거나 이미 채워진 칸이면 방향 전환
		if nx < 0 || nx >= n || ny < 0 || ny >= n || matrix[nx][ny] != 0 {
			dir = (dir + 1) % 4
			nx, ny = x+dx[dir], y+dy[dir]
		}

		x, y = nx, ny
	}

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
