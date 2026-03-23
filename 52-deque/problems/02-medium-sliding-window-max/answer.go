package main

import (
	"bufio"
	"fmt"
	"os"
)

// slidingWindowMax는 모노톤 덱을 이용하여 크기 K인 슬라이딩 윈도우의 최댓값 배열을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - k: 윈도우 크기
//
// [반환값]
//   - []int: 각 윈도우 위치의 최댓값 배열
//
// [알고리즘 힌트]
//
//	모노톤 덱에 인덱스를 저장하며, 앞쪽이 항상 최댓값의 인덱스이다.
//	윈도우 범위를 벗어난 인덱스를 앞에서 제거하고,
//	현재 원소보다 작거나 같은 원소의 인덱스를 뒤에서 제거한다.
func slidingWindowMax(arr []int, k int) []int {
	n := len(arr)
	deque := make([]int, 0, n)
	result := make([]int, 0, n-k+1)

	for i := 0; i < n; i++ {
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && arr[deque[len(deque)-1]] <= arr[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, arr[deque[0]])
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	result := slidingWindowMax(arr, k)
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
