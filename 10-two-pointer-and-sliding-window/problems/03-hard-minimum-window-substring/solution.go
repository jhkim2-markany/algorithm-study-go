package main

import (
	"bufio"
	"fmt"
	"os"
)

// minWindowSubstring은 문자열 s에서 t의 모든 문자를 포함하는
// 가장 짧은 부분 문자열의 길이를 반환한다.
//
// [매개변수]
//   - s: 원본 문자열
//   - t: 포함해야 하는 문자들로 이루어진 문자열
//
// [반환값]
//   - int: 조건을 만족하는 최소 부분 문자열의 길이 (없으면 0)
func minWindowSubstring(s, t string) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 문자열 S와 T 입력
	var s, t string
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &t)

	// 핵심 함수 호출
	result := minWindowSubstring(s, t)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
