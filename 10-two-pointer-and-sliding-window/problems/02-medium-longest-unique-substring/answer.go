package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestUniqueSubstring은 중복 문자가 없는 가장 긴 부분 문자열의 길이를 반환한다.
//
// [매개변수]
//   - s: 알파벳 소문자로 이루어진 문자열
//
// [반환값]
//   - int: 중복 문자가 없는 가장 긴 부분 문자열의 길이
//
// [알고리즘 힌트]
//
//	슬라이딩 윈도우: 배열로 각 문자의 등장 횟수를 추적한다.
//	right 포인터를 확장하며 문자를 추가하고,
//	중복이 발생하면 left를 이동하여 중복을 해소한다.
//	매 단계에서 윈도우 크기의 최댓값을 갱신한다.
//
//	시간복잡도: O(N), 공간복잡도: O(26) = O(1)
func longestUniqueSubstring(s string) int {
	var count [26]int
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right] - 'a'
		count[ch]++
		for count[ch] > 1 {
			count[s[left]-'a']--
			left++
		}
		if length := right - left + 1; length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 문자열 입력
	var s string
	fmt.Fscan(reader, &s)

	// 핵심 함수 호출
	result := longestUniqueSubstring(s)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
