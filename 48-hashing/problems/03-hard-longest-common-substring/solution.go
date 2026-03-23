package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestCommonSubstringLen은 이분 탐색과 해싱을 이용하여
// 두 문자열의 가장 긴 공통 부분 문자열의 길이를 반환한다.
//
// [매개변수]
//   - a: 첫 번째 문자열
//   - b: 두 번째 문자열
//
// [반환값]
//   - int: 가장 긴 공통 부분 문자열의 길이
func longestCommonSubstringLen(a, b string) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	fmt.Fprintln(writer, longestCommonSubstringLen(a, b))
}
