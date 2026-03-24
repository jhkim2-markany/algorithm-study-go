package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

// samAndSubstrings는 숫자 문자열의 모든 부분 문자열 합을 반환한다.
//
// [매개변수]
//   - s: 숫자 문자열
//
// [반환값]
//   - int: 모든 부분 문자열의 합 (mod 10^9+7)
//
// [알고리즘 힌트]
//
//	f(i) = i번째 자릿수로 끝나는 모든 부분 문자열의 합
//	f(i) = f(i-1)*10 + digit[i]*(i+1)
//	전체 합 = Σf(i)
func samAndSubstrings(s string) int {
	totalSum := 0 // 전체 부분 문자열 합
	fi := 0       // 현재 위치로 끝나는 부분 문자열들의 합

	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')

		// 이전 위치로 끝나는 부분 문자열에 현재 자릿수를 붙임 + 현재 자릿수 단독
		fi = (fi*10 + digit*(i+1)) % mod

		// 전체 합에 누적
		totalSum = (totalSum + fi) % mod
	}

	return totalSum
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
