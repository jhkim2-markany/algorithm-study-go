package main

import (
	"bufio"
	"fmt"
	"os"
)

// factorize는 자연수 N을 소인수분해하여 소인수 목록을 오름차순으로 반환한다.
// 같은 소인수가 여러 번 나누어지면 그 횟수만큼 포함한다.
//
// [매개변수]
//   - n: 소인수분해할 자연수 (n >= 2)
//
// [반환값]
//   - []int: 소인수 목록 (오름차순, 중복 포함)
func factorize(n int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	factors := factorize(n)
	for _, f := range factors {
		fmt.Fprintln(writer, f)
	}
}
