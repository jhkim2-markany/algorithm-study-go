package main

import "fmt"

// 모듈로 곱셈 역원 (Modular Multiplicative Inverse)
// 페르마 소정리, 확장 유클리드 알고리즘, 이항 계수 모듈러 계산의 기본 구현

const MOD = 1000000007 // 10^9 + 7 (소수)

// modPow는 빠른 거듭제곱으로 a^b mod m을 계산한다
// 시간 복잡도: O(log b)
func modPow(a, b, m int64) int64 {
	a %= m
	result := int64(1)
	for b > 0 {
		// b의 마지막 비트가 1이면 결과에 a를 곱한다
		if b%2 == 1 {
			result = result * a % m
		}
		b /= 2
		a = a * a % m
	}
	return result
}

// fermatInverse는 페르마 소정리를 이용하여 a의 모듈러 역원을 구한다
// M이 소수일 때: a^(-1) ≡ a^(M-2) (mod M)
// 시간 복잡도: O(log M)
func fermatInverse(a, m int64) int64 {
	return modPow(a, m-2, m)
}

// extGCD는 확장 유클리드 알고리즘으로 ax + by = gcd(a, b)의 해를 구한다
// 반환값: gcd, x, y
func extGCD(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := extGCD(b, a%b)
	// a%b = a - (a/b)*b 이므로 역추적한다
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}

// extEuclidInverse는 확장 유클리드 알고리즘으로 a의 모듈러 역원을 구한다
// gcd(a, M) = 1일 때 사용 가능 (M이 소수가 아니어도 됨)
// 시간 복잡도: O(log M)
func extEuclidInverse(a, m int64) int64 {
	_, x, _ := extGCD(a, m)
	// x가 음수일 수 있으므로 양수로 변환한다
	return (x%m + m) % m
}

// 팩토리얼과 역팩토리얼 전처리를 위한 전역 배열
var (
	fact    []int64 // fact[i] = i! mod M
	invFact []int64 // invFact[i] = (i!)^(-1) mod M
)

// initFactorial은 팩토리얼과 역팩토리얼 배열을 전처리한다
// 시간 복잡도: O(N + log M)
func initFactorial(n int, m int64) {
	fact = make([]int64, n+1)
	invFact = make([]int64, n+1)

	// 팩토리얼 계산
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i) % m
	}

	// 역팩토리얼 계산: fact[n]의 역원을 구한 뒤 역순으로 계산한다
	invFact[n] = fermatInverse(fact[n], m)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % m
	}
}

// nCr은 이항 계수 C(n, r) mod M을 O(1)에 반환한다
// initFactorial을 먼저 호출해야 한다
func nCr(n, r int, m int64) int64 {
	if r < 0 || r > n {
		return 0
	}
	return fact[n] % m * invFact[r] % m * invFact[n-r] % m
}

// precomputeInverse는 1부터 n까지의 모듈러 역원을 O(N)에 전처리한다
// 점화식: inv[i] = -(M/i) * inv[M%i] mod M
func precomputeInverse(n int, m int64) []int64 {
	inv := make([]int64, n+1)
	inv[1] = 1
	for i := int64(2); i <= int64(n); i++ {
		// M = (M/i)*i + (M%i) 에서 유도된 점화식
		inv[i] = (m - (m/i)*inv[m%i]%m) % m
	}
	return inv
}

func main() {
	// 1. 페르마 소정리를 이용한 역원 계산
	fmt.Println("=== 페르마 소정리 역원 ===")
	a := int64(3)
	inv := fermatInverse(a, MOD)
	fmt.Printf("%d의 역원 (mod %d) = %d\n", a, MOD, inv)
	fmt.Printf("검증: %d × %d mod %d = %d\n", a, inv, MOD, a*inv%MOD)

	// 2. 확장 유클리드 알고리즘을 이용한 역원 계산
	fmt.Println("\n=== 확장 유클리드 역원 ===")
	b := int64(5)
	inv2 := extEuclidInverse(b, MOD)
	fmt.Printf("%d의 역원 (mod %d) = %d\n", b, MOD, inv2)
	fmt.Printf("검증: %d × %d mod %d = %d\n", b, inv2, MOD, b*inv2%MOD)

	// M이 소수가 아닌 경우에도 확장 유클리드는 동작한다
	m2 := int64(12)
	c := int64(5) // gcd(5, 12) = 1이므로 역원 존재
	inv3 := extEuclidInverse(c, m2)
	fmt.Printf("%d의 역원 (mod %d) = %d\n", c, m2, inv3)
	fmt.Printf("검증: %d × %d mod %d = %d\n", c, inv3, m2, c*inv3%m2)

	// 3. 모듈러 나눗셈
	fmt.Println("\n=== 모듈러 나눗셈 ===")
	// 10 / 3 mod MOD 계산
	numerator := int64(10)
	denominator := int64(3)
	divResult := numerator % MOD * fermatInverse(denominator, MOD) % MOD
	fmt.Printf("%d / %d mod %d = %d\n", numerator, denominator, MOD, divResult)
	fmt.Printf("검증: %d × %d mod %d = %d (원래 분자)\n", divResult, denominator, MOD, divResult*denominator%MOD)

	// 4. 이항 계수 모듈러 계산
	fmt.Println("\n=== 이항 계수 (nCr mod p) ===")
	initFactorial(100, MOD)
	fmt.Printf("C(10, 3) mod %d = %d (실제값: 120)\n", MOD, nCr(10, 3, MOD))
	fmt.Printf("C(20, 10) mod %d = %d (실제값: 184756)\n", MOD, nCr(20, 10, MOD))
	fmt.Printf("C(100, 50) mod %d = %d\n", MOD, nCr(100, 50, MOD))

	// 5. 역원 배열 전처리
	fmt.Println("\n=== 역원 배열 전처리 ===")
	invArr := precomputeInverse(10, MOD)
	for i := 1; i <= 10; i++ {
		fmt.Printf("inv[%d] = %d (검증: %d × %d mod %d = %d)\n",
			i, invArr[i], int64(i), invArr[i], MOD, int64(i)*invArr[i]%MOD)
	}
}
