package main

import (
	"bufio"
	"fmt"
	"os"
)

// tomatoRipening은 모든 토마토가 익는 데 걸리는 최소 일수를 반환한다.
//
// [매개변수]
//   - box: N×M 크기의 상자 (1: 익은 토마토, 0: 안 익은 토마토, -1: 빈 칸)
//   - n: 행 수
//   - m: 열 수
//
// [반환값]
//   - int: 모든 토마토가 익는 최소 일수 (이미 다 익었으면 0, 불가능하면 -1)
func tomatoRipening(box [][]int, n, m int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 열 수 M, 행 수 N 입력
	var m, n int
	fmt.Fscan(reader, &m, &n)

	// 상자 정보 입력
	box := make([][]int, n)
	for i := 0; i < n; i++ {
		box[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &box[i][j])
		}
	}

	// 핵심 함수 호출
	result := tomatoRipening(box, n, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
