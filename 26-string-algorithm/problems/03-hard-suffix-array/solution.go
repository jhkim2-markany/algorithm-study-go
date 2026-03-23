package main

import (
	"bufio"
	"fmt"
	"os"
)

// buildSuffixArray는 문자열의 접미사 배열을 구축한다.
//
// [매개변수]
//   - s: 입력 문자열
//
// [반환값]
//   - []int: 접미사 배열 (각 접미사의 시작 인덱스)
func buildSuffixArray(s string) []int {
	// 여기에 코드를 작성하세요
	return nil
}

// buildLCP는 접미사 배열로부터 LCP 배열을 구축한다.
//
// [매개변수]
//   - s: 입력 문자열
//   - sa: 접미사 배열
//
// [반환값]
//   - []int: LCP 배열 (인접한 접미사 간의 최장 공통 접두사 길이)
func buildLCP(s string, sa []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	// 핵심 함수 호출
	sa := buildSuffixArray(s)
	lcp := buildLCP(s, sa)

	// 접미사 배열 출력
	for i, v := range sa {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	// LCP 배열 출력
	for i, v := range lcp {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
