package main

import (
	"bufio"
	"fmt"
	"os"
)

// findTarget은 정렬된 배열에서 target의 인덱스를 이진 탐색으로 찾아 반환한다.
//
// [매개변수]
//   - arr: 오름차순 정렬된 정수 배열
//   - target: 찾을 값
//
// [반환값]
//   - int: target의 인덱스 (0-indexed), 없으면 -1
//
// [알고리즘 힌트]
//
//	이진 탐색: lo, hi 두 포인터로 탐색 범위를 절반씩 줄인다.
//	mid = (lo + hi) / 2에서 arr[mid]와 target을 비교하여
//	같으면 반환, 작으면 lo = mid + 1, 크면 hi = mid - 1.
//
//	시간복잡도: O(log N)
func findTarget(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
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
		var target int
		fmt.Fscan(reader, &target)
		fmt.Fprintln(writer, findTarget(arr, target))
	}
}
