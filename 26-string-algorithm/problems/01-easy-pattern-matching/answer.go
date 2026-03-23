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
//
// [알고리즘 힌트]
//
//	브루트포스로 텍스트의 각 위치에서 패턴의 모든 문자를 비교한다.
//	텍스트의 i번째 위치부터 패턴 길이만큼 비교하여
//	모든 문자가 일치하면 해당 위치(1-based)를 기록한다.
func findPattern(text, pattern string) []int {
	n := len(text)
	m := len(pattern)

	positions := []int{}
	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			positions = append(positions, i+1)
		}
	}
	return positions
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
