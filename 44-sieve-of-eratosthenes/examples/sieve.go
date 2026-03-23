package main

import (
	"fmt"
	"math"
)

// 에라토스테네스의 체 (Sieve of Eratosthenes)
// 기본 체, 최적화 체, 세그먼트 체, 소인수분해 전처리를 구현한다.

// 기본 에라토스테네스의 체
// N 이하의 모든 소수를 구한다.
// 시간 복잡도: O(N log log N)
// 공간 복잡도: O(N)
func basicSieve(n int) []bool {
	// 소수 여부를 저장하는 배열
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// 2부터 √N까지 순회하며 배수를 제거한다
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// i의 배수를 i*i부터 제거한다 (i*i 미만은 이미 처리됨)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

// 최적화된 에라토스테네스의 체 (짝수 제외)
// 2를 별도 처리하고 홀수만 검사하여 메모리와 시간을 절약한다.
// 시간 복잡도: O(N log log N)
// 공간 복잡도: O(N/2)
func optimizedSieve(n int) []int {
	// 결과 소수 목록
	primes := []int{}
	if n < 2 {
		return primes
	}
	primes = append(primes, 2)

	// 홀수만 저장하는 배열 (인덱스 i는 수 2*i+1에 대응)
	size := (n - 1) / 2
	isComposite := make([]bool, size+1)

	// 3부터 홀수만 순회한다
	for i := 0; i <= size; i++ {
		if !isComposite[i] {
			// 현재 수: 2*i + 3 (i=0이면 3, i=1이면 5, ...)
			// 단, 여기서는 인덱스 i가 수 2*i+3에 대응하도록 조정
			p := 2*i + 3
			if p > n {
				break
			}
			primes = append(primes, p)

			// p의 홀수 배수를 제거한다 (p*p부터 시작)
			if int64(p)*int64(p) <= int64(n) {
				// p*p에 대응하는 인덱스부터 시작
				start := (p*p - 3) / 2
				for j := start; j <= size; j += p {
					isComposite[j] = true
				}
			}
		}
	}
	return primes
}

// 소인수분해 전처리 (Smallest Prime Factor)
// 각 수의 최소 소인수를 기록하여 빠른 소인수분해를 가능하게 한다.
// 시간 복잡도: O(N log log N) 전처리, O(log N) 쿼리
// 공간 복잡도: O(N)
func buildSPF(n int) []int {
	// spf[i]: i의 최소 소인수
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		spf[i] = i // 초기값은 자기 자신
	}

	// 에라토스테네스의 체와 유사하게 최소 소인수를 기록한다
	for i := 2; i*i <= n; i++ {
		if spf[i] == i { // i가 소수인 경우
			for j := i * i; j <= n; j += i {
				if spf[j] == j { // 아직 갱신되지 않은 경우만
					spf[j] = i
				}
			}
		}
	}
	return spf
}

// SPF 배열을 이용한 소인수분해
// 수 x를 소인수분해하여 (소인수, 지수) 쌍의 목록을 반환한다.
func factorize(x int, spf []int) [][2]int {
	factors := [][2]int{}
	for x > 1 {
		p := spf[x]
		cnt := 0
		// 같은 소인수를 모두 나눈다
		for x%p == 0 {
			x /= p
			cnt++
		}
		factors = append(factors, [2]int{p, cnt})
	}
	return factors
}

// 세그먼트 체 (Segmented Sieve)
// [L, R] 구간의 소수를 구한다.
// 시간 복잡도: O((R-L+1) log log R + √R)
// 공간 복잡도: O(√R + (R-L+1))
func segmentedSieve(l, r int64) []int64 {
	// 1단계: √R까지의 소수를 기본 체로 구한다
	limit := int64(math.Sqrt(float64(r))) + 1
	smallPrimes := []int64{}
	isPrime := make([]bool, limit+1)
	for i := int64(2); i <= limit; i++ {
		isPrime[i] = true
	}
	for i := int64(2); i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := int64(2); i <= limit; i++ {
		if isPrime[i] {
			smallPrimes = append(smallPrimes, i)
		}
	}

	// 2단계: [L, R] 구간에서 소수를 걸러낸다
	size := r - l + 1
	seg := make([]bool, size)
	for i := range seg {
		seg[i] = true
	}

	// L이 1 이하인 경우 0과 1을 제외한다
	if l <= 1 {
		for i := int64(0); i <= 1-l && i < size; i++ {
			seg[i] = false
		}
	}

	// 각 소수의 배수를 구간에서 제거한다
	for _, p := range smallPrimes {
		// 구간 내 p의 배수 시작점을 구한다
		start := ((l + p - 1) / p) * p
		if start == p {
			start += p // p 자체는 소수이므로 제외하지 않는다
		}
		for j := start; j <= r; j += p {
			seg[j-l] = false
		}
	}

	// 결과 수집
	result := []int64{}
	for i := int64(0); i < size; i++ {
		if seg[i] {
			result = append(result, l+i)
		}
	}
	return result
}

func main() {
	// 1. 기본 에라토스테네스의 체
	fmt.Println("=== 기본 에라토스테네스의 체 (N=50) ===")
	isPrime := basicSieve(50)
	fmt.Print("소수: ")
	for i := 2; i <= 50; i++ {
		if isPrime[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// 2. 최적화된 체 (짝수 제외)
	fmt.Println("\n=== 최적화된 체 (N=50) ===")
	primes := optimizedSieve(50)
	fmt.Print("소수: ")
	for _, p := range primes {
		fmt.Printf("%d ", p)
	}
	fmt.Println()

	// 3. 소인수분해 전처리
	fmt.Println("\n=== 소인수분해 전처리 (SPF) ===")
	spf := buildSPF(100)
	testNums := []int{12, 30, 60, 97}
	for _, n := range testNums {
		factors := factorize(n, spf)
		fmt.Printf("%d = ", n)
		for i, f := range factors {
			if i > 0 {
				fmt.Print(" × ")
			}
			if f[1] == 1 {
				fmt.Printf("%d", f[0])
			} else {
				fmt.Printf("%d^%d", f[0], f[1])
			}
		}
		fmt.Println()
	}

	// 4. 세그먼트 체
	fmt.Println("\n=== 세그먼트 체 [90, 110] ===")
	segPrimes := segmentedSieve(90, 110)
	fmt.Print("소수: ")
	for _, p := range segPrimes {
		fmt.Printf("%d ", p)
	}
	fmt.Println()
}
