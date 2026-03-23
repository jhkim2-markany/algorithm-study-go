package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcd는 두 정수의 최대공약수를 반환한다.
//
// [매개변수]
//   - a: 첫 번째 정수
//   - b: 두 번째 정수
//
// [반환값]
//   - int: a와 b의 최대공약수
//
// [알고리즘 힌트]
//
//	유클리드 호제법: b가 0이 될 때까지 a, b = b, a%b를 반복한다.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// queryRangeGCD는 Sparse Table을 이용하여 배열의 구간 GCD 쿼리를 처리한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - queries: 각 쿼리는 [l, r] 형태의 1-indexed 구간
//
// [반환값]
//   - []int: 각 쿼리에 대한 구간 GCD 결과 배열
//
// [알고리즘 힌트]
//
//	GCD는 멱등 연산이므로 Sparse Table로 O(1) 쿼리가 가능하다.
//	sparse[k][i] = gcd(arr[i], ..., arr[i+2^k-1])
//	쿼리: gcd(sparse[k][l], sparse[k][r-2^k+1])
func queryRangeGCD(arr []int, queries [][2]int) []int {
	n := len(arr)

	logTable := make([]int, n+2)
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	maxK := logTable[n] + 1
	sparse := make([][]int, maxK)
	for k := 0; k < maxK; k++ {
		sparse[k] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		sparse[0][i] = arr[i]
	}
	for k := 1; k < maxK; k++ {
		for i := 0; i+(1<<k)-1 < n; i++ {
			sparse[k][i] = gcd(sparse[k-1][i], sparse[k-1][i+(1<<(k-1))])
		}
	}

	results := make([]int, len(queries))
	for idx, q := range queries {
		l, r := q[0]-1, q[1]-1
		length := r - l + 1
		k := logTable[length]
		results[idx] = gcd(sparse[k][l], sparse[k][r-(1<<k)+1])
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	queries := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	results := queryRangeGCD(arr, queries)
	for _, v := range results {
		fmt.Fprintln(writer, v)
	}
}
