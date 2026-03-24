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
func isBalanced(s string) string {
	// 여기에 코드를 작성하세요
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
