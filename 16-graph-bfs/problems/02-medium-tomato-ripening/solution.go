package main

import (
	"bufio"
	"fmt"
	"os"
)

// 상하좌우 방향 배열
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 열 수 M, 행 수 N 입력
	var m, n int
	fmt.Fscan(reader, &m, &n)

	// 상자 정보 입력
	box := make([][]int, n)
	type point struct{ x, y int }
	queue := []point{}

	for i := 0; i < n; i++ {
		box[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &box[i][j])
			// 익은 토마토를 모두 큐에 넣는다 (다중 시작점 BFS)
			if box[i][j] == 1 {
				queue = append(queue, point{i, j})
			}
		}
	}

	// 다중 시작점 BFS 수행
	// 익은 토마토의 값(1)을 기준으로 날짜를 누적한다
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx[d], cur.y+dy[d]

			// 범위 확인
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			// 익지 않은 토마토만 처리
			if box[nx][ny] != 0 {
				continue
			}

			// 인접한 익지 않은 토마토를 익게 만들고 날짜를 기록
			box[nx][ny] = box[cur.x][cur.y] + 1
			queue = append(queue, point{nx, ny})
		}
	}

	// 최대 날짜를 구하고, 익지 않은 토마토가 있는지 확인
	maxDay := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if box[i][j] == 0 {
				// 익지 못한 토마토가 존재
				fmt.Fprintln(writer, -1)
				return
			}
			if box[i][j] > maxDay {
				maxDay = box[i][j]
			}
		}
	}

	// 시작값이 1이었으므로 1을 빼서 일수를 구한다
	fmt.Fprintln(writer, maxDay-1)
}
