package main

import (
	"bufio"
	"fmt"
	"os"
)

// countEqualPairs는 이중 다항식 해싱을 이용하여 동일한 문자열 쌍의 개수를 반환한다.
//
// [매개변수]
//   - strs: 문자열 배열
//
// [반환값]
//   - int64: 동일한 문자열 쌍의 개수 (k개의 같은 문자열이면 k*(k-1)/2)
//
// [알고리즘 힌트]
//
//	이중 해싱(base1=31/mod1=10^9+7, base2=37/mod2=10^9+9)으로 충돌 확률을 줄인다.
//	같은 해시 쌍을 가진 문자열이 k개이면 쌍의 수는 k*(k-1)/2이다.
func countEqualPairs(strs []string) int64 {
	const (
		base1 = 31
		mod1  = 1000000007
		base2 = 37
		mod2  = 1000000009
	)

	type hashPair struct {
		h1, h2 int64
	}
	count := make(map[hashPair]int64)

	for _, s := range strs {
		var h1, h2 int64
		var p1, p2 int64 = 1, 1
		for _, ch := range s {
			val := int64(ch-'a') + 1
			h1 = (h1 + val*p1) % mod1
			h2 = (h2 + val*p2) % mod2
			p1 = p1 * base1 % mod1
			p2 = p2 * base2 % mod2
		}
		count[hashPair{h1, h2}]++
	}

	var result int64
	for _, k := range count {
		result += k * (k - 1) / 2
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &strs[i])
	}

	fmt.Fprintln(writer, countEqualPairs(strs))
}
