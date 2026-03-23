package main

import (
	"bufio"
	"fmt"
	"os"
)

// lisWithPath는 최장 증가 부분 수열(LIS)의 길이와 실제 LIS를 반환한다.
//
// [매개변수]
//   - a: 정수 수열
//
// [반환값]
//   - int: LIS의 길이
//   - []int: 실제 LIS 원소 배열
func lisWithPath(a []int) (int, []int) {
	// 여기에 코드를 작성하세요
	return 0, nil
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

	lisLen, result := lisWithPath(a)

	fmt.Fprintln(writer, lisLen)
	for i := 0; i < lisLen; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, result[i])
	}
	fmt.Fprintln(writer)
}
