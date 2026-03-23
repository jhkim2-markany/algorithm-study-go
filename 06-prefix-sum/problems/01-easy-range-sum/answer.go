package main

import (
	"bufio"
	"fmt"
	"os"
)

// rangeSum은 배열과 구간 쿼리 목록을 받아 각 구간의 합을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열 (0-indexed)
//   - queries: 구간 쿼리 배열, 각 쿼리는 [l, r] (1-indexed)
//
// [반환값]
//   - []int: 각 쿼리에 대한 구간 합 결과 배열
//
// [알고리즘 힌트]
//
//	누적합(Prefix Sum) 배열을 미리 구축하면 각 쿼리를 O(1)에 처리할 수 있다.
//	prefix[i] = arr[0] + arr[1] + ... + arr[i-1] (1-indexed)
//	구간 [l, r]의 합 = prefix[r] - prefix[l-1]
//
//	시간복잡도: 전처리 O(N), 쿼리당 O(1), 전체 O(N + M)
func rangeSum(arr []int, queries [][2]int) []int {
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	results := make([]int, len(queries))
	for i, q := range queries {
		results[i] = prefix[q[1]] - prefix[q[0]-1]
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 N과 쿼리 수 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 쿼리 입력
	queries := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	// 핵심 함수 호출
	results := rangeSum(arr, queries)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
