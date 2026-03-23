package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

// Matrix는 2x2 행렬 타입이다.
type Matrix [2][2]int64

func matMul(a, b Matrix) Matrix {
	var c Matrix
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][j] = (c[i][j] + a[i][k]*b[k][j]) % MOD
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

// fibonacci는 N번째 피보나치 수를 행렬 거듭제곱으로 구한다.
//
// [매개변수]
//   - n: 구할 피보나치 수의 인덱스 (0 이상)
//
// [반환값]
//   - int64: F(n) mod 10^9+7
//
// [알고리즘 힌트]
//
//	피보나치 전이 행렬 {{1,1},{1,0}}을 (n-1)번 거듭제곱한다.
//	result[0][0]이 F(n)이다.
//	시간복잡도: O(log N) (2x2 행렬 곱셈은 상수)
func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	base := Matrix{{1, 1}, {1, 0}}
	result := matPow(base, n-1)
	return result[0][0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int64
	fmt.Fscan(reader, &n)

	fmt.Fprintln(writer, fibonacci(n))
}
