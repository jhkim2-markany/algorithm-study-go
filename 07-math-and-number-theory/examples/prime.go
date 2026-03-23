package main

import (
	"fmt"
	"math"
)

// 소수 판별 및 에라토스테네스의 체 예시
// 소수 판별 시간 복잡도: O(√N)
// 에라토스테네스의 체 시간 복잡도: O(N log log N)

// isPrime 함수는 주어진 수가 소수인지 판별한다
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	// 2부터 √N까지만 확인하면 충분하다
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// sieve 함수는 에라토스테네스의 체로 N 이하의 소수 목록을 반환한다
func sieve(n int) []int {
	// 모든 수를 소수 후보로 초기화한다
	isP := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isP[i] = true
	}

	// 2부터 √N까지 순회하며 배수를 제거한다
	for i := 2; i*i <= n; i++ {
		if isP[i] {
			// i의 배수를 모두 합성수로 표시한다
			for j := i * i; j <= n; j += i {
				isP[j] = false
			}
		}
	}

	// 소수만 모아서 반환한다
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isP[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// modPow 함수는 a^n mod m을 빠른 거듭제곱으로 계산한다
func modPow(a, n, m int) int {
	result := 1
	a = a % m
	// n의 이진 표현을 이용하여 O(log n)에 계산한다
	for n > 0 {
		// n이 홀수이면 결과에 a를 곱한다
		if n%2 == 1 {
			result = result * a % m
		}
		n /= 2
		a = a * a % m
	}
	return result
}

func main() {
	// 소수 판별 예시
	fmt.Println("=== 소수 판별 ===")
	testNums := []int{1, 2, 3, 4, 17, 18, 97, 100}
	for _, n := range testNums {
		if isPrime(n) {
			fmt.Printf("%d: 소수\n", n)
		} else {
			fmt.Printf("%d: 소수 아님\n", n)
		}
	}

	// 에라토스테네스의 체 예시
	fmt.Println("\n=== 에라토스테네스의 체 (50 이하 소수) ===")
	primes := sieve(50)
	fmt.Println(primes)
	fmt.Printf("50 이하 소수의 개수: %d\n", len(primes))

	// 빠른 거듭제곱 예시
	fmt.Println("\n=== 빠른 거듭제곱 (모듈러) ===")
	fmt.Printf("2^10 mod 1000 = %d\n", modPow(2, 10, 1000))
	fmt.Printf("3^13 mod 1000000007 = %d\n", modPow(3, 13, 1000000007))
}
