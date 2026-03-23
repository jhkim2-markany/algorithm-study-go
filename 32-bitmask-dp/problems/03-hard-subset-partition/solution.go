package main

import (
	"bufio"
	"fmt"
	"os"
)

// minPartitionDiff는 집합을 두 부분집합으로 나눌 때 합의 차이의 최솟값을 반환한다.
//
// [매개변수]
//   - n: 원소의 수
//   - a: 원소 배열 (길이 n)
//
// [반환값]
//   - int: 두 부분집합 합의 차이의 최솟값
func minPartitionDiff(n int, a []int) int {
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

	fmt.Fprintln(writer, minPartitionDiff(n, a))
}
