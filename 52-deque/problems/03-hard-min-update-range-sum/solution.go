package main

import (
	"bufio"
	"fmt"
	"os"
)

// minUpdateRangeSum은 배열에서 각 원소가 포함되는 크기 K 윈도우의 최솟값 중
// 최댓값을 B[i]로 정의할 때, B[i]의 총합을 반환한다.
//
// [매개변수]
//   - a: 정수 배열
//   - k: 윈도우 크기
//
// [반환값]
//   - int64: B[i]의 총합
func minUpdateRangeSum(a []int, k int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	fmt.Fprintln(writer, minUpdateRangeSum(a, k))
}
