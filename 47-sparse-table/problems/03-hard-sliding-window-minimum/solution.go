package main

import (
	"bufio"
	"fmt"
	"os"
)

// slidingWindowMinMax는 Sparse Table을 이용하여 크기 K인 슬라이딩 윈도우의
// 최솟값 배열과 최댓값 배열을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - k: 윈도우 크기
//
// [반환값]
//   - []int: 각 윈도우 위치의 최솟값 배열
//   - []int: 각 윈도우 위치의 최댓값 배열
func slidingWindowMinMax(arr []int, k int) ([]int, []int) {
	// 여기에 코드를 작성하세요
	return nil, nil
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

	mins, maxs := slidingWindowMinMax(arr, k)

	for i, v := range mins {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	for i, v := range maxs {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
