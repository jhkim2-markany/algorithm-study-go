package main

import (
	"bufio"
	"fmt"
	"os"
)

// lowerBound는 arr에서 target 이상인 첫 번째 인덱스를 반환한다.
func lowerBound(arr []int, target int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// upperBound는 arr에서 target을 초과하는 첫 번째 인덱스를 반환한다.
func upperBound(arr []int, target int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// countOccurrences는 정렬된 배열에서 x의 등장 횟수를 반환한다.
//
// [매개변수]
//   - arr: 오름차순 정렬된 정수 배열
//   - x: 등장 횟수를 구할 값
//
// [반환값]
//   - int: x의 등장 횟수
//
// [알고리즘 힌트]
//
//	lower_bound와 upper_bound를 이진 탐색으로 구한다.
//	lower_bound: x 이상인 첫 번째 인덱스 (arr[mid] < x → lo = mid+1, else hi = mid)
//	upper_bound: x 초과인 첫 번째 인덱스 (arr[mid] <= x → lo = mid+1, else hi = mid)
//	등장 횟수 = upper_bound - lower_bound
//
//	시간복잡도: O(log N)
func countOccurrences(arr []int, x int) int {
	return upperBound(arr, x) - lowerBound(arr, x)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 N과 질의 수 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 정렬된 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 각 질의에 대해 핵심 함수 호출
	for q := 0; q < m; q++ {
		var x int
		fmt.Fscan(reader, &x)
		fmt.Fprintln(writer, countOccurrences(arr, x))
	}
}
