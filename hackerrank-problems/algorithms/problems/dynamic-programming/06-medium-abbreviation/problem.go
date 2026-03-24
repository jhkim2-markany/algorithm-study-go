package main

import (
	"bufio"
	"fmt"
	"os"
)

// abbreviation은 문자열 a를 변환하여 b와 같게 만들 수 있는지 판별한다.
//
// [매개변수]
//   - a: 원본 문자열 (대소문자 혼합)
//   - b: 목표 문자열 (대문자만)
//
// [반환값]
//   - string: "YES" 또는 "NO"
func abbreviation(a, b string) string {
	// 여기에 코드를 작성하세요
	return "NO"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 쿼리 수 입력
	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		// 문자열 a, b 입력
		var a, b string
		fmt.Fscan(reader, &a)
		fmt.Fscan(reader, &b)

		// 핵심 함수 호출 및 결과 출력
		result := abbreviation(a, b)
		fmt.Fprintln(writer, result)
	}
}
