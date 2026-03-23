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

// gridShortestPath는 격자에서 (0,0)부터 (n-1,m-1)까지의 0-1 BFS 최단 경로를 구한다.
// '.'은 비용 0, '#'은 비용 1로 이동하며, 덱 기반 0-1 BFS를 사용한다.
//
// [매개변수]
//   - grid: 격자 정보 ('.'은 빈 칸, '#'은 벽)
//   - n: 격자의 행 수
//   - m: 격자의 열 수
//
// [반환값]
//   - int: (0,0)에서 (n-1,m-1)까지의 최소 비용
//
// [알고리즘 힌트]
//
//	0-1 BFS: 간선 가중치가 0 또는 1인 그래프에서 덱(deque)을 사용하여
//	최단 거리를 구한다. 가중치 0인 간선은 덱 앞에, 가중치 1인 간선은
//	덱 뒤에 추가하여 다익스트라와 동일한 결과를 O(V+E)에 얻는다.
//	좌표를 r*m+c로 인코딩하여 덱에 저장한다.
func gridShortestPath(grid []string, n, m int) int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = INF
		}
	}
	dist[0][0] = 0

	// 덱을 슬라이스로 구현 (좌표를 r*m+c로 인코딩)
	deque := []int{0}

	for len(deque) > 0 {
		cur := deque[0]
		deque = deque[1:]
		r, c := cur/m, cur%m

		for d := 0; d < 4; d++ {
			nr, nc := r+dr[d], c+dc[d]
			if nr < 0 || nr >= n || nc < 0 || nc >= m {
				continue
			}

			w := 0
			if grid[nr][nc] == '#' {
				w = 1
			}

			newDist := dist[r][c] + w
			if newDist < dist[nr][nc] {
				dist[nr][nc] = newDist
				encoded := nr*m + nc
				if w == 0 {
					deque = append([]int{encoded}, deque...)
				} else {
					deque = append(deque, encoded)
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

	// 입력: 격자 정보
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, gridShortestPath(grid, n, m))
}
