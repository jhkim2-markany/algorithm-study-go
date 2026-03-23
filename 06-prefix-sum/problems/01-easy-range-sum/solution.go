package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 N과 쿼리 수 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 누적합 배열 구축 (1-indexed)
	// prefix[i] = arr[0] + arr[1] + ... + arr[i-1]
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	// 각 쿼리에 대해 구간 합 계산
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		// 구간 [l, r]의 합 = prefix[r] - prefix[l-1] (1-indexed)
		fmt.Fprintln(writer, prefix[r]-prefix[l-1])
	}
}
