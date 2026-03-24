package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// twoArrays는 두 배열을 재배열하여 모든 쌍의 합이 K 이상이 될 수 있는지 판별한다.
//
// [매개변수]
//   - k: 목표 합
//   - A: 첫 번째 배열
//   - B: 두 번째 배열
//
// [반환값]
//   - string: "YES" 또는 "NO"
//
// [알고리즘 힌트]
//
//	A를 오름차순, B를 내림차순으로 정렬한다.
//	A의 가장 작은 값과 B의 가장 큰 값을 매칭하여 합이 K 이상인지 확인한다.
func twoArrays(k int, A []int, B []int) string {
	// A를 오름차순으로 정렬
	sort.Ints(A)
	// B를 내림차순으로 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	// 각 쌍의 합이 K 이상인지 확인
	for i := 0; i < len(A); i++ {
		if A[i]+B[i] < k {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 쿼리 수 입력
	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		// N, K 입력
		var n, k int
		fmt.Fscan(reader, &n, &k)

		// 배열 A 입력
		A := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &A[i])
		}

		// 배열 B 입력
		B := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &B[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := twoArrays(k, A, B)
		fmt.Fprintln(writer, result)
	}
}
