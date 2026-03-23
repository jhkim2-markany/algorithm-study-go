package main

import (
	"bufio"
	"fmt"
	"os"
)

// findPattern은 텍스트에서 패턴이 등장하는 모든 위치를 찾는다.
//
// [매개변수]
//   - text: 검색 대상 문자열
//   - pattern: 찾을 패턴 문자열
//
// [반환값]
//   - []int: 패턴이 등장하는 위치 목록 (1-based 인덱스)
func findPattern(text, pattern string) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var text, pattern string
	fmt.Fscan(reader, &text)
	fmt.Fscan(reader, &pattern)

	// 핵심 함수 호출
	positions := findPattern(text, pattern)

	// 결과 출력
	fmt.Fprintln(writer, len(positions))
	if len(positions) > 0 {
		for i, pos := range positions {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, pos)
		}
		fmt.Fprintln(writer)
	}
}
