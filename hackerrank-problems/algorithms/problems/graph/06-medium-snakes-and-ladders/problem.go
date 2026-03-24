package main

import (
	"bufio"
	"fmt"
	"os"
)

// quickestWayUp은 뱀과 사다리 게임에서 100번 칸에 도달하는 최소 주사위 횟수를 반환한다.
//
// [매개변수]
//   - ladders: 사다리 목록 (각 원소는 [2]int{시작, 끝})
//   - snakes: 뱀 목록 (각 원소는 [2]int{시작, 끝})
//
// [반환값]
//   - int: 최소 주사위 횟수 (-1이면 도달 불가)
func quickestWayUp(ladders [][2]int, snakes [][2]int) int {
	// 여기에 코드를 작성하세요
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(reader, &n)
		ladders := make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &ladders[i][0], &ladders[i][1])
		}

		var m int
		fmt.Fscan(reader, &m)
		snakes := make([][2]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &snakes[i][0], &snakes[i][1])
		}

		result := quickestWayUp(ladders, snakes)
		fmt.Fprintln(writer, result)
	}
}
