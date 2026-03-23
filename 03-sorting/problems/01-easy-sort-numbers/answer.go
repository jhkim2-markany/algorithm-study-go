package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// sortNumbers는 정수 배열을 오름차순으로 정렬하여 반환한다.
//
// [매개변수]
//   - arr: 정렬할 정수 배열 (길이 N, 1 ≤ N ≤ 1,000,000)
//
// [반환값]
//   - []int: 오름차순으로 정렬된 정수 배열
//
// [알고리즘 힌트]
//
//	Go 표준 라이브러리의 sort.Ints()를 활용하면 간단하게 정렬할 수 있다.
//	sort.Ints()는 내부적으로 인트로소트(Introsort)를 사용하며,
//	평균 시간복잡도는 O(N log N)이다.
//
//	주의: sort.Ints()는 슬라이스를 직접 수정(in-place)한다.
//
//	예시: arr=[5, 3, 1, 4, 2]
//	  → [1, 2, 3, 4, 5]
func sortNumbers(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정수 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// N개의 정수 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	sorted := sortNumbers(arr)

	// 결과 출력
	for i := 0; i < len(sorted); i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, sorted[i])
	}
	fmt.Fprintln(writer)
}
