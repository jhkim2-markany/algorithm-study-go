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

// 모듈러 거듭제곱: base^exp mod m 을 계산한다
func power(base, exp, m int64) int64 {
	result := int64(1)
	base %= m
	for exp > 0 {
		// 지수가 홀수이면 결과에 밑을 곱한다
		if exp%2 == 1 {
			result = result * base % m
		}
		exp /= 2
		base = base * base % m
	}
	return result
}

// 팩토리얼과 역팩토리얼을 전처리한다
func precompute() {
	fact[0] = 1
	for i := 1; i < maxN; i++ {
		// i! = (i-1)! × i
		fact[i] = fact[i-1] * int64(i) % mod
	}
	// 페르마의 소정리로 (maxN-1)!의 역원을 구한다
	invFact[maxN-1] = power(fact[maxN-1], mod-2, mod)
	for i := maxN - 2; i >= 0; i-- {
		// (i!)⁻¹ = ((i+1)!)⁻¹ × (i+1)
		invFact[i] = invFact[i+1] * int64(i+1) % mod
	}
}

// nCr mod p 를 계산한다
func comb(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	// nCr = n! × (r!)⁻¹ × ((n-r)!)⁻¹ mod p
	return fact[n] % mod * invFact[r] % mod * invFact[n-r] % mod
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 팩토리얼 전처리
	precompute()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)

		// 전처리된 값으로 이항 계수를 계산하여 출력한다
		fmt.Fprintln(writer, comb(n, r))
	}
}
