package main

import (
	"bufio"
	"fmt"
	"os"
)

// solveNQueen은 N×N 체스판에 N개의 퀸을 서로 공격하지 않게 배치하는 경우의 수를 반환한다.
//
// [매개변수]
//   - n: 체스판의 크기 및 퀸의 개수
//
// [반환값]
//   - int: 가능한 배치의 수
//
// [알고리즘 힌트]
//
//	행 단위로 퀸을 배치하며 백트래킹한다.
//	열(col), 우상향 대각선(diag1: r+c), 좌상향 대각선(diag2: r-c+n-1)에
//	퀸이 있는지 불리언 배열로 추적한다.
//	충돌이 없는 열에만 퀸을 배치하고, 모든 행을 채우면 카운트를 증가시킨다.
func solveNQueen(n int) int {
	col := make([]bool, n)
	diag1 := make([]bool, 2*n)
	diag2 := make([]bool, 2*n)
	count := 0

	var solve func(row int)
	solve = func(row int) {
		if row == n {
			count++
			return
		}
		for c := 0; c < n; c++ {
			if col[c] || diag1[row+c] || diag2[row-c+n-1] {
				continue
			}
			col[c] = true
			diag1[row+c] = true
			diag2[row-c+n-1] = true
			solve(row + 1)
			col[c] = false
			diag1[row+c] = false
			diag2[row-c+n-1] = false
		}
	}

	solve(0)
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	result := solveNQueen(n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
