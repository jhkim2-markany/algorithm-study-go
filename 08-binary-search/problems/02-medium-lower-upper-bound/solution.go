package main

import (
	"bufio"
	"fmt"
	"os"
)

// lowerBound는 arr에서 target 이상인 첫 번째 인덱스를 반환한다
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

// upperBound는 arr에서 target을 초과하는 첫 번째 인덱스를 반환한다
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

	// 각 질의에 대해 등장 횟수 계산
	for q := 0; q < m; q++ {
		var x int
		fmt.Fscan(reader, &x)

		// upper_bound - lower_bound = x의 등장 횟수
		count := upperBound(arr, x) - lowerBound(arr, x)
		fmt.Fprintln(writer, count)
	}
}
