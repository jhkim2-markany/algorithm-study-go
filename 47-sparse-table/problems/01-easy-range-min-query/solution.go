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
func queryRangeMin(arr []int, queries [][2]int) []int {
	// 여기에 코드를 작성하세요
	return nil
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
