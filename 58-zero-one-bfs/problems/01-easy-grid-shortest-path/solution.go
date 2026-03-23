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
func gridShortestPath(grid []string, n, m int) int {
	// 여기에 코드를 작성하세요
	return 0
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
