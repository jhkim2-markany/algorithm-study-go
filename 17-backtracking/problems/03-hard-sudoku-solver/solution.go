package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	board [9][9]int
	row   [9][10]bool // row[r][num]: r행에 num이 있는지
	colB  [9][10]bool // colB[c][num]: c열에 num이 있는지
	box   [9][10]bool // box[b][num]: b번 박스에 num이 있는지
	blank [][2]int    // 빈 칸의 좌표 목록
)

// boxIdx 함수는 (r, c) 좌표가 속하는 3×3 박스 번호를 반환한다
func boxIdx(r, c int) int {
	return (r/3)*3 + c/3
}

// solve 함수는 idx번째 빈 칸부터 채워나간다
func solve(idx int) bool {
	// 종료 조건: 모든 빈 칸을 채웠으면 성공
	if idx == len(blank) {
		return true
	}

	r, c := blank[idx][0], blank[idx][1]
	b := boxIdx(r, c)

	// 1부터 9까지 시도
	for num := 1; num <= 9; num++ {
		// 가지치기: 행, 열, 박스에 이미 같은 숫자가 있으면 건너뛴다
		if row[r][num] || colB[c][num] || box[b][num] {
			continue
		}

		// 숫자 배치
		board[r][c] = num
		row[r][num] = true
		colB[c][num] = true
		box[b][num] = true

		// 다음 빈 칸으로 재귀 호출
		if solve(idx + 1) {
			return true
		}

		// 되돌리기: 숫자 제거
		board[r][c] = 0
		row[r][num] = false
		colB[c][num] = false
		box[b][num] = false
	}

	// 모든 숫자가 실패하면 이전 칸으로 되돌아간다
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 스도쿠 입력 및 초기 상태 설정
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fscan(reader, &board[i][j])
			if board[i][j] == 0 {
				// 빈 칸 좌표 저장
				blank = append(blank, [2]int{i, j})
			} else {
				// 이미 채워진 숫자를 제약 조건에 등록
				num := board[i][j]
				row[i][num] = true
				colB[j][num] = true
				box[boxIdx(i, j)][num] = true
			}
		}
	}

	// 백트래킹으로 스도쿠 풀기
	solve(0)

	// 결과 출력
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, board[i][j])
		}
		fmt.Fprintln(writer)
	}
}
