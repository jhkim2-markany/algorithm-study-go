package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathUpdate는 오일러 투어와 펜윅 트리(차분 배열)를 이용하여
// 루트→v 경로 갱신과 노드 값 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (u, v 쌍)
//   - queries: 질의 목록 (타입, 인자들)
//
// [반환값]
//   - []int: 노드 값 질의(타입 2)의 결과 배열
func pathUpdate(n int, edges [][2]int, queries [][]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var v, x int
			fmt.Fscan(reader, &v, &x)
			queries[i] = []int{t, v, x}
		} else {
			var v int
			fmt.Fscan(reader, &v)
			queries[i] = []int{t, v}
		}
	}

	results := pathUpdate(n, edges, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
