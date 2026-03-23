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

	// 원소 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 0부터 2^N - 1까지 모든 비트마스크를 순회한다
	total := 1 << n
	for mask := 0; mask < total; mask++ {
		first := true
		// 각 비트를 확인하여 포함된 원소를 출력한다
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if !first {
					fmt.Fprint(writer, " ")
				}
				// 원소는 1부터 시작하므로 i+1을 출력한다
				fmt.Fprint(writer, i+1)
				first = false
			}
		}
		fmt.Fprintln(writer)
	}
}
