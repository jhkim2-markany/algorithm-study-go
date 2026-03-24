package main

import (
	"bufio"
	"fmt"
	"os"
)

// minimumMoves는 격자에서 시작 위치부터 목표 위치까지의 최소 이동 횟수를 반환한다.
//
// [매개변수]
//   - grid: N × N 격자 (각 행은 문자열)
//   - startX: 시작 행
//   - startY: 시작 열
//   - goalX: 목표 행
//   - goalY: 목표 열
//
// [반환값]
//   - int: 최소 이동 횟수
func minimumMoves(grid []string, startX, startY, goalX, goalY int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 격자 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 격자 입력
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	// 시작/목표 위치 입력
	var startX, startY, goalX, goalY int
	fmt.Fscan(reader, &startX, &startY, &goalX, &goalY)

	// 핵심 함수 호출
	result := minimumMoves(grid, startX, startY, goalX, goalY)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
