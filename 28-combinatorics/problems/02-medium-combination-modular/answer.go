package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxN = 1000001
const mod = 1000000007

// 팩토리얼과 역팩토리얼 배열
var fact [maxN]int64
var invFact [maxN]int64

// power는 모듈러 거듭제곱 base^exp mod m 을 계산한다
func power(base, exp, m int64) int64 {
	result := int64(1)
	base %= m
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % m
		}
		exp /= 2
		base = base * base % m
	}
	return result
}

// precompute는 팩토리얼과 역팩토리얼을 전처리한다
func precompute() {
	fact[0] = 1
	for i := 1; i < maxN; i++ {
		fact[i] = fact[i-1] * int64(i) % mod
	}
	invFact[maxN-1] = power(fact[maxN-1], mod-2, mod)
	for i := maxN - 2; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % mod
	}
}

// combMod는 nCr mod p 를 계산한다.
//
// [매개변수]
//   - n: 전체 원소 수
//   - r: 선택할 원소 수
//
// [반환값]
//   - int64: C(n, r) mod 1000000007
//
// [알고리즘 힌트]
//
//	nCr = n! × (r!)⁻¹ × ((n-r)!)⁻¹ mod p
//	팩토리얼과 역팩토리얼을 전처리하여 O(1)로 계산한다.
//	역팩토리얼은 페르마의 소정리로 구한다: a⁻¹ ≡ a^(p-2) mod p
func combMod(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	return fact[n] % mod * invFact[r] % mod * invFact[n-r] % mod
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 팩토리얼 전처리
	precompute()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)

		// 핵심 함수 호출
		fmt.Fprintln(writer, combMod(n, r))
	}
}
