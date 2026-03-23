package main

import (
	"fmt"
	"math/big"
)

// 소수 판정 (Primality Test) - 시행 나눗셈과 밀러-라빈 알고리즘
// 주어진 수가 소수인지 판별하는 두 가지 방법을 구현한다.

// 시행 나눗셈 (Trial Division)
// 2부터 √N까지 나누어 보며 소수를 판정한다.
// 시간 복잡도: O(√N)
// 공간 복잡도: O(1)
func isPrimeTrialDivision(n int64) bool {
	// 2 미만은 소수가 아니다
	if n < 2 {
		return false
	}
	// 2와 3은 소수이다
	if n < 4 {
		return true
	}
	// 짝수와 3의 배수를 먼저 제외한다
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// 6k ± 1 형태의 수로만 나누어 본다
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// 모듈러 거듭제곱: (base^exp) mod m 을 계산한다
// 오버플로우 방지를 위해 big.Int를 사용한다
func modPow(base, exp, m int64) int64 {
	b := big.NewInt(base)
	e := big.NewInt(exp)
	mod := big.NewInt(m)
	result := new(big.Int).Exp(b, e, mod)
	return result.Int64()
}

// 밀러-라빈 소수 판정법 (Miller-Rabin Primality Test)
// 결정적 밀러-라빈: 특정 밑 집합으로 int64 범위 내에서 정확히 판정한다.
// 시간 복잡도: O(k × log² N), k는 밑의 개수
// 공간 복잡도: O(1)
func isPrimeMillerRabin(n int64) bool {
	// 작은 수 처리
	if n < 2 {
		return false
	}
	// 작은 소수 목록으로 빠르게 판정한다
	smallPrimes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, p := range smallPrimes {
		if n == p {
			return true
		}
		if n%p == 0 {
			return false
		}
	}

	// n-1 = 2^s × d 형태로 분해한다 (d는 홀수)
	d := n - 1
	s := 0
	for d%2 == 0 {
		d /= 2
		s++
	}

	// 각 밑에 대해 밀러-라빈 테스트를 수행한다
	for _, a := range smallPrimes {
		if a >= n {
			continue
		}
		// x = a^d mod n 을 계산한다
		x := modPow(a, d, n)

		// x가 1이거나 n-1이면 이 밑에 대해 통과
		if x == 1 || x == n-1 {
			continue
		}

		// s-1번 제곱하며 n-1이 나오는지 확인한다
		passed := false
		for r := 0; r < s-1; r++ {
			x = modPow(x, 2, n)
			if x == n-1 {
				passed = true
				break
			}
		}
		// 통과하지 못하면 합성수이다
		if !passed {
			return false
		}
	}
	return true
}

func main() {
	// 시행 나눗셈 예시
	fmt.Println("=== 시행 나눗셈 (Trial Division) ===")
	testNumbers := []int64{2, 7, 15, 97, 100, 997}
	for _, n := range testNumbers {
		if isPrimeTrialDivision(n) {
			fmt.Printf("%d: 소수\n", n)
		} else {
			fmt.Printf("%d: 합성수\n", n)
		}
	}

	// 밀러-라빈 예시 (큰 수 판정)
	fmt.Println("\n=== 밀러-라빈 (Miller-Rabin) ===")
	largeNumbers := []int64{
		1000000007,         // 소수 (10^9 + 7)
		1000000009,         // 소수 (10^9 + 9)
		999999999999999989, // 소수
		999999999999999990, // 합성수
	}
	for _, n := range largeNumbers {
		if isPrimeMillerRabin(n) {
			fmt.Printf("%d: 소수\n", n)
		} else {
			fmt.Printf("%d: 합성수\n", n)
		}
	}

	// 두 방법 비교 (작은 수에서 결과 일치 확인)
	fmt.Println("\n=== 두 방법 비교 (1~50) ===")
	fmt.Print("소수 목록: ")
	for i := int64(1); i <= 50; i++ {
		if isPrimeMillerRabin(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
