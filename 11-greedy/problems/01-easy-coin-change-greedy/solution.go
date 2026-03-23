package main

import (
	"bufio"
	"fmt"
	"os"
)

// coinChangeGreedy는 주어진 금액을 거슬러 주는 데 필요한 최소 동전 수를 반환한다.
//
// [매개변수]
//   - n: 거슬러 줘야 할 금액 (10의 배수)
//
// [반환값]
//   - int: 필요한 최소 동전 수
func coinChangeGreedy(n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	result := coinChangeGreedy(n)
	fmt.Fprintln(writer, result)
}
