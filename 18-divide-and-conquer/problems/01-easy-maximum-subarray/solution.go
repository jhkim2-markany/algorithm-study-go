package main

import (
	"bufio"
	"fmt"
	"os"
)

// maxSubarraySum은 분할 정복으로 배열의 최대 부분 배열 합을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - int: 최대 부분 배열 합
func maxSubarraySum(arr []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := maxSubarraySum(arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
