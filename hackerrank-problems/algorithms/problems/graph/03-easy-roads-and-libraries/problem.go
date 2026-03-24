package main

import (
	"bufio"
	"fmt"
	"os"
)

// roadsAndLibraries는 모든 도시에서 도서관에 접근 가능하게 하는 최소 비용을 반환한다.
//
// [매개변수]
//   - n: 도시 수
//   - cLib: 도서관 건설 비용
//   - cRoad: 도로 수리 비용
//   - edges: 도로 목록 (각 원소는 [2]int{u, v})
//
// [반환값]
//   - int64: 최소 비용
func roadsAndLibraries(n int, cLib int, cRoad int, edges [][2]int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		var n, m, cLib, cRoad int
		fmt.Fscan(reader, &n, &m, &cLib, &cRoad)

		edges := make([][2]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &edges[i][0], &edges[i][1])
		}

		result := roadsAndLibraries(n, cLib, cRoad, edges)
		fmt.Fprintln(writer, result)
	}
}
