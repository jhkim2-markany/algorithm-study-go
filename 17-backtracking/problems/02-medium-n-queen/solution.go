package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n     int
	count int
	col   []bool // col[c]: c열에 퀸이 있는지
	diag1 []bool // diag1[r+c]: 우상향 대각선에 퀸이 있는지
	diag2 []bool // diag2[r-c+n-1]: 좌상향 대각선에 퀸이 있는지
)

// solve 함수는 row번째 행에 퀸을 배치한다
func solve(row int) {
	// 종료 조건: 모든 행에 퀸을 배치했으면 경우의 수 증가
	if row == n {
		count++
		return
	}

	for c := 0; c < n; c++ {
		// 가지치기: 같은 열 또는 대각선에 퀸이 있으면 건너뛴다
		if col[c] || diag1[row+c] || diag2[row-c+n-1] {
			continue
		}

		// 퀸 배치
		col[c] = true
		diag1[row+c] = true
		diag2[row-c+n-1] = true

		// 다음 행으로 재귀 호출
		solve(row + 1)

		// 되돌리기: 퀸 제거
		col[c] = false
		diag1[row+c] = false
		diag2[row-c+n-1] = false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	fmt.Fscan(reader, &n)

	col = make([]bool, n)
	diag1 = make([]bool, 2*n)
	diag2 = make([]bool, 2*n)

	// 백트래킹으로 N-Queen 풀기
	solve(0)

	// 결과 출력
	fmt.Fprintln(writer, count)
}
