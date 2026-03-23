package main

import (
	"bufio"
	"fmt"
	"os"
)

// longestCommonSubstringLen은 이분 탐색과 해싱을 이용하여
// 두 문자열의 가장 긴 공통 부분 문자열의 길이를 반환한다.
//
// [매개변수]
//   - a: 첫 번째 문자열
//   - b: 두 번째 문자열
//
// [반환값]
//   - int: 가장 긴 공통 부분 문자열의 길이
//
// [알고리즘 힌트]
//
//	이분 탐색으로 길이를 결정하고, 이중 해싱으로 해당 길이의 공통 부분 문자열 존재 여부를 확인한다.
//	접두사 해시를 전처리하여 부분 문자열 해시를 O(1)에 구한다.
//	시간 복잡도: O((N+M) log min(N,M))
func longestCommonSubstringLen(a, b string) int {
	const (
		base1 = 31
		mod1  = 1000000007
		base2 = 37
		mod2  = 1000000009
	)

	type hashPair struct {
		h1, h2 int64
	}

	buildHash := func(s string) ([]int64, []int64, []int64, []int64) {
		n := len(s)
		h1 := make([]int64, n+1)
		h2 := make([]int64, n+1)
		p1 := make([]int64, n+1)
		p2 := make([]int64, n+1)
		p1[0] = 1
		p2[0] = 1
		for i := 0; i < n; i++ {
			val := int64(s[i]-'a') + 1
			h1[i+1] = (h1[i] + val*p1[i]) % mod1
			h2[i+1] = (h2[i] + val*p2[i]) % mod2
			p1[i+1] = p1[i] * base1 % mod1
			p2[i+1] = p2[i] * base2 % mod2
		}
		return h1, h2, p1, p2
	}

	getHash := func(h1, h2, p1, p2 []int64, l, length int) hashPair {
		r := l + length
		sh1 := (h1[r] - h1[l]%mod1*p1[length]%mod1 + mod1*2) % mod1
		sh2 := (h2[r] - h2[l]%mod2*p2[length]%mod2 + mod2*2) % mod2
		return hashPair{sh1, sh2}
	}

	ah1, ah2, ap1, ap2 := buildHash(a)
	bh1, bh2, bp1, bp2 := buildHash(b)

	hasCommon := func(length int) bool {
		if length == 0 {
			return true
		}
		hashSet := make(map[hashPair]bool)
		for i := 0; i+length <= len(a); i++ {
			h := getHash(ah1, ah2, ap1, ap2, i, length)
			hashSet[h] = true
		}
		for i := 0; i+length <= len(b); i++ {
			h := getHash(bh1, bh2, bp1, bp2, i, length)
			if hashSet[h] {
				return true
			}
		}
		return false
	}

	lo, hi := 0, len(a)
	if len(b) < hi {
		hi = len(b)
	}
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if hasCommon(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	fmt.Fprintln(writer, longestCommonSubstringLen(a, b))
}
