package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

// colorFillQueries는 색상이 칠해진 격자에서 플러드 필 방식의
// 색상 변경 쿼리들을 유니온 파인드를 이용하여 효율적으로 처리한다.
// 각 쿼리는 (r, c) 위치와 연결된 같은 색상 영역 전체를 새 색상으로 변경한다.
//
// [매개변수]
//   - grid: n×m 크기의 격자 (각 칸에 색상 번호)
//   - n: 격자의 행 수
//   - m: 격자의 열 수
//   - queries: 쿼리 목록 ([r, c, newColor], 1-indexed 좌표)
//
// [반환값]
//   - [][]int: 모든 쿼리 처리 후의 최종 격자 상태
func colorFillQueries(grid [][]int, n, m int, queries [][3]int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	queries := make([][3]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}

	result := colorFillQueries(grid, n, m, queries)

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if c > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, result[r][c])
		}
		fmt.Fprintln(writer)
	}
}
