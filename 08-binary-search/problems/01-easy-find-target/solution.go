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
func findTarget(arr []int, target int) int {
	// 여기에 코드를 작성하세요
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
