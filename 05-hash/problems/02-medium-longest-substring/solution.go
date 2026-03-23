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

	// 해시맵: 각 문자의 마지막 등장 인덱스를 저장
	lastSeen := make(map[byte]int)

	maxLen := 0
	left := 0

	// 슬라이딩 윈도우로 중복 없는 가장 긴 부분 문자열을 찾는다
	for right := 0; right < len(s); right++ {
		ch := s[right]

		// 현재 문자가 윈도우 내에 이미 존재하면 왼쪽 포인터를 이동
		if idx, ok := lastSeen[ch]; ok && idx >= left {
			left = idx + 1
		}

		// 현재 문자의 위치를 해시맵에 갱신
		lastSeen[ch] = right

		// 윈도우 크기의 최댓값 갱신
		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, maxLen)
}
