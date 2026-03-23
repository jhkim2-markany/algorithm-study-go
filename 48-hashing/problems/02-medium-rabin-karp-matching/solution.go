package main

import (
	"bufio"
	"fmt"
	"os"
)

// rabinKarpSearch는 라빈-카프 알고리즘으로 텍스트에서 패턴의 모든 출현 위치를 찾는다.
//
// [매개변수]
//   - text: 검색 대상 텍스트
//   - pattern: 찾을 패턴 문자열
//
// [반환값]
//   - []int: 패턴이 출현하는 1-indexed 위치 배열
func rabinKarpSearch(text, pattern string) []int {
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

	positions := rabinKarpSearch(text, pattern)

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
