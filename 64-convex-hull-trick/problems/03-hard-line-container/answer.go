package main

import (
	"bufio"
	"fmt"
	"os"
)

// lineContainer는 Li Chao Tree를 이용하여 직선 추가와 최솟값 쿼리를 처리한다.
//
// [매개변수]
//   - ops: 연산 목록 (op=1: 직선 추가 [1, m, b], op=2: 최솟값 쿼리 [2, x])
//
// [반환값]
//   - []int64: 최솟값 쿼리(op=2)의 결과 배열
//
// [알고리즘 힌트]
//   1. Li Chao Tree는 세그먼트 트리 기반 자료구조로 직선을 관리한다
//   2. 삽입 시 중간점에서 새 직선과 기존 직선을 비교하여 우세 직선을 결정한다
//   3. 밀려난 직선은 한쪽 자식으로 내려보낸다
//   4. 쿼리 시 루트에서 리프까지 내려가며 각 노드의 직선 값 중 최솟값을 반환한다
func lineContainer(ops [][]int64) []int64 {
	const (
		MINX = int64(-1000001)
		MAXX = int64(1000002)
		INF  = int64(2e18)
	)

	type Line struct {
		m, b int64
	}
	type Node struct {
		line        Line
		hasLine     bool
		left, right *Node
	}

	var insert func(node **Node, lo, hi int64, newLine Line)
	insert = func(node **Node, lo, hi int64, newLine Line) {
		if *node == nil {
			*node = &Node{line: newLine, hasLine: true}
			return
		}
		nd := *node
		if !nd.hasLine {
			nd.line = newLine
			nd.hasLine = true
			return
		}
		mid := (lo + hi) / 2
		leftBetter := newLine.m*lo+newLine.b < nd.line.m*lo+nd.line.b
		midBetter := newLine.m*mid+newLine.b < nd.line.m*mid+nd.line.b
		if midBetter {
			nd.line, newLine = newLine, nd.line
			leftBetter = !leftBetter
		}
		if lo+1 >= hi {
			return
		}
		if leftBetter {
			insert(&nd.left, lo, mid, newLine)
		} else {
			insert(&nd.right, mid, hi, newLine)
		}
	}

	var query func(node *Node, lo, hi, x int64) int64
	query = func(node *Node, lo, hi, x int64) int64 {
		if node == nil || !node.hasLine {
			return INF
		}
		res := node.line.m*x + node.line.b
		mid := (lo + hi) / 2
		if lo+1 >= hi {
			return res
		}
		var childRes int64
		if x < mid {
			childRes = query(node.left, lo, mid, x)
		} else {
			childRes = query(node.right, mid, hi, x)
		}
		if childRes < res {
			return childRes
		}
		return res
	}

	var root *Node
	var results []int64

	for _, op := range ops {
		if op[0] == 1 {
			insert(&root, MINX, MAXX, Line{op[1], op[2]})
		} else {
			results = append(results, query(root, MINX, MAXX, op[1]))
		}
	}

	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	ops := make([][]int64, q)
	for i := 0; i < q; i++ {
		var op int64
		fmt.Fscan(reader, &op)
		if op == 1 {
			var m, b int64
			fmt.Fscan(reader, &m, &b)
			ops[i] = []int64{op, m, b}
		} else {
			var x int64
			fmt.Fscan(reader, &x)
			ops[i] = []int64{op, x}
		}
	}

	results := lineContainer(ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
