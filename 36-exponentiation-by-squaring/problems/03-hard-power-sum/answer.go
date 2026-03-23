package main

import (
	"bufio"
	"fmt"
	"os"
)

// Matrix는 2x2 행렬 타입이다.
type Matrix [2][2]int64

var mod int64

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

func matPow(base Matrix, exp int64) Matrix {
	result := Matrix{{1, 0}, {0, 1}}
	for exp > 0 {
		if exp%2 == 1 {
			result = matMul(result, base)
		}
		base = matMul(base, base)
		exp /= 2
	}
	return result
}

// powerSum은 S(N) = A^1 + A^2 + ... + A^N mod M을 계산한다.
//
// [매개변수]
//   - a: 밑 (1 이상)
//   - n: 지수 합의 상한 (1 이상)
//   - m: 모듈러 값 (2 이상)
//
// [반환값]
//   - int64: (A^1 + A^2 + ... + A^N) mod M
//
// [알고리즘 힌트]
//
//	행렬 거듭제곱으로 거듭제곱 합을 구한다.
//	전이 행렬: {{1, A}, {0, A}}, 초기 벡터: [0, 1].
//	n번 거듭제곱 후 result[0][1]이 S(n)이다.
//	시간복잡도: O(log N) (2x2 행렬 곱셈은 상수)
func powerSum(a, n, m int64) int64 {
	mod = m
	if n == 0 {
		return 0
	}

	amod := a % mod
	base := Matrix{
		{1, amod},
		{0, amod},
	}

	result := matPow(base, n)
	return result[0][1] % mod
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, n, m int64
	fmt.Fscan(reader, &a, &n, &m)

	fmt.Fprintln(writer, powerSum(a, n, m))
}
