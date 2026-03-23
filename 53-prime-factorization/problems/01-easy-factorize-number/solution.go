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

	// 입력: 자연수 N
	var n int
	fmt.Fscan(reader, &n)

	// 시행 나눗셈으로 소인수분해한다
	// 2부터 시작하여 나누어 떨어지면 반복하여 나눈다
	for d := 2; d*d <= n; d++ {
		for n%d == 0 {
			fmt.Fprintln(writer, d)
			n /= d
		}
	}

	// 남은 수가 1보다 크면 그 자체가 소인수이다
	if n > 1 {
		fmt.Fprintln(writer, n)
	}
}
