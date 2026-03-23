package main

import (
	"bufio"
	"fmt"
	"os"
)

// countInversions는 병합 정렬을 이용하여 배열의 역전 수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열 (함수 내에서 정렬됨)
//
// [반환값]
//   - int64: 역전의 수
func countInversions(arr []int) int64 {
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
	result := countInversions(arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
