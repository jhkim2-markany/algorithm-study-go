package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestCommonSubsequence는 두 수열의 최장 공통 부분 수열을 반환한다.
//
// [매개변수]
//   - a: 첫 번째 수열
//   - b: 두 번째 수열
//
// [반환값]
//   - []int: 최장 공통 부분 수열
func longestCommonSubsequence(a, b []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 수열 길이 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 첫 번째 수열 입력
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 두 번째 수열 입력
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := longestCommonSubsequence(a, b)
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
