package main

import (
	"bufio"
	"fmt"
	"os"
)

// countOccurrences는 정렬된 배열에서 x의 등장 횟수를 반환한다.
//
// [매개변수]
//   - arr: 오름차순 정렬된 정수 배열
//   - x: 등장 횟수를 구할 값
//
// [반환값]
//   - int: x의 등장 횟수
func countOccurrences(arr []int, x int) int {
	// 여기에 코드를 작성하세요
	return 0
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
