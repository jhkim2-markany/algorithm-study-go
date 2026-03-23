package main

import (
	"bufio"
	"fmt"
	"os"
)

// solveSudoku는 9×9 스도쿠 퍼즐을 풀어 완성된 보드를 반환한다.
//
// [매개변수]
//   - board: 9×9 스도쿠 보드 (0은 빈 칸)
//
// [반환값]
//   - [9][9]int: 완성된 스도쿠 보드
func solveSudoku(board [9][9]int) [9][9]int {
	// 여기에 코드를 작성하세요
	return board
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 스도쿠 입력
	var board [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fscan(reader, &board[i][j])
		}
	}

	// 핵심 함수 호출
	result := solveSudoku(board)

	// 결과 출력
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, result[i][j])
		}
		fmt.Fprintln(writer)
	}
}
