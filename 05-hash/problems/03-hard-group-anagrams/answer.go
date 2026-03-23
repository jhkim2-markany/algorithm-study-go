package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// sortString은 문자열의 문자를 정렬하여 아나그램의 대표 키를 생성한다.
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// groupAnagrams는 문자열 배열에서 아나그램 관계인 문자열들을 그룹화하여 반환한다.
//
// [매개변수]
//   - words: 알파벳 소문자로 이루어진 문자열 배열
//
// [반환값]
//   - [][]string: 아나그램 그룹 배열 (각 그룹 내 사전순, 그룹 간 첫 문자열 기준 사전순)
//
// [알고리즘 힌트]
//
//	각 문자열을 정렬하여 아나그램의 대표 키를 생성한다.
//	해시맵에 대표 키를 기준으로 원본 문자열을 그룹화한다.
//	예: "eat" → "aet", "tea" → "aet" → 같은 그룹
//	각 그룹 내 문자열을 사전순 정렬하고,
//	그룹 간에는 첫 번째 문자열 기준으로 사전순 정렬한다.
//
//	시간복잡도: O(N × K log K), K는 문자열 최대 길이
func groupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)
	for _, w := range words {
		key := sortString(w)
		groups[key] = append(groups[key], w)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		sort.Strings(group)
		result = append(result, group)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
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
