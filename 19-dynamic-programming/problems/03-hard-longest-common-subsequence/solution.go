package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestCommonSubsequence는 두 문자열의 LCS 길이와 LCS 문자열을 반환한다.
//
// [매개변수]
//   - a: 첫 번째 문자열
//   - b: 두 번째 문자열
//
// [반환값]
//   - int: LCS의 길이
//   - string: LCS 문자열 (길이가 0이면 빈 문자열)
func longestCommonSubsequence(a, b string) (int, string) {
	// 여기에 코드를 작성하세요
	return 0, ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	// 핵심 함수 호출
	lcsLen, lcsStr := longestCommonSubsequence(a, b)

	// LCS 길이 출력
	fmt.Fprintln(writer, lcsLen)

	// LCS 문자열 출력 (길이가 0이 아닌 경우)
	if lcsLen > 0 {
		fmt.Fprintln(writer, lcsStr)
	}
}
