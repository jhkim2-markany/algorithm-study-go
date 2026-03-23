package main

import "fmt"

// 분할 정복 거듭제곱 - 모듈러 거듭제곱과 행렬 거듭제곱 예시
// 시간 복잡도: O(log n) (정수), O(k³ log n) (행렬)
// 공간 복잡도: O(1) (정수), O(k²) (행렬)

const MOD = 1000000007

// 모듈러 거듭제곱: base^exp mod MOD를 O(log exp)에 계산한다
func modPow(base, exp, mod int64) int64 {
	// 결과값을 1로 초기화한다
	result := int64(1)
	base %= mod

	// 지수가 0보다 큰 동안 반복한다
	for exp > 0 {
		// 지수가 홀수이면 결과에 밑을 곱한다
		if exp%2 == 1 {
			result = result * base % mod
		}
		// 밑을 제곱하고 지수를 반으로 나눈다
		base = base * base % mod
		exp /= 2
	}
	return result
}

// 2x2 행렬 타입 정의
type Matrix [2][2]int64

// 2x2 행렬 곱셈: 두 행렬을 곱하고 mod를 취한다
func matMul(a, b Matrix, mod int64) Matrix {
	var c Matrix
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][j] = (c[i][j] + a[i][k]*b[k][j]) % mod
			}
		}
	}
	return c
}

// 행렬 거듭제곱: 행렬 base를 exp번 거듭제곱한다
func matPow(base Matrix, exp int64, mod int64) Matrix {
	// 단위 행렬로 초기화한다
	result := Matrix{{1, 0}, {0, 1}}

	for exp > 0 {
		// 지수가 홀수이면 결과에 행렬을 곱한다
		if exp%2 == 1 {
			result = matMul(result, base, mod)
		}
		// 행렬을 제곱하고 지수를 반으로 나눈다
		base = matMul(base, base, mod)
		exp /= 2
	}
	return result
}

// 행렬 거듭제곱으로 피보나치 수열의 n번째 항을 구한다
func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	// 피보나치 전이 행렬: {{1,1},{1,0}}
	base := Matrix{{1, 1}, {1, 0}}
	result := matPow(base, n-1, MOD)
	// result[0][0]이 F(n)이다
	return result[0][0]
}

func main() {
	// 모듈러 거듭제곱 예시
	fmt.Println("=== 모듈러 거듭제곱 ===")
	fmt.Printf("2^10 mod %d = %d\n", MOD, modPow(2, 10, MOD))
	fmt.Printf("3^20 mod %d = %d\n", MOD, modPow(3, 20, MOD))
	fmt.Printf("7^1000000 mod %d = %d\n", MOD, modPow(7, 1000000, MOD))

	// 행렬 거듭제곱을 이용한 피보나치 수열 예시
	fmt.Println("\n=== 행렬 거듭제곱 피보나치 ===")
	for _, n := range []int64{5, 10, 50, 100} {
		fmt.Printf("F(%d) mod %d = %d\n", n, MOD, fibonacci(n))
	}
}
