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

	// 텍스트와 패턴 입력
	var text, pattern string
	fmt.Fscan(reader, &text)
	fmt.Fscan(reader, &pattern)

	n := len(text)
	m := len(pattern)

	// 브루트포스로 패턴이 등장하는 모든 위치를 찾는다
	positions := []int{}
	for i := 0; i <= n-m; i++ {
		match := true
		// 현재 위치에서 패턴의 각 문자를 비교한다
		for j := 0; j < m; j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		// 모든 문자가 일치하면 위치를 기록한다 (1-based 인덱스)
		if match {
			positions = append(positions, i+1)
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
