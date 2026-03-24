package main

import (
	"bufio"
	"fmt"
	"os"
)

// evenTree는 각 연결 요소의 노드 수가 짝수가 되도록 제거할 수 있는 간선의 최대 수를 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 원소는 [2]int{u, v}, v는 u의 부모)
//
// [반환값]
//   - int: 제거 가능한 간선의 최대 수
func evenTree(n int, edges [][2]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	result := evenTree(n, edges)
	fmt.Fprintln(writer, result)
}
