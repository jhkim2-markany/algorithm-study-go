package main

import (
	"bufio"
	"fmt"
	"os"
)

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

		// 스택을 이용한 괄호 검사
		stack := []rune{}
		valid := true

		for _, ch := range s {
			switch ch {
			case '(', '{', '[':
				// 여는 괄호는 스택에 Push
				stack = append(stack, ch)
			case ')':
				// 닫는 괄호: 스택이 비어 있거나 짝이 맞지 않으면 실패
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					valid = false
				} else {
					stack = stack[:len(stack)-1]
				}
			case '}':
				if len(stack) == 0 || stack[len(stack)-1] != '{' {
					valid = false
				} else {
					stack = stack[:len(stack)-1]
				}
			case ']':
				if len(stack) == 0 || stack[len(stack)-1] != '[' {
					valid = false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if !valid {
				break
			}
		}

		// 스택이 비어 있어야 모든 괄호가 짝지어진 것
		if valid && len(stack) == 0 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
