package main

import (
	"bufio"
	"fmt"
	"os"
)

// kthMinimum은 구간 [l, r]에서 k번째로 작은 원소를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - queries: 쿼리 목록 (각 쿼리는 [l, r, k])
//
// [반환값]
//   - []int: 각 쿼리에 대한 결과
func kthMinimum(arr []int, queries [][]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 쿼리 개수 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 쿼리 입력
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var l, r, k int
		fmt.Fscan(reader, &l, &r, &k)
		queries[i] = []int{l, r, k}
	}

	// 핵심 함수 호출
	results := kthMinimum(arr, queries)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
