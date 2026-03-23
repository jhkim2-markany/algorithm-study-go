package main

import (
	"bufio"
	"fmt"
	"os"
)

// matrixRank는 모듈러 가우스 소거법으로 행렬의 랭크를 구한다.
//
// [매개변수]
//   - n: 행 수
//   - m: 열 수
//   - a: N×M 행렬 (모듈러 998244353 위의 값)
//
// [반환값]
//   - int: 행렬의 랭크
//
// [알고리즘 힌트]
//   1. 각 열에서 0이 아닌 피벗 행을 찾아 교환한다
//   2. 피벗 행을 정규화한다 (피벗 원소의 모듈러 역원을 곱한다)
//   3. 모듈러 역원은 페르마 소정리로 계산한다: a^(p-2) mod p
//   4. 다른 모든 행에서 해당 열을 소거한다
//   5. 피벗이 존재하는 열의 수가 랭크이다
func matrixRank(n, m int, a [][]int64) int {
	const mod = int64(998244353)

	power := func(base, exp, mod int64) int64 {
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

	modInverse := func(a, mod int64) int64 {
		return power(a, mod-2, mod)
	}

	rank := 0
	for col := 0; col < m && rank < n; col++ {
		pivotRow := -1
		for row := rank; row < n; row++ {
			if a[row][col] != 0 {
				pivotRow = row
				break
			}
		}
		if pivotRow == -1 {
			continue
		}

		a[rank], a[pivotRow] = a[pivotRow], a[rank]

		inv := modInverse(a[rank][col], mod)
		for j := col; j < m; j++ {
			a[rank][j] = a[rank][j] * inv % mod
		}

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

	return rank
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	const mod = int64(998244353)
	a := make([][]int64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a[i][j])
			a[i][j] = (a[i][j]%mod + mod) % mod
		}
	}

	fmt.Fprintln(writer, matrixRank(n, m, a))
}
