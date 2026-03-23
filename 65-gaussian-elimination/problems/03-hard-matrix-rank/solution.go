package main

import (
	"bufio"
	"fmt"
	"os"
)

// 모듈러 가우스 소거법으로 행렬 랭크를 구한다.
// 소수 모듈러에서 역원을 페르마 소정리로 계산한다.

const mod = 998244353

// power는 밑 base의 exp 거듭제곱을 mod로 나눈 나머지를 구한다.
func power(base, exp, mod int64) int64 {
	result := int64(1)
	base %= mod
	if base < 0 {
		base += mod
	}
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % mod
		}
		exp /= 2
		base = base * base % mod
	}
	return result
}

// modInverse는 a의 모듈러 역원을 구한다 (페르마 소정리).
func modInverse(a, mod int64) int64 {
	return power(a, mod-2, mod)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 행 수 N, 열 수 M
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 입력: N×M 행렬
	a := make([][]int64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a[i][j])
			a[i][j] = (a[i][j]%mod + mod) % mod
		}
	}

	// 모듈러 가우스 소거법으로 랭크 계산
	rank := 0
	for col := 0; col < m && rank < n; col++ {
		// 현재 열에서 0이 아닌 피벗 행을 찾는다
		pivotRow := -1
		for row := rank; row < n; row++ {
			if a[row][col] != 0 {
				pivotRow = row
				break
			}
		}
		if pivotRow == -1 {
			continue // 이 열에는 피벗이 없다
		}

		// 피벗 행과 현재 행을 교환
		a[rank], a[pivotRow] = a[pivotRow], a[rank]

		// 피벗 행을 정규화 (피벗 원소를 1로 만든다)
		inv := modInverse(a[rank][col], mod)
		for j := col; j < m; j++ {
			a[rank][j] = a[rank][j] * inv % mod
		}

		// 피벗 아래의 모든 행을 소거
		for row := 0; row < n; row++ {
			if row != rank && a[row][col] != 0 {
				factor := a[row][col]
				for j := col; j < m; j++ {
					a[row][j] = (a[row][j] - factor*a[rank][j]%mod + mod) % mod
				}
			}
		}

		rank++
	}

	// 출력: 랭크
	fmt.Fprintln(writer, rank)
}
