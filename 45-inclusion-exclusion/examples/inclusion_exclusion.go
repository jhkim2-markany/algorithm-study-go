package main

import (
	"fmt"
	"math/bits"
)

// 포함 배제의 원리 (Inclusion-Exclusion Principle)
// 비트마스크를 이용한 합집합 크기 계산, 배수 문제, 오일러 피 함수를 구현한다.

// 비트마스크를 이용한 포함 배제
// N 이하에서 주어진 수들의 배수인 수의 개수를 구한다.
// 시간 복잡도: O(2^K) (K는 주어진 수의 개수)
// 공간 복잡도: O(K)
func countMultiples(n int, divisors []int) int {
	k := len(divisors)
	result := 0

	// 비트마스크로 모든 비어있지 않은 부분집합을 열거한다
	for mask := 1; mask < (1 << k); mask++ {
		// 선택된 수들의 최소공배수(여기서는 곱)를 계산한다
		product := 1
		overflow := false
		for i := 0; i < k; i++ {
			if mask&(1<<i) != 0 {
				product *= divisors[i]
				// 곱이 N을 초과하면 해당 배수는 0개이다
				if product > n {
					overflow = true
					break
				}
			}
		}

		if overflow {
			continue
		}

		// 켜진 비트 수가 홀수이면 더하고, 짝수이면 뺀다
		count := n / product
		if bits.OnesCount(uint(mask))%2 == 1 {
			result += count
		} else {
			result -= count
		}
	}
	return result
}

// 오일러 피 함수 (Euler's Totient Function)
// 1부터 n까지의 수 중 n과 서로소인 수의 개수를 구한다.
// 포함 배제의 원리를 이용하여 계산한다.
// 시간 복잡도: O(√n)
// 공간 복잡도: O(1)
func eulerPhi(n int) int {
	result := n

	// n의 소인수를 구하며 포함 배제를 적용한다
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			// p는 n의 소인수이다
			// φ(n) = n × (1 - 1/p) = n - n/p
			result -= result / p
			// n에서 p를 모두 나눈다
			for n%p == 0 {
				n /= p
			}
		}
	}

	// n이 1보다 크면 남은 소인수가 하나 있다
	if n > 1 {
		result -= result / n
	}
	return result
}

// 비트마스크 포함 배제로 서로소인 수의 개수를 구한다
// 1부터 n까지의 수 중 n과 서로소인 수의 개수를 반환한다.
// 시간 복잡도: O(2^K + √n) (K는 소인수의 개수)
func countCoprime(n int) int {
	// n의 소인수를 구한다
	primeFactors := []int{}
	temp := n
	for p := 2; p*p <= temp; p++ {
		if temp%p == 0 {
			primeFactors = append(primeFactors, p)
			for temp%p == 0 {
				temp /= p
			}
		}
	}
	if temp > 1 {
		primeFactors = append(primeFactors, temp)
	}

	// 포함 배제로 소인수의 배수인 수의 개수를 구한다
	multiples := countMultiples(n, primeFactors)

	// 서로소인 수의 개수 = 전체 - 배수인 수의 개수
	return n - multiples
}

// 교란 순열 (Derangement) 개수
// n개의 원소에서 어떤 원소도 원래 위치에 있지 않은 순열의 수를 구한다.
// 포함 배제의 원리를 적용한 공식: D(n) = n! × Σ(k=0 to n) (-1)^k / k!
// 시간 복잡도: O(n)
// 공간 복잡도: O(1)
func derangement(n int) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}

	// D(n) = (n-1) × (D(n-1) + D(n-2)) 점화식을 사용한다
	prev2 := int64(1) // D(0) = 1
	prev1 := int64(0) // D(1) = 0

	for i := 2; i <= n; i++ {
		current := int64(i-1) * (prev1 + prev2)
		prev2 = prev1
		prev1 = current
	}
	return prev1
}

func main() {
	// 1. 비트마스크 포함 배제: N 이하에서 배수의 개수
	fmt.Println("=== 비트마스크 포함 배제 ===")
	n := 30
	divisors := []int{2, 3, 5}
	multCount := countMultiples(n, divisors)
	fmt.Printf("%d 이하에서 %v 중 하나 이상의 배수인 수의 개수: %d\n", n, divisors, multCount)
	fmt.Printf("%d 이하에서 %v 모두와 서로소인 수의 개수: %d\n", n, divisors, n-multCount)

	// 2. 오일러 피 함수
	fmt.Println("\n=== 오일러 피 함수 ===")
	testNums := []int{1, 6, 12, 30, 97}
	for _, num := range testNums {
		fmt.Printf("φ(%d) = %d\n", num, eulerPhi(num))
	}

	// 3. 비트마스크 포함 배제로 서로소 개수 검증
	fmt.Println("\n=== 서로소 개수 (비트마스크 포함 배제) ===")
	for _, num := range testNums {
		fmt.Printf("1~%d 중 %d와 서로소인 수의 개수: %d\n", num, num, countCoprime(num))
	}

	// 4. 교란 순열
	fmt.Println("\n=== 교란 순열 ===")
	for i := 0; i <= 10; i++ {
		fmt.Printf("D(%d) = %d\n", i, derangement(i))
	}
}
