package main

import (
	"bufio"
	"fmt"
	"os"
)

// samAndSubstrings는 숫자 문자열의 모든 부분 문자열 합을 반환한다.
//
// [매개변수]
//   - s: 숫자 문자열
//
// [반환값]
//   - int: 모든 부분 문자열의 합 (mod 10^9+7)
func samAndSubstrings(s string) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 숫자 문자열 입력
	var s string
	fmt.Fscan(reader, &s)

	// 핵심 함수 호출 및 결과 출력
	result := samAndSubstrings(s)
	fmt.Fprintln(writer, result)
}
