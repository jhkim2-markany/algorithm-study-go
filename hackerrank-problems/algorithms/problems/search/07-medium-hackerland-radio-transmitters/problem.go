package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// hackerlandRadioTransmitters는 모든 집에 신호를 보내기 위한 최소 송신기 수를 반환한다.
//
// [매개변수]
//   - x: 집 위치 배열
//   - k: 송신기 도달 범위
//
// [반환값]
//   - int: 최소 송신기 수
func hackerlandRadioTransmitters(x []int, k int) int {
	// 여기에 코드를 작성하세요
	_ = sort.Ints
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	fmt.Fprintln(writer, hackerlandRadioTransmitters(x, k))
}
