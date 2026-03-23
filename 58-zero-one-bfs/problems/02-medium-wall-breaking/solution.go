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

	// 거리 배열 초기화 (부숴야 하는 벽의 최소 개수)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = INF
		}
	}
	dist[0][0] = 0

	// 덱 기반 0-1 BFS
	type Pos struct{ r, c int }
	deque := []Pos{{0, 0}}

	for len(deque) > 0 {
		// 덱 앞에서 꺼내기
		cur := deque[0]
		deque = deque[1:]

		for d := 0; d < 4; d++ {
			nr, nc := cur.r+dr[d], cur.c+dc[d]
			if nr < 0 || nr >= n || nc < 0 || nc >= m {
				continue // 범위 밖
			}

			// 빈 칸이면 비용 0, 벽이면 비용 1 (벽 부수기)
			w := int(grid[nr][nc] - '0')

			newDist := dist[cur.r][cur.c] + w
			if newDist < dist[nr][nc] {
				dist[nr][nc] = newDist
				if w == 0 {
					// 빈 칸: 덱 앞에 추가
					deque = append([]Pos{{nr, nc}}, deque...)
				} else {
					// 벽: 덱 뒤에 추가
					deque = append(deque, Pos{nr, nc})
				}
			}
		}
	}

	// 출력: 부숴야 하는 벽의 최소 개수
	fmt.Fprintln(writer, dist[n-1][m-1])
}
