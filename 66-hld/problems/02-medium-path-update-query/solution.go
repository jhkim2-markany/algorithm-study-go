package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathUpdateQuery는 HLD와 Lazy Propagation 세그먼트 트리를 이용하여
// 트리에서 경로 갱신(구간 덧셈)과 경로 합 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (u, v 쌍)
//   - ops: 연산 목록 (op=1: 경로 갱신 [1,u,v,w], op=2: 경로 질의 [2,u,v])
//
// [반환값]
//   - []int64: 경로 합 질의(op=2)의 결과 배열
func pathUpdateQuery(n int, edges [][2]int, ops [][]int) []int64 {
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

	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)
		if op == 1 {
			var u, v, w int
			fmt.Fscan(reader, &u, &v, &w)
			ops[i] = []int{op, u, v, w}
		} else {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			ops[i] = []int{op, u, v}
		}
	}

	results := pathUpdateQuery(n, edges, ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
