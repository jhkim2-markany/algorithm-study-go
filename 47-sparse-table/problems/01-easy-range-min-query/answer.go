package main

import (
	"bufio"
	"fmt"
	"os"
)

// queryRangeMin은 Sparse Table을 이용하여 배열의 구간 최솟값 쿼리를 처리한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - queries: 각 쿼리는 [l, r] 형태의 1-indexed 구간
//
// [반환값]
//   - []int: 각 쿼리에 대한 구간 최솟값 결과 배열
//
// [알고리즘 힌트]
//
//	Sparse Table 전처리: sparse[k][i] = 인덱스 i에서 길이 2^k 구간의 최솟값
//	쿼리: 겹치는 두 구간 sparse[k][l]과 sparse[k][r-2^k+1]의 최솟값
//	전처리 O(N log N), 쿼리 O(1)
func queryRangeMin(arr []int, queries [][2]int) []int {
	n := len(arr)

	// 로그 값 전처리
	logTable := make([]int, n+2)
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	// Sparse Table 구성
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
			left := sparse[k-1][i]
			right := sparse[k-1][i+(1<<(k-1))]
			if left <= right {
				sparse[k][i] = left
			} else {
				sparse[k][i] = right
			}
		}
	}

	// 쿼리 처리
	results := make([]int, len(queries))
	for idx, q := range queries {
		l, r := q[0]-1, q[1]-1
		length := r - l + 1
		k := logTable[length]
		left := sparse[k][l]
		right := sparse[k][r-(1<<k)+1]
		if left <= right {
			results[idx] = left
		} else {
			results[idx] = right
		}
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

	results := queryRangeMin(arr, queries)
	for _, v := range results {
		fmt.Fprintln(writer, v)
	}
}
