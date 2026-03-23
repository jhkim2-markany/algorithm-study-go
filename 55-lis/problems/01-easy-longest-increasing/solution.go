package main

import (
	"bufio"
	"fmt"
	"os"
)

// lisLength는 주어진 수열에서 최장 증가 부분 수열(LIS)의 길이를 반환한다.
//
// [매개변수]
//   - a: 정수 수열
//
// [반환값]
//   - int: LIS의 길이
func lisLength(a []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	fmt.Fprintln(writer, lisLength(a))
}
