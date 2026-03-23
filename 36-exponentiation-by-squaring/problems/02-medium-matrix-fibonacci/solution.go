package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

// 2x2 행렬 타입
type Matrix [2][2]int64

// 행렬 곱셈: 두 2x2 행렬을 곱하고 mod를 취한다
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

// 행렬 거듭제곱: 행렬 base를 exp번 거듭제곱한다
func matPow(base Matrix, exp int64) Matrix {
	// 단위 행렬로 초기화한다
	result := Matrix{{1, 0}, {0, 1}}

	// 지수를 반씩 줄이며 행렬을 거듭제곱한다
	for exp > 0 {
		// 지수가 홀수이면 결과에 행렬을 곱한다
		if exp%2 == 1 {
			result = matMul(result, base)
		}
		// 행렬을 제곱하고 지수를 반으로 나눈다
		base = matMul(base, base)
		exp /= 2
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N 입력
	var n int64
	fmt.Fscan(reader, &n)

	// F(0) = 0, F(1) = 1 기저 처리
	if n <= 1 {
		fmt.Fprintln(writer, n)
		return
	}

	// 피보나치 전이 행렬 {{1,1},{1,0}}을 (n-1)번 거듭제곱한다
	base := Matrix{{1, 1}, {1, 0}}
	result := matPow(base, n-1)

	// result[0][0]이 F(n)이다
	fmt.Fprintln(writer, result[0][0])
}
