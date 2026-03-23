package main

import (
	"bufio"
	"fmt"
	"os"
)

// computeFailure 함수는 패턴의 실패 함수(부분 일치 테이블)를 계산한다
func computeFailure(pattern string) []int {
	m := len(pattern)
	fail := make([]int, m)

	// j: 현재까지 일치한 접두사 길이
	j := 0
	for i := 1; i < m; i++ {
		// 불일치 시 실패 함수를 따라가며 일치 위치를 찾는다
		for j > 0 && pattern[i] != pattern[j] {
			j = fail[j-1]
		}
		// 문자가 일치하면 접두사 길이를 증가시킨다
		if pattern[i] == pattern[j] {
			j++
		}
		fail[i] = j
	}
	return fail
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 텍스트와 패턴 입력
	var text, pattern string
	fmt.Fscan(reader, &text)
	fmt.Fscan(reader, &pattern)

	n := len(text)
	m := len(pattern)

	// 실패 함수를 미리 계산한다
	fail := computeFailure(pattern)

	// KMP 알고리즘으로 패턴 검색
	positions := []int{}
	j := 0 // 패턴에서 현재 비교 중인 위치
	for i := 0; i < n; i++ {
		// 불일치 시 실패 함수를 이용해 패턴 위치를 조정한다
		for j > 0 && text[i] != pattern[j] {
			j = fail[j-1]
		}
		// 현재 문자가 일치하면 패턴 포인터를 전진시킨다
		if text[i] == pattern[j] {
			j++
		}
		// 패턴 전체가 일치하면 시작 위치를 기록한다 (1-based)
		if j == m {
			positions = append(positions, i-m+2)
			// 다음 매칭을 위해 실패 함수로 위치를 조정한다
			j = fail[j-1]
		}
	}

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
