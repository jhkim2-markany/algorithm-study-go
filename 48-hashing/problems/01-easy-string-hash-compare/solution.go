package main

import (
	"bufio"
	"fmt"
	"os"
)

// 다항식 해싱을 이용한 문자열 쌍 개수 세기
// 이중 해싱으로 충돌 확률을 줄인다

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

	// 입력: 문자열 개수
	var n int
	fmt.Fscan(reader, &n)

	// 이중 해시 값을 키로 사용하여 같은 문자열의 개수를 센다
	type hashPair struct {
		h1, h2 int64
	}
	count := make(map[hashPair]int64)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)

		// 다항식 해시 계산 (두 개의 해시 함수 사용)
		var h1, h2 int64
		var p1, p2 int64 = 1, 1
		for _, ch := range s {
			val := int64(ch-'a') + 1
			h1 = (h1 + val*p1) % mod1
			h2 = (h2 + val*p2) % mod2
			p1 = p1 * base1 % mod1
			p2 = p2 * base2 % mod2
		}

		key := hashPair{h1, h2}
		count[key]++
	}

	// 같은 해시를 가진 문자열이 k개이면 쌍의 수는 k*(k-1)/2이다
	var result int64
	for _, k := range count {
		result += k * (k - 1) / 2
	}

	fmt.Fprintln(writer, result)
}
