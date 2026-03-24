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
func gridChallenge(grid []string) string {
	// 여기에 코드를 작성하세요
	return "NO"
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

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Strings
}
