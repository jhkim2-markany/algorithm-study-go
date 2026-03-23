package main

import (
	"bufio"
	"fmt"
	"os"
)

// activitySelection은 회의 목록에서 겹치지 않게 최대한 많은 회의를 선택한 수를 반환한다.
//
// [매개변수]
//   - starts: 각 회의의 시작 시간 배열
//   - ends: 각 회의의 종료 시간 배열
//
// [반환값]
//   - int: 겹치지 않게 선택할 수 있는 최대 회의 수
func activitySelection(starts, ends []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	starts := make([]int, n)
	ends := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &starts[i], &ends[i])
	}

	result := activitySelection(starts, ends)
	fmt.Fprintln(writer, result)
}
