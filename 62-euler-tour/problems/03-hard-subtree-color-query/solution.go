package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// subtreeColorQuery는 오일러 투어와 Mo's Algorithm(with Updates)을 이용하여
// 서브트리 내 서로 다른 색상 수 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - c: 색상 종류 수
//   - initColor: 각 노드의 초기 색상 (1-indexed)
//   - edges: 간선 목록 (u, v 쌍)
//   - ops: 연산 목록 (타입1: 색상 변경, 타입2: 서브트리 색상 수 질의)
//
// [반환값]
//   - []int: 서브트리 색상 수 질의(타입 2)의 결과 배열
func subtreeColorQuery(n, c int, initColor []int, edges [][2]int, ops [][]int) []int {
	// 여기에 코드를 작성하세요
	_ = math.Max
	_ = sort.Slice
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, c, q int
	fmt.Fscan(reader, &n, &c, &q)

	initColor := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &initColor[i])
	}

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var v, newC int
			fmt.Fscan(reader, &v, &newC)
			ops[i] = []int{t, v, newC}
		} else {
			var v int
			fmt.Fscan(reader, &v)
			ops[i] = []int{t, v}
		}
	}

	results := subtreeColorQuery(n, c, initColor, edges, ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
