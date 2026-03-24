package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// maxMin은 배열에서 K개의 원소를 선택하여 최소 불공정도를 반환한다.
//
// [매개변수]
//   - k: 선택할 원소의 수
//   - arr: 정수 배열
//
// [반환값]
//   - int: 최소 불공정도
//
// [알고리즘 힌트]
//
//	배열을 정렬하면 최소 불공정도는 연속된 K개 원소의 윈도우에서 발생한다.
//	슬라이딩 윈도우로 arr[i+k-1] - arr[i]의 최솟값을 구한다.
func maxMin(k int, arr []int) int {
	// 배열을 오름차순으로 정렬
	sort.Ints(arr)

	// 최솟값을 최대 정수로 초기화
	minUnfairness := math.MaxInt64

	// 크기 K인 슬라이딩 윈도우로 최소 불공정도 탐색
	for i := 0; i+k-1 < len(arr); i++ {
		unfairness := arr[i+k-1] - arr[i]
		if unfairness < minUnfairness {
			minUnfairness = unfairness
		}
	}

	return minUnfairness
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N, K 입력
	var n, k int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := maxMin(k, arr)
	fmt.Fprintln(writer, result)
}
