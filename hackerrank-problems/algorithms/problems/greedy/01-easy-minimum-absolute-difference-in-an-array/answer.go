package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// minimumAbsoluteDifference는 배열에서 서로 다른 두 원소의 차이의 절댓값 중 최솟값을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - int: 최소 절대 차이
//
// [알고리즘 힌트]
//
//	배열을 정렬하면 최소 절대 차이는 반드시 인접한 두 원소 사이에서 발생한다.
//	정렬 후 인접 쌍의 차이만 비교하면 된다.
func minimumAbsoluteDifference(arr []int) int {
	// 배열을 오름차순으로 정렬
	sort.Ints(arr)

	// 최솟값을 최대 정수로 초기화
	minDiff := math.MaxInt64

	// 인접한 두 원소의 차이를 비교하여 최솟값 갱신
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := minimumAbsoluteDifference(arr)
	fmt.Fprintln(writer, result)
}
