package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathMaxEdge는 HLD와 세그먼트 트리(최댓값)를 이용하여 트리에서
// 간선 가중치 갱신과 경로 최대 간선 가중치 질의를 처리한다.
// 간선 가중치는 자식 노드에 매핑하여 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - edgeList: 간선 목록 (a, b, w)
//   - ops: 연산 목록 (op=1: 간선 갱신 [1,idx,w], op=2: 경로 질의 [2,u,v])
//
// [반환값]
//   - []int: 경로 최대 간선 가중치 질의(op=2)의 결과 배열
func pathMaxEdge(n int, edgeList [][3]int, ops [][]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	edgeList := make([][3]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edgeList[i][0], &edgeList[i][1], &edgeList[i][2])
	}

	var q int
	fmt.Fscan(reader, &q)

	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)
		if op == 1 {
			var idx, w int
			fmt.Fscan(reader, &idx, &w)
			ops[i] = []int{op, idx, w}
		} else {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			ops[i] = []int{op, u, v}
		}
	}

	results := pathMaxEdge(n, edgeList, ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
