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

	// 문자열 S와 T 입력
	var s, t string
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &t)

	// T에 포함된 각 문자의 필요 횟수를 기록
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// 윈도우 내 각 문자의 등장 횟수
	var window [128]int

	// formed: T의 문자 중 윈도우에서 필요 횟수를 충족한 고유 문자 수
	// required: T에 포함된 고유 문자 수
	required := 0
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			required++
		}
	}

	formed := 0
	left := 0
	minLen := len(s) + 1

	// 슬라이딩 윈도우: right를 확장하며 T의 모든 문자를 포함하는 구간 탐색
	for right := 0; right < len(s); right++ {
		ch := s[right]
		window[ch]++

		// 현재 문자가 T에 필요한 만큼 충족되었는지 확인
		if need[ch] > 0 && window[ch] == need[ch] {
			formed++
		}

		// T의 모든 문자가 충족되면 left를 축소하며 최소 길이 갱신
		for formed == required {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}

			// left 문자를 윈도우에서 제거
			leftCh := s[left]
			window[leftCh]--
			if need[leftCh] > 0 && window[leftCh] < need[leftCh] {
				formed--
			}
			left++
		}
	}

	// 결과 출력
	if minLen > len(s) {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, minLen)
	}
}
