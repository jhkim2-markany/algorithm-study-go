package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

// minWallBreaks는 격자에서 (0,0)부터 (n-1,m-1)까지 벽을 최소로 부수며 도달하는 비용을 구한다.
// '0'은 빈 칸(비용 0), '1'은 벽(비용 1)이며, 덱 기반 0-1 BFS를 사용한다.
//
// [매개변수]
//   - grid: 격자 정보 ('0'은 빈 칸, '1'은 벽)
//   - n: 격자의 행 수
//   - m: 격자의 열 수
//
// [반환값]
//   - int: (0,0)에서 (n-1,m-1)까지 부숴야 하는 벽의 최소 개수
//
// [알고리즘 힌트]
//
//	0-1 BFS를 사용하여 벽 부수기 최소 비용을 구한다.
//	빈 칸('0')으로 이동하면 비용 0이므로 덱 앞에 추가하고,
//	벽('1')을 부수고 이동하면 비용 1이므로 덱 뒤에 추가한다.
//	Pos 구조체로 좌표를 관리하며, grid[nr][nc]-'0'으로 비용을 계산한다.
func minWallBreaks(grid []string, n, m int) int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = INF
		}
	}
	dist[0][0] = 0

	type Pos struct{ r, c int }
	deque := []Pos{{0, 0}}

	for len(deque) > 0 {
		cur := deque[0]
		deque = deque[1:]

		for d := 0; d < 4; d++ {
			nr, nc := cur.r+dr[d], cur.c+dc[d]
			if nr < 0 || nr >= n || nc < 0 || nc >= m {
				continue
			}

			w := int(grid[nr][nc] - '0')

			newDist := dist[cur.r][cur.c] + w
			if newDist < dist[nr][nc] {
				dist[nr][nc] = newDist
				if w == 0 {
					deque = append([]Pos{{nr, nc}}, deque...)
				} else {
					deque = append(deque, Pos{nr, nc})
				}
			}
		}
	}

	return dist[n-1][m-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 격자 크기
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 입력: 격자 정보 (0: 빈 칸, 1: 벽)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, minWallBreaks(grid, n, m))
}
