package main

import "fmt"

// 순열 (Permutation) - 팩토리얼을 이용한 순열 계산
// n개에서 r개를 순서를 고려하여 선택하는 경우의 수를 구한다.
// 시간 복잡도: O(N) 전처리, O(1) 쿼리
// 공간 복잡도: O(N)

const maxN = 1000001
const mod = 1000000007

// 팩토리얼 배열과 역팩토리얼 배열
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
func precompute(n int) {
	fact[0] = 1
	for i := 1; i <= n; i++ {
		// i! = (i-1)! × i
		fact[i] = fact[i-1] * int64(i) % mod
	}
	// 페르마의 소정리로 역원을 구한다: (n!)⁻¹ ≡ (n!)^(p-2) mod p
	invFact[n] = power(fact[n], mod-2, mod)
	for i := n - 1; i >= 0; i-- {
		// (i!)⁻¹ = ((i+1)!)⁻¹ × (i+1)
		invFact[i] = invFact[i+1] * int64(i+1) % mod
	}
}

// nPr mod p 를 계산한다
func perm(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	// nPr = n! / (n-r)! = n! × ((n-r)!)⁻¹
	return fact[n] % mod * invFact[n-r] % mod
}

// nCr mod p 를 계산한다
func comb(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	// nCr = n! / (r! × (n-r)!) = n! × (r!)⁻¹ × ((n-r)!)⁻¹
	return fact[n] % mod * invFact[r] % mod * invFact[n-r] % mod
}

func main() {
	// 전처리
	precompute(100000)

	// 예시 1: 5P2 = 20
	fmt.Printf("P(5, 2) = %d\n", perm(5, 2))

	// 예시 2: 10P3 = 720
	fmt.Printf("P(10, 3) = %d\n", perm(10, 3))

	// 예시 3: 큰 수의 조합 (모듈러 연산)
	fmt.Printf("C(100000, 50000) mod 10^9+7 = %d\n", comb(100000, 50000))

	// 예시 4: 5C3 = 10
	fmt.Printf("C(5, 3) = %d\n", comb(5, 3))
}
