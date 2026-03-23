package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// compress는 정수 수열에 대해 좌표 압축을 수행하여 각 원소의 순위를 반환한다.
//
// [매개변수]
//   - arr: 원본 정수 수열
//
// [반환값]
//   - []int: 각 원소의 압축된 좌표(순위) 배열 (0-indexed)
//
// [알고리즘 힌트]
//   1. 원본 배열을 복사한 뒤 정렬한다.
//   2. 정렬된 배열에서 중복을 제거하여 고유 좌표 배열을 만든다.
//   3. 각 원소에 대해 이진 탐색(sort.SearchInts)으로 고유 좌표 배열에서의 위치를 찾는다.
func compress(arr []int) []int {
	n := len(arr)

	// 정렬 후 중복 제거
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	unique := []int{sorted[0]}
	for i := 1; i < n; i++ {
		if sorted[i] != sorted[i-1] {
			unique = append(unique, sorted[i])
		}
	}

	// 이진 탐색으로 각 원소의 압축된 좌표를 구한다
	ranks := make([]int, n)
	for i := 0; i < n; i++ {
		ranks[i] = sort.SearchInts(unique, arr[i])
	}
	return ranks
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	ranks := compress(arr)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, ranks[i])
	}
	fmt.Fprintln(writer)
}
