package main

import (
	"bufio"
	"fmt"
	"os"
)

// buildingBridges는 기둥의 높이와 무게가 주어질 때, 첫 번째와 마지막 기둥을
// 반드시 남기면서 최소 비용으로 다리를 놓는 비용을 Li Chao Tree를 이용하여 구한다.
//
// [매개변수]
//   - n: 기둥 수
//   - h: 각 기둥의 높이
//   - w: 각 기둥의 무게 (제거 비용)
//
// [반환값]
//   - int64: 최소 비용
func buildingBridges(n int, h, w []int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	h := make([]int64, n)
	w := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &w[i])
	}

	fmt.Fprintln(writer, buildingBridges(n, h, w))
}
