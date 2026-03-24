package main

import (
	"bufio"
	"fmt"
	"os"
)

// isBalanced는 괄호 문자열이 균형 잡혀 있는지 판별한다.
//
// [매개변수]
//   - s: 괄호 문자열
//
// [반환값]
//   - string: 균형 잡혀 있으면 "YES", 아니면 "NO"
//
// [알고리즘 힌트]
//
//	스택을 사용하여 여는 괄호를 저장하고,
//	닫는 괄호를 만나면 스택 최상위와 짝이 맞는지 확인한다.
func isBalanced(s string) string {
	// 괄호 매칭 맵 정의
	matching := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 스택 초기화
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' || ch == '[' || ch == '{' {
			// 여는 괄호는 스택에 푸시
			stack = append(stack, ch)
		} else {
			// 닫는 괄호: 스택이 비어있거나 짝이 맞지 않으면 실패
			if len(stack) == 0 || stack[len(stack)-1] != matching[ch] {
				return "NO"
			}
			// 짝이 맞으면 팝
			stack = stack[:len(stack)-1]
		}
	}

	// 스택이 비어있으면 균형 잡힌 문자열
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 개수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(reader, &s)

		// 핵심 함수 호출 및 결과 출력
		fmt.Fprintln(writer, isBalanced(s))
	}
}
