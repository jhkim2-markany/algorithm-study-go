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
//
// [알고리즘 힌트]
//
//	빈 칸 목록을 미리 수집한 뒤 백트래킹으로 채운다.
//	행(row), 열(col), 3×3 박스(box)에 대한 제약 조건을 불리언 배열로 관리한다.
//	각 빈 칸에 1~9를 시도하며, 제약 조건에 위배되면 가지치기한다.
//	모든 빈 칸을 채우면 성공, 실패하면 되돌린다.
func solveSudoku(board [9][9]int) [9][9]int {
	var row [9][10]bool
	var colB [9][10]bool
	var box [9][10]bool
	var blank [][2]int

	boxIdx := func(r, c int) int {
		return (r/3)*3 + c/3
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				blank = append(blank, [2]int{i, j})
			} else {
				num := board[i][j]
				row[i][num] = true
				colB[j][num] = true
				box[boxIdx(i, j)][num] = true
			}
		}
	}

	var solve func(idx int) bool
	solve = func(idx int) bool {
		if idx == len(blank) {
			return true
		}
		r, c := blank[idx][0], blank[idx][1]
		b := boxIdx(r, c)

		for num := 1; num <= 9; num++ {
			if row[r][num] || colB[c][num] || box[b][num] {
				continue
			}
			board[r][c] = num
			row[r][num] = true
			colB[c][num] = true
			box[b][num] = true

			if solve(idx + 1) {
				return true
			}

			board[r][c] = 0
			row[r][num] = false
			colB[c][num] = false
			box[b][num] = false
		}
		return false
	}

	solve(0)
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
