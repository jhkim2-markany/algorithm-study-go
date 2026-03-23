package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// abs는 정수의 절댓값을 반환한다
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// checkPermutation은 정렬된 배열의 순열 중 인접 원소 차이가
// 모두 k 이하인 순열을 찾아 반환한다.
//
// [매개변수]
//   - arr: 오름차순 정렬된 정수 배열
//   - k: 인접 원소 간 허용되는 최대 차이
//
// [반환값]
//   - []int: 조건을 만족하는 사전순 가장 앞선 순열
//   - bool: 조건을 만족하는 순열을 찾았으면 true, 없으면 false
func checkPermutation(arr []int, k int) ([]int, bool) {
	// 여기에 코드를 작성하세요
	return nil, false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 사전순으로 가장 앞서는 순열을 찾기 위해 정렬
	sort.Ints(arr)

	// 핵심 함수 호출
	result, found := checkPermutation(arr, k)

	// 결과 출력
	if found {
		for i, v := range result {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
