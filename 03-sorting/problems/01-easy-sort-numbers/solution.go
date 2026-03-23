package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정수 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// N개의 정수 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 오름차순 정렬
	sort.Ints(arr)

	// 결과 출력
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, arr[i])
	}
	fmt.Fprintln(writer)
}
