package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// sortString 함수는 문자열의 문자를 정렬하여 아나그램의 대표 키를 생성한다.
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 문자열 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 해시맵: 정렬된 문자열(키) → 원본 문자열 목록(값)
	groups := make(map[string][]string)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)

		// 문자열을 정렬하여 아나그램 대표 키 생성
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}

	// 각 그룹 내 문자열을 사전순 정렬
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		sort.Strings(group)
		result = append(result, group)
	}

	// 그룹 간 정렬: 각 그룹의 첫 번째 문자열 기준 사전순
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	// 결과 출력
	for _, group := range result {
		fmt.Fprintln(writer, strings.Join(group, " "))
	}
}
