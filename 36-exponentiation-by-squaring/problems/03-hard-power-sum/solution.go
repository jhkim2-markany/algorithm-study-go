package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2x2 행렬 타입
type Matrix [2][2]int64

var mod int64

// 행렬 곱셈: 두 2x2 행렬을 곱하고 mod를 취한다
func matMul(a, b Matrix) Matrix {
	var c Matrix
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][j] = (c[i][j] + a[i][k]%mod*(b[k][j]%mod)) % mod
			}
		}
	}
	return c
}

// 행렬 거듭제곱: 행렬 base를 exp번 거듭제곱한다
func matPow(base Matrix, exp int64) Matrix {
	// 단위 행렬로 초기화한다
	result := Matrix{{1, 0}, {0, 1}}

	// 지수를 반씩 줄이며 행렬을 거듭제곱한다
	for exp > 0 {
		if exp%2 == 1 {
			result = matMul(result, base)
		}
		base = matMul(base, base)
		exp /= 2
	}
	return result
}

// 거듭제곱 합 S(N) = A^1 + A^2 + ... + A^N을 행렬 거듭제곱으로 구한다
// 점화식: S(n) = S(n-1) + A^n, A^n = A * A^(n-1)
// | S(n) |   | 1  A |   | S(n-1)   |
// | A^n  | = | 0  A | × | A^(n-1)  |
// 초기값: [S(0), A^0] = [0, 1]
func powerSum(a, n int64) int64 {
	if n == 0 {
		return 0
	}

	// 전이 행렬 구성
	amod := a % mod
	base := Matrix{
		{1, amod},
		{0, amod},
	}

	// 행렬을 n번 거듭제곱한다
	result := matPow(base, n)

	// result × [0, 1]^T 에서 result[0][1]이 S(n)이다
	return result[0][1] % mod
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// A, N, M 입력
	var a, n int64
	fmt.Fscan(reader, &a, &n, &mod)

	// 거듭제곱 합 계산 후 출력
	fmt.Fprintln(writer, powerSum(a, n))
}
