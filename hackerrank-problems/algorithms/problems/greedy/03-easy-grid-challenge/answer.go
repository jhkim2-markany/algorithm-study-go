package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// gridChallenge는 각 행을 재배열하여 모든 열이 알파벳 순서가 될 수 있는지 판별한다.
//
// [매개변수]
//   - grid: 문자열 배열 (N×N 그리드)
//
// [반환값]
//   - string: "YES" 또는 "NO"
//
// [알고리즘 힌트]
//
//	각 행을 정렬한 뒤, 열 방향으로 비내림차순인지 확인한다.
//	행 정렬이 그리디 선택이며, 이것이 최적임이 보장된다.
func gridChallenge(grid []string) string {
	n := len(grid)

	// 각 행을 알파벳 순서로 정렬
	sorted := make([][]byte, n)
	for i := 0; i < n; i++ {
		sorted[i] = []byte(grid[i])
		sort.Slice(sorted[i], func(a, b int) bool {
			return sorted[i][a] < sorted[i][b]
		})
	}

	// 각 열이 비내림차순인지 확인
	for col := 0; col < n; col++ {
		for row := 1; row < n; row++ {
			if sorted[row][col] < sorted[row-1][col] {
				return "NO"
			}
		}
	}

	return "YES"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		// 그리드 크기 입력
		var n int
		fmt.Fscan(reader, &n)

		// 그리드 입력
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &grid[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := gridChallenge(grid)
		fmt.Fprintln(writer, result)
	}
}
