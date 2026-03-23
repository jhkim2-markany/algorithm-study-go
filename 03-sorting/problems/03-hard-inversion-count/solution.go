package main

import (
	"bufio"
	"fmt"
	"os"
)

// countInversions는 배열에서 역전(Inversion)의 총 개수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열 (길이 N, 1 ≤ N ≤ 500,000)
//
// [반환값]
//   - int64: 역전의 총 개수 (i < j이면서 arr[i] > arr[j]인 쌍의 수)
func countInversions(arr []int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := countInversions(arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
