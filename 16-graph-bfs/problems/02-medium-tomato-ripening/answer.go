package main

import (
	"bufio"
	"fmt"
	"os"
)

// tomatoRipening은 모든 토마토가 익는 데 걸리는 최소 일수를 반환한다.
//
// [매개변수]
//   - box: N×M 크기의 상자 (1: 익은 토마토, 0: 안 익은 토마토, -1: 빈 칸)
//   - n: 행 수
//   - m: 열 수
//
// [반환값]
//   - int: 모든 토마토가 익는 최소 일수 (이미 다 익었으면 0, 불가능하면 -1)
//
// [알고리즘 힌트]
//
//	다중 시작점 BFS를 사용한다.
//	모든 익은 토마토(값 1)를 큐에 넣고 동시에 BFS를 시작한다.
//	인접한 안 익은 토마토(값 0)를 익게 만들며 날짜를 누적한다.
//	BFS 완료 후 안 익은 토마토가 남아있으면 -1을 반환한다.
func tomatoRipening(box [][]int, n, m int) int {
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}

	type point struct{ x, y int }
	var queue []point

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if box[i][j] == 1 {
				queue = append(queue, point{i, j})
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for d := 0; d < 4; d++ {
			nx, ny := cur.x+dx[d], cur.y+dy[d]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if box[nx][ny] != 0 {
				continue
			}
			box[nx][ny] = box[cur.x][cur.y] + 1
			queue = append(queue, point{nx, ny})
		}
	}

	maxDay := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if box[i][j] == 0 {
				return -1
			}
			if box[i][j] > maxDay {
				maxDay = box[i][j]
			}
		}
	}

	return maxDay - 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 열 수 M, 행 수 N 입력
	var m, n int
	fmt.Fscan(reader, &m, &n)

	// 상자 정보 입력
	box := make([][]int, n)
	for i := 0; i < n; i++ {
		box[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &box[i][j])
		}
	}

	// 핵심 함수 호출
	result := tomatoRipening(box, n, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
