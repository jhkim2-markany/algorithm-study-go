package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestSubstring은 중복 문자가 없는 가장 긴 부분 문자열의 길이를 반환한다.
//
// [매개변수]
//   - s: 알파벳 소문자로 이루어진 문자열
//
// [반환값]
//   - int: 중복 문자가 없는 가장 긴 부분 문자열의 길이
//
// [알고리즘 힌트]
//
//	슬라이딩 윈도우 + 해시맵을 사용한다.
//	해시맵에 각 문자의 마지막 등장 인덱스를 저장하고,
//	현재 문자가 윈도우 내에 이미 존재하면 왼쪽 포인터를
//	해당 문자의 다음 위치로 이동시킨다.
//	매 단계에서 윈도우 크기(right - left + 1)의 최댓값을 갱신한다.
//
//	시간복잡도: O(N), 공간복잡도: O(min(N, 26))
func longestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, ok := lastSeen[ch]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[ch] = right
		if windowLen := right - left + 1; windowLen > maxLen {
			maxLen = windowLen
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
	result := longestSubstring(s)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
