package main

import (
	"bufio"
	"fmt"
	"os"
)

// slidingWindowMax는 모노톤 덱을 이용하여 크기 K인 슬라이딩 윈도우의 최댓값 배열을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - k: 윈도우 크기
//
// [반환값]
//   - []int: 각 윈도우 위치의 최댓값 배열
func slidingWindowMax(arr []int, k int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	result := slidingWindowMax(arr, k)
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
