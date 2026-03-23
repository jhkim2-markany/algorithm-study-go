package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	// 각 질의에 대해 이진 탐색 수행
	for q := 0; q < m; q++ {
		var target int
		fmt.Fscan(reader, &target)

		// 이진 탐색으로 target의 인덱스를 찾는다
		lo, hi := 0, n-1
		result := -1

		for lo <= hi {
			mid := (lo + hi) / 2

			if arr[mid] == target {
				// 값을 찾은 경우
				result = mid
				break
			} else if arr[mid] < target {
				// 오른쪽 절반 탐색
				lo = mid + 1
			} else {
				// 왼쪽 절반 탐색
				hi = mid - 1
			}
		}

		fmt.Fprintln(writer, result)
	}
}
