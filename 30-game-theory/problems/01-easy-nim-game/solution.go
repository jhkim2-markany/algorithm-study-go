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

	// 돌 더미의 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 모든 더미의 돌 개수를 XOR한다
	xorSum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		xorSum ^= a
	}

	// XOR 합이 0이 아니면 선수 승리, 0이면 후수 승리
	if xorSum != 0 {
		fmt.Fprintln(writer, "First")
	} else {
		fmt.Fprintln(writer, "Second")
	}
}
