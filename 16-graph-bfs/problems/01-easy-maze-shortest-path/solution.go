package main

import (
	"bufio"
	"fmt"
	"os"
)

// mazeShortestPath는 미로에서 (0,0)부터 (n-1,m-1)까지의 최단 거리를 반환한다.
//
// [매개변수]
//   - maze: N×M 크기의 미로 ('1'은 이동 가능, '0'은 벽)
//   - n: 행 수
//   - m: 열 수
//
// [반환값]
//   - int: 시작점에서 도착점까지의 최단 거리 (시작 칸 포함)
func mazeShortestPath(maze []string, n, m int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 미로 크기 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 미로 입력
	maze := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &maze[i])
	}

	// 핵심 함수 호출
	result := mazeShortestPath(maze, n, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
