package main

import (
	"bufio"
	"fmt"
	"os"
)

// isValidParentheses는 괄호 문자열이 올바르게 짝지어져 있는지 판별한다.
//
// [매개변수]
//   - s: 소괄호, 중괄호, 대괄호로만 이루어진 문자열
//
// [반환값]
//   - bool: 올바른 괄호 문자열이면 true, 아니면 false
func isValidParentheses(s string) bool {
	// 여기에 코드를 작성하세요
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		// 괄호 문자열 입력
		var s string
		fmt.Fscan(reader, &s)

		// 핵심 함수 호출
		if isValidParentheses(s) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
