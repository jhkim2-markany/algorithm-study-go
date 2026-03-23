package main

import (
	"bufio"
	"fmt"
	"os"
)

// jobScheduling은 마감 기한과 보상이 주어진 작업들에서 최대 보상을 반환한다.
//
// [매개변수]
//   - deadlines: 각 작업의 마감 기한 배열
//   - profits: 각 작업의 보상 배열
//
// [반환값]
//   - int: 얻을 수 있는 최대 보상
func jobScheduling(deadlines, profits []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	deadlines := make([]int, n)
	profits := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &deadlines[i], &profits[i])
	}

	result := jobScheduling(deadlines, profits)
	fmt.Fprintln(writer, result)
}
