package main

import "fmt"

// 문자열 해싱 (Polynomial Hashing)
// 다항식 해싱, 롤링 해시, 라빈-카프 알고리즘의 기본 구현

const (
	base1 = 31         // 첫 번째 해시 기저
	mod1  = 1000000007 // 첫 번째 모듈러 (10^9 + 7)
	base2 = 37         // 두 번째 해시 기저 (이중 해싱용)
	mod2  = 1000000009 // 두 번째 모듈러 (10^9 + 9)
)

// polyHash는 문자열의 다항식 해시 값을 계산한다
// hash(S) = s0*p^0 + s1*p^1 + ... + sn-1*p^(n-1) mod M
func polyHash(s string, base, mod int64) int64 {
	var hash int64
	var power int64 = 1
	for _, ch := range s {
		// 문자를 1부터 시작하는 정수로 변환한다 ('a' = 1)
		val := int64(ch-'a') + 1
		hash = (hash + val*power) % mod
		power = power * base % mod
	}
	return hash
}

// PrefixHash는 접두사 해시 배열을 관리하는 구조체이다
type PrefixHash struct {
	h1, h2 []int64 // 두 해시 함수의 접두사 해시 배열
	p1, p2 []int64 // 거듭제곱 배열
}

// NewPrefixHash는 문자열의 접두사 해시를 전처리한다
// 시간 복잡도: O(N)
func NewPrefixHash(s string) *PrefixHash {
	n := len(s)
	ph := &PrefixHash{
		h1: make([]int64, n+1),
		h2: make([]int64, n+1),
		p1: make([]int64, n+1),
		p2: make([]int64, n+1),
	}

	ph.p1[0] = 1
	ph.p2[0] = 1

	// 접두사 해시와 거듭제곱을 동시에 전처리한다
	for i := 0; i < n; i++ {
		val := int64(s[i]-'a') + 1
		ph.h1[i+1] = (ph.h1[i] + val*ph.p1[i]) % mod1
		ph.h2[i+1] = (ph.h2[i] + val*ph.p2[i]) % mod2
		ph.p1[i+1] = ph.p1[i] * base1 % mod1
		ph.p2[i+1] = ph.p2[i] * base2 % mod2
	}

	return ph
}

// SubHash는 부분 문자열 s[l..r] (0-indexed, 양 끝 포함)의 이중 해시를 O(1)에 반환한다
func (ph *PrefixHash) SubHash(l, r int) (int64, int64) {
	length := r - l + 1
	// 첫 번째 해시: (H[r+1] - H[l] * P[length]) mod M
	h1 := (ph.h1[r+1] - ph.h1[l]%mod1*ph.p1[length]%mod1 + mod1*2) % mod1
	h2 := (ph.h2[r+1] - ph.h2[l]%mod2*ph.p2[length]%mod2 + mod2*2) % mod2
	return h1, h2
}

// rabinKarp는 라빈-카프 알고리즘으로 텍스트에서 패턴의 모든 출현 위치를 찾는다
// 시간 복잡도: 평균 O(N+M), 최악 O(NM)
func rabinKarp(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	if m > n {
		return nil
	}

	var result []int

	// 패턴의 해시 값을 계산한다
	patHash := polyHash(pattern, base1, mod1)

	// 텍스트의 첫 번째 윈도우 해시를 계산한다
	var winHash int64
	var power int64 = 1
	for i := 0; i < m; i++ {
		val := int64(text[i]-'a') + 1
		winHash = (winHash + val*power) % mod1
		if i < m-1 {
			power = power * base1 % mod1
		}
	}

	// 첫 위치 확인
	if winHash == patHash && text[:m] == pattern {
		result = append(result, 0)
	}

	// 롤링 해시로 나머지 위치를 확인한다
	for i := 1; i+m-1 < n; i++ {
		// 이전 문자를 제거하고 새 문자를 추가한다
		oldVal := int64(text[i-1]-'a') + 1
		newVal := int64(text[i+m-1]-'a') + 1
		winHash = (winHash - oldVal + mod1) % mod1
		winHash = winHash * modInverse(base1, mod1) % mod1
		winHash = (winHash + newVal*power) % mod1

		// 해시가 일치하면 실제 문자열을 비교한다 (충돌 방지)
		if winHash == patHash && text[i:i+m] == pattern {
			result = append(result, i)
		}
	}

	return result
}

// modInverse는 페르마 소정리를 이용한 모듈러 역원을 구한다
func modInverse(a, m int64) int64 {
	return modPow(a, m-2, m)
}

// modPow는 빠른 거듭제곱으로 a^b mod m을 계산한다
func modPow(a, b, m int64) int64 {
	result := int64(1)
	a = a % m
	for b > 0 {
		if b%2 == 1 {
			result = result * a % m
		}
		b /= 2
		a = a * a % m
	}
	return result
}

func main() {
	// 1. 다항식 해싱 기본 예제
	fmt.Println("=== 다항식 해싱 ===")
	words := []string{"abc", "abd", "abc", "xyz"}
	for _, w := range words {
		h := polyHash(w, base1, mod1)
		fmt.Printf("hash(\"%s\") = %d\n", w, h)
	}
	fmt.Println("같은 문자열은 같은 해시 값을 가진다")

	// 2. 접두사 해시를 이용한 부분 문자열 비교
	fmt.Println("\n=== 접두사 해시 (부분 문자열 비교) ===")
	s := "abcabcabc"
	ph := NewPrefixHash(s)
	fmt.Printf("문자열: \"%s\"\n", s)

	// s[0..2] = "abc"와 s[3..5] = "abc" 비교
	h1a, h2a := ph.SubHash(0, 2)
	h1b, h2b := ph.SubHash(3, 5)
	fmt.Printf("s[0..2] 해시: (%d, %d)\n", h1a, h2a)
	fmt.Printf("s[3..5] 해시: (%d, %d)\n", h1b, h2b)
	if h1a == h1b && h2a == h2b {
		fmt.Println("s[0..2]와 s[3..5]는 같은 문자열이다")
	}

	// s[0..2] = "abc"와 s[1..3] = "bca" 비교
	h1c, h2c := ph.SubHash(1, 3)
	fmt.Printf("s[1..3] 해시: (%d, %d)\n", h1c, h2c)
	if h1a != h1c || h2a != h2c {
		fmt.Println("s[0..2]와 s[1..3]는 다른 문자열이다")
	}

	// 3. 라빈-카프 패턴 매칭
	fmt.Println("\n=== 라빈-카프 패턴 매칭 ===")
	text := "ababababab"
	pattern := "abab"
	positions := rabinKarp(text, pattern)
	fmt.Printf("텍스트: \"%s\"\n", text)
	fmt.Printf("패턴: \"%s\"\n", pattern)
	fmt.Printf("출현 위치: %v\n", positions)

	// 4. 이중 해싱 예제
	fmt.Println("\n=== 이중 해싱 ===")
	testWords := []string{"hello", "world", "hello"}
	for _, w := range testWords {
		h1 := polyHash(w, base1, mod1)
		h2 := polyHash(w, base2, mod2)
		fmt.Printf("hash(\"%s\") = (%d, %d)\n", w, h1, h2)
	}
	fmt.Println("이중 해싱은 충돌 확률을 크게 줄인다")
}
