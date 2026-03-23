package main

import (
	"bufio"
	"fmt"
	"os"
)

// enumerateSubsets는 N개 원소의 모든 부분집합을 비트마스크로 열거하여 반환한다.
//
// [매개변수]
//   - n: 원소의 수
//
// [반환값]
//   - [][]int: 각 부분집합에 포함된 원소 목록 (1-based)
func enumerateSubsets(n int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	subsets := enumerateSubsets(n)

	// 결과 출력
	for _, subset := range subsets {
		for i, v := range subset {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
