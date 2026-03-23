package main

import (
	"bufio"
	"fmt"
	"os"
)

// 직선 컨테이너 (Line Container)
// 직선 추가와 최솟값 쿼리가 임의 순서로 주어진다.
// Li Chao Tree를 사용하여 삽입 O(log X), 쿼리 O(log X)에 처리한다.
//
// Li Chao Tree:
// - 세그먼트 트리 기반 자료구조
// - 각 노드는 해당 x 구간에서 "우세한" 직선을 저장
// - 삽입 시 중간점에서 비교하여 우세 직선을 결정하고, 나머지를 자식으로 내려보냄
// - 쿼리 시 루트에서 리프까지 내려가며 각 노드의 직선 값 중 최솟값을 반환

const (
	MINX = -1000001 // x 좌표 범위 하한
	MAXX = 1000002  // x 좌표 범위 상한 (exclusive)
	INF  = int64(2e18)
)

// Line은 일차 함수 y = m*x + b를 나타낸다
type Line struct {
	m, b int64
}

func (l Line) eval(x int64) int64 {
	return l.m*x + l.b
}

// Node는 Li Chao Tree의 노드이다
type Node struct {
	line        Line
	hasLine     bool
	left, right *Node
}

// insert는 Li Chao Tree에 직선을 삽입한다
func insert(node **Node, lo, hi int64, newLine Line) {
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
	// 왼쪽 끝에서 새 직선이 더 좋은지 확인
	leftBetter := newLine.eval(lo) < nd.line.eval(lo)
	// 중간점에서 새 직선이 더 좋은지 확인
	midBetter := newLine.eval(mid) < nd.line.eval(mid)

	if midBetter {
		// 새 직선이 중간점에서 더 좋으므로 현재 노드의 직선과 교체
		nd.line, newLine = newLine, nd.line
		leftBetter = !leftBetter
	}

	// 구간이 1이면 더 이상 내려갈 필요 없음
	if lo+1 >= hi {
		return
	}

	if leftBetter {
		// 밀려난 직선이 왼쪽 구간에서 더 좋을 수 있음
		insert(&nd.left, lo, mid, newLine)
	} else {
		// 밀려난 직선이 오른쪽 구간에서 더 좋을 수 있음
		insert(&nd.right, mid, hi, newLine)
	}
}

// query는 x에서의 최솟값을 반환한다
func query(node *Node, lo, hi, x int64) int64 {
	if node == nil || !node.hasLine {
		return INF
	}

	res := node.line.eval(x)
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	var root *Node

	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			// 직선 추가: y = mx + b
			var m, b int64
			fmt.Fscan(reader, &m, &b)
			insert(&root, int64(MINX), int64(MAXX), Line{m, b})
		} else {
			// 최솟값 쿼리
			var x int64
			fmt.Fscan(reader, &x)
			result := query(root, int64(MINX), int64(MAXX), x)
			fmt.Fprintln(writer, result)
		}
	}
}
