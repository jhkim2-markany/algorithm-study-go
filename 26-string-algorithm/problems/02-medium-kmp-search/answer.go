package main

import (
	"bufio"
	"fmt"
	"os"
)

// computeFailure는 패턴의 실패 함수(부분 일치 테이블)를 계산한다.
//
// [매개변수]
//   - pattern: 패턴 문자열
//
// [반환값]
//   - []int: 실패 함수 배열
//
// [알고리즘 힌트]
//
//	j는 현재까지 일치한 접두사 길이를 나타낸다.
//	불일치 시 실패 함수를 따라가며 일치 위치를 찾고,
//	문자가 일치하면 접두사 길이를 증가시킨다.
func computeFailure(pattern string) []int {
	m := len(pattern)
	fail := make([]int, m)

	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = fail[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		fail[i] = j
	}
	return fail
}

// kmpSearch는 KMP 알고리즘으로 텍스트에서 패턴이 등장하는 모든 위치를 찾는다.
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
//	실패 함수를 미리 계산한 후, 텍스트를 순회하며 패턴과 비교한다.
//	불일치 시 실패 함수를 이용해 패턴 위치를 조정하여
//	불필요한 비교를 건너뛴다. 패턴 전체가 일치하면 위치를 기록하고
//	실패 함수로 다음 매칭 위치를 조정한다.
func kmpSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)

	fail := computeFailure(pattern)

	positions := []int{}
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = fail[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			positions = append(positions, i-m+2)
			j = fail[j-1]
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
	positions := kmpSearch(text, pattern)

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
