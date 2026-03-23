package main

import (
	"bufio"
	"fmt"
	"os"
)

// subarraySumEqualsK는 합이 k인 연속 부분배열의 개수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - k: 목표 합
//
// [반환값]
//   - int: 합이 k인 연속 부분배열의 개수
func subarraySumEqualsK(arr []int, k int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 N과 목표 합 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	count := subarraySumEqualsK(arr, k)

	// 결과 출력
	fmt.Fprintln(writer, count)
}
