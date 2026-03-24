package main

import (
	"bufio"
	"fmt"
	"os"
)

// twoStacks는 두 스택에서 꺼낼 수 있는 최대 원소 개수를 반환한다.
//
// [매개변수]
//   - maxSum: 최대 허용 합
//   - a: 스택 A의 원소 배열 (위에서 아래 순서)
//   - b: 스택 B의 원소 배열 (위에서 아래 순서)
//
// [반환값]
//   - int: 꺼낼 수 있는 최대 원소 개수
func twoStacks(maxSum int, a, b []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 게임 횟수 입력
	var g int
	fmt.Fscan(reader, &g)

	for ; g > 0; g-- {
		var n, m, maxSum int
		fmt.Fscan(reader, &n, &m, &maxSum)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		b := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &b[i])
		}

		// 핵심 함수 호출
		fmt.Fprintln(writer, twoStacks(maxSum, a, b))
	}
}
