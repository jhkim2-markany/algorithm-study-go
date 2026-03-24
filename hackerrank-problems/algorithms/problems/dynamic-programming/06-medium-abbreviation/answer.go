package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// abbreviation은 문자열 a를 변환하여 b와 같게 만들 수 있는지 판별한다.
//
// [매개변수]
//   - a: 원본 문자열 (대소문자 혼합)
//   - b: 목표 문자열 (대문자만)
//
// [반환값]
//   - string: "YES" 또는 "NO"
//
// [알고리즘 힌트]
//
//	2차원 DP를 사용한다. dp[i][j]는 a의 처음 i글자로 b의 처음 j글자를
//	만들 수 있는지 여부이다. 대문자는 반드시 매칭, 소문자는 변환 또는 삭제.
func abbreviation(a, b string) string {
	la, lb := len(a), len(b)

	// dp[i][j] = a[:i]로 b[:j]를 만들 수 있는지
	dp := make([][]bool, la+1)
	for i := range dp {
		dp[i] = make([]bool, lb+1)
	}

	// 빈 문자열로 빈 문자열을 만들 수 있음
	dp[0][0] = true

	// a의 앞부분이 모두 소문자이면 삭제하여 빈 문자열을 만들 수 있음
	for i := 1; i <= la; i++ {
		if unicode.IsLower(rune(a[i-1])) {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			if unicode.IsUpper(rune(a[i-1])) {
				// 대문자: 반드시 매칭해야 함
				if a[i-1] == b[j-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				// 소문자: 대문자로 변환하여 매칭하거나 삭제
				if unicode.ToUpper(rune(a[i-1])) == rune(b[j-1]) {
					// 변환하여 매칭 또는 삭제
					dp[i][j] = dp[i-1][j-1] || dp[i-1][j]
				} else {
					// 매칭 불가, 삭제만 가능
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	if dp[la][lb] {
		return "YES"
	}
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
