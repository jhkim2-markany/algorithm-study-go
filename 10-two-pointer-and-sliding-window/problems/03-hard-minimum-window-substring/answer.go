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
//
// [알고리즘 힌트]
//
//	슬라이딩 윈도우: need 배열에 t의 각 문자 필요 횟수를 기록한다.
//	right를 확장하며 윈도우에 문자를 추가하고,
//	t의 모든 문자가 충족되면 left를 축소하며 최소 길이를 갱신한다.
//	formed(충족된 고유 문자 수) == required(t의 고유 문자 수)일 때 조건 만족.
//
//	시간복잡도: O(|S| + |T|), 공간복잡도: O(128)
func minWindowSubstring(s, t string) int {
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := 0
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			required++
		}
	}

	var window [128]int
	formed := 0
	left := 0
	minLen := len(s) + 1

	for right := 0; right < len(s); right++ {
		ch := s[right]
		window[ch]++
		if need[ch] > 0 && window[ch] == need[ch] {
			formed++
		}

		for formed == required {
			if length := right - left + 1; length < minLen {
				minLen = length
			}
			leftCh := s[left]
			window[leftCh]--
			if need[leftCh] > 0 && window[leftCh] < need[leftCh] {
				formed--
			}
			left++
		}
	}

	if minLen > len(s) {
		return 0
	}
	return minLen
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
