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

	// 문자열 입력
	var s string
	fmt.Fscan(reader, &s)

	// 각 문자의 등장 횟수를 추적하는 배열 (영문 소문자 26개)
	var count [26]int
	maxLen := 0
	left := 0

	// 슬라이딩 윈도우: right 포인터를 확장하며 탐색
	for right := 0; right < len(s); right++ {
		ch := s[right] - 'a'
		count[ch]++

		// 중복 문자가 발생하면 left를 이동하여 중복 해소
		for count[ch] > 1 {
			count[s[left]-'a']--
			left++
		}

		// 현재 윈도우 길이로 최대값 갱신
		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, maxLen)
}
