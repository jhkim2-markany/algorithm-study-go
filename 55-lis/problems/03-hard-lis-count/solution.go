package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

// lisCount는 최장 증가 부분 수열(LIS)의 길이와 개수를 반환한다.
//
// [매개변수]
//   - a: 정수 수열
//
// [반환값]
//   - int: LIS의 길이
//   - int: LIS의 개수 (mod 1,000,000,007)
func lisCount(a []int) (int, int) {
	// 여기에 코드를 작성하세요
	return 0, 0
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

	lisLen, totalCount := lisCount(a)

	fmt.Fprintln(writer, lisLen)
	fmt.Fprintln(writer, totalCount)
}
