package main

import (
	"bufio"
	"fmt"
	"os"
)

// countLuck은 시작점에서 포트키까지 경로의 갈림길 수가 k와 같은지 판별한다.
//
// [매개변수]
//   - matrix: N×M 격자 ('.' 이동 가능, 'X' 나무, 'M' 시작, '*' 포트키)
//   - k: 예상 갈림길 수
//
// [반환값]
//   - string: "Impressed" 또는 "Oops!"
//
// [알고리즘 힌트]
//
//	DFS로 시작점에서 포트키까지 경로를 찾으면서
//	이동 가능한 방향이 2개 이상인 갈림길을 센다.
func countLuck(matrix []string, k int) string {
	n := len(matrix)
	m := len(matrix[0])

	// 시작 위치 찾기
	var sr, sc int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 'M' {
				sr, sc = i, j
			}
		}
	}

	// 방문 배열
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	// 4방향 이동 벡터
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	// DFS: 포트키까지의 경로에서 갈림길 수를 반환
	var dfs func(r, c int) (bool, int)
	dfs = func(r, c int) (bool, int) {
		visited[r][c] = true

		// 포트키에 도달
		if matrix[r][c] == '*' {
			return true, 0
		}

		// 현재 셀에서 이동 가능한 방향 수 계산
		choices := 0
		for d := 0; d < 4; d++ {
			nr, nc := r+dr[d], c+dc[d]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && !visited[nr][nc] && matrix[nr][nc] != 'X' {
				choices++
			}
		}

		// 인접 셀 탐색
		for d := 0; d < 4; d++ {
			nr, nc := r+dr[d], c+dc[d]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && !visited[nr][nc] && matrix[nr][nc] != 'X' {
				found, wands := dfs(nr, nc)
				if found {
					// 갈림길이면 지팡이 횟수 추가
					if choices > 1 {
						wands++
					}
					return true, wands
				}
			}
		}

		return false, 0
	}

	_, wands := dfs(sr, sc)

	if wands == k {
		return "Impressed"
	}
	return "Oops!"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		matrix := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &matrix[i])
		}

		var k int
		fmt.Fscan(reader, &k)

		fmt.Fprintln(writer, countLuck(matrix, k))
	}
}
