package main

import (
	"bufio"
	"fmt"
	"os"
)

// Sparse Table을 이용한 구간 GCD 쿼리
// GCD는 멱등 연산이므로 O(1) 쿼리가 가능하다
// 전처리: O(N log N), 쿼리: O(1)

var sparseGCD [17][100001]int
var logTbl [100002]int

// 최대공약수를 구하는 함수
func gcdFunc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기와 쿼리 수
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 로그 값 전처리
	for i := 2; i <= n; i++ {
		logTbl[i] = logTbl[i/2] + 1
	}

	// Sparse Table 구성: 길이 1인 구간 초기화
	for i := 0; i < n; i++ {
		sparseGCD[0][i] = arr[i]
	}

	// 길이 2^k인 구간: 두 구간의 GCD를 합친다
	maxK := logTbl[n] + 1
	for k := 1; k < maxK; k++ {
		for i := 0; i+(1<<k)-1 < n; i++ {
			sparseGCD[k][i] = gcdFunc(sparseGCD[k-1][i], sparseGCD[k-1][i+(1<<(k-1))])
		}
	}

	// 쿼리 처리
	for q := 0; q < m; q++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		// 1-indexed를 0-indexed로 변환한다
		l--
		r--
		// 구간 길이에 해당하는 로그 값을 구한다
		length := r - l + 1
		k := logTbl[length]
		// 겹치는 두 구간의 GCD를 합친다 (멱등 연산)
		result := gcdFunc(sparseGCD[k][l], sparseGCD[k][r-(1<<k)+1])
		fmt.Fprintln(writer, result)
	}
}
