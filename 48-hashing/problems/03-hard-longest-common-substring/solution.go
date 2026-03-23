package main

import (
	"bufio"
	"fmt"
	"os"
)

// 이분 탐색 + 해싱으로 가장 긴 공통 부분 문자열의 길이를 구한다
// 시간 복잡도: O((N+M) log min(N,M))

const (
	base1 = 31
	mod1  = 1000000007
	base2 = 37
	mod2  = 1000000009
)

// hashPair는 이중 해시 값을 저장하는 구조체이다
type hashPair struct {
	h1, h2 int64
}

// buildPrefixHash는 문자열의 접두사 해시와 거듭제곱 배열을 전처리한다
func buildPrefixHash(s string) ([]int64, []int64, []int64, []int64) {
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

// getSubHash는 부분 문자열 s[l..l+length-1]의 이중 해시를 O(1)에 반환한다
func getSubHash(h1, h2, p1, p2 []int64, l, length int) hashPair {
	r := l + length
	sh1 := (h1[r] - h1[l]%mod1*p1[length]%mod1 + mod1*2) % mod1
	sh2 := (h2[r] - h2[l]%mod2*p2[length]%mod2 + mod2*2) % mod2
	return hashPair{sh1, sh2}
}

// hasCommonSubstring는 길이 length인 공통 부분 문자열이 존재하는지 확인한다
func hasCommonSubstring(a, b string, length int,
	ah1, ah2, ap1, ap2, bh1, bh2, bp1, bp2 []int64) bool {

	if length == 0 {
		return true
	}

	// 문자열 A의 길이 length인 모든 부분 문자열 해시를 집합에 저장한다
	hashSet := make(map[hashPair]bool)
	for i := 0; i+length <= len(a); i++ {
		h := getSubHash(ah1, ah2, ap1, ap2, i, length)
		hashSet[h] = true
	}

	// 문자열 B의 길이 length인 부분 문자열 해시가 집합에 있는지 확인한다
	for i := 0; i+length <= len(b); i++ {
		h := getSubHash(bh1, bh2, bp1, bp2, i, length)
		if hashSet[h] {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 두 문자열
	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)

	// 접두사 해시 전처리
	ah1, ah2, ap1, ap2 := buildPrefixHash(a)
	bh1, bh2, bp1, bp2 := buildPrefixHash(b)

	// 이분 탐색으로 가장 긴 공통 부분 문자열의 길이를 구한다
	// 길이 mid인 공통 부분 문자열이 존재하면 더 긴 것을 시도한다
	lo, hi := 0, len(a)
	if len(b) < hi {
		hi = len(b)
	}

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if hasCommonSubstring(a, b, mid, ah1, ah2, ap1, ap2, bh1, bh2, bp1, bp2) {
			lo = mid // 길이 mid가 가능하면 더 긴 것을 시도한다
		} else {
			hi = mid - 1 // 불가능하면 더 짧은 것을 시도한다
		}
	}

	fmt.Fprintln(writer, lo)
}
