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
func matrixRank(n, m int, a [][]int64) int {
	// 여기에 코드를 작성하세요
	return 0
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
