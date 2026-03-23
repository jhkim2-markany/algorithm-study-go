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
func compress(arr []int) []int {
	// 여기에 코드를 작성하세요
	return nil
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

	_ = sort.SearchInts // 패키지 사용 보장
}
