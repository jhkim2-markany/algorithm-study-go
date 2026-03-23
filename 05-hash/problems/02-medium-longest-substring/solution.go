package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestSubstring은 중복 문자가 없는 가장 긴 부분 문자열의 길이를 반환한다.
//
// [매개변수]
//   - s: 알파벳 소문자로 이루어진 문자열
//
// [반환값]
//   - int: 중복 문자가 없는 가장 긴 부분 문자열의 길이
func longestSubstring(s string) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 문자열 입력
	var s string
	fmt.Fscan(reader, &s)

	// 핵심 함수 호출
	result := longestSubstring(s)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
