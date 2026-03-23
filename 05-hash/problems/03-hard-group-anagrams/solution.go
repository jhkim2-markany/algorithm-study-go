package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// groupAnagrams는 문자열 배열에서 아나그램 관계인 문자열들을 그룹화하여 반환한다.
//
// [매개변수]
//   - words: 알파벳 소문자로 이루어진 문자열 배열
//
// [반환값]
//   - [][]string: 아나그램 그룹 배열 (각 그룹 내 사전순, 그룹 간 첫 문자열 기준 사전순)
func groupAnagrams(words []string) [][]string {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 문자열 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 문자열 입력
	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &words[i])
	}

	// 핵심 함수 호출
	result := groupAnagrams(words)

	// 결과 출력
	for _, group := range result {
		fmt.Fprintln(writer, strings.Join(group, " "))
	}
}
