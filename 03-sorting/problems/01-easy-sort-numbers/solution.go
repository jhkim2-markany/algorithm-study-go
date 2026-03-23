package main

import (
	"bufio"
	"fmt"
	"os"
)

// sortNumbers는 정수 배열을 오름차순으로 정렬하여 반환한다.
//
// [매개변수]
//   - arr: 정렬할 정수 배열 (길이 N, 1 ≤ N ≤ 1,000,000)
//
// [반환값]
//   - []int: 오름차순으로 정렬된 정수 배열
func sortNumbers(arr []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정수 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// N개의 정수 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	sorted := sortNumbers(arr)

	// 결과 출력
	for i := 0; i < len(sorted); i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, sorted[i])
	}
	fmt.Fprintln(writer)
}
