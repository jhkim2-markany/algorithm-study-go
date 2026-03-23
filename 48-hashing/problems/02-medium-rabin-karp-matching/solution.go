package main

import (
	"bufio"
	"fmt"
	"os"
)

// 라빈-카프 알고리즘을 이용한 문자열 패턴 매칭
// 롤링 해시로 O(N+M) 평균 시간에 패턴의 모든 출현 위치를 찾는다

const (
	base1 = 31
	mod1  = 1000000007
	base2 = 37
	mod2  = 1000000009
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 텍스트와 패턴
	var text, pattern string
	fmt.Fscan(reader, &text)
	fmt.Fscan(reader, &pattern)

	n := len(text)
	m := len(pattern)

	if m > n {
		fmt.Fprintln(writer, 0)
		return
	}

	// 패턴의 이중 해시 값을 계산한다
	var patH1, patH2 int64
	var pw1, pw2 int64 = 1, 1
	for i := 0; i < m; i++ {
		val := int64(pattern[i]-'a') + 1
		patH1 = (patH1 + val*pw1) % mod1
		patH2 = (patH2 + val*pw2) % mod2
		if i < m-1 {
			pw1 = pw1 * base1 % mod1
			pw2 = pw2 * base2 % mod2
		}
	}
	// pw1, pw2는 base^(m-1) mod M 값이다

	// 텍스트의 접두사 해시를 전처리한다
	h1 := make([]int64, n+1)
	h2 := make([]int64, n+1)
	p1 := make([]int64, n+1)
	p2 := make([]int64, n+1)
	p1[0] = 1
	p2[0] = 1
	for i := 0; i < n; i++ {
		val := int64(text[i]-'a') + 1
		h1[i+1] = (h1[i] + val*p1[i]) % mod1
		h2[i+1] = (h2[i] + val*p2[i]) % mod2
		p1[i+1] = p1[i] * base1 % mod1
		p2[i+1] = p2[i] * base2 % mod2
	}

	// 각 위치에서 부분 문자열 해시를 O(1)에 구하여 패턴과 비교한다
	var positions []int
	for i := 0; i+m-1 < n; i++ {
		// 부분 문자열 text[i..i+m-1]의 해시를 구한다
		subH1 := (h1[i+m] - h1[i]%mod1*p1[m]%mod1 + mod1*2) % mod1
		subH2 := (h2[i+m] - h2[i]%mod2*p2[m]%mod2 + mod2*2) % mod2

		// 이중 해시가 모두 일치하면 매칭으로 판정한다
		if subH1 == patH1 && subH2 == patH2 {
			positions = append(positions, i+1) // 1-indexed
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
