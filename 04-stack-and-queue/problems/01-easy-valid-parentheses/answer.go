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
//
// [알고리즘 힌트]
//
//	스택을 사용하여 여는 괄호를 Push하고,
//	닫는 괄호를 만나면 스택 맨 위와 짝이 맞는지 확인한다.
//	짝이 맞으면 Pop, 맞지 않으면 false를 반환한다.
//	문자열을 모두 처리한 후 스택이 비어 있으면 올바른 괄호 문자열이다.
//
//	시간복잡도: O(N), 공간복잡도: O(N)
func isValidParentheses(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
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
