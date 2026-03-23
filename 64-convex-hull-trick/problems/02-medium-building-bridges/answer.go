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
//
// [알고리즘 힌트]
//   1. dp[i] = h[i]² + sumW[i-1] + min((-2·h[j])·h[i] + dp[j] + h[j]² - sumW[j])
//   2. f_j(x) = (-2·h[j])·x + (dp[j] + h[j]² - sumW[j])로 일차 함수를 정의한다
//   3. h가 정렬되지 않으므로 Li Chao Tree를 사용하여 삽입/쿼리를 O(log X)에 처리한다
//   4. Li Chao Tree의 각 노드는 중간점에서 우세한 직선을 저장한다
func buildingBridges(n int, h, w []int64) int64 {
	const INF = int64(1e18)

	sumW := make([]int64, n)
	sumW[0] = w[0]
	for i := 1; i < n; i++ {
		sumW[i] = sumW[i-1] + w[i]
	}

	type Line struct {
		m, b int64
	}
	type Node struct {
		line        Line
		hasLine     bool
		left, right *Node
	}

	var lo, hi int64 = 0, 1000001

	var insert func(node **Node, l, r int64, newLine Line)
	insert = func(node **Node, l, r int64, newLine Line) {
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
		mid := (l + r) / 2
		leftBetter := newLine.m*l+newLine.b < nd.line.m*l+nd.line.b
		midBetter := newLine.m*mid+newLine.b < nd.line.m*mid+nd.line.b
		if midBetter {
			nd.line, newLine = newLine, nd.line
			leftBetter = !leftBetter
		}
		if l+1 >= r {
			return
		}
		if leftBetter {
			insert(&nd.left, l, mid, newLine)
		} else {
			insert(&nd.right, mid, r, newLine)
		}
	}

	var query func(node *Node, l, r, x int64) int64
	query = func(node *Node, l, r, x int64) int64 {
		if node == nil || !node.hasLine {
			return INF
		}
		mid := (l + r) / 2
		res := node.line.m*x + node.line.b
		if x < mid {
			childRes := query(node.left, l, mid, x)
			if childRes < res {
				res = childRes
			}
		} else {
			childRes := query(node.right, mid, r, x)
			if childRes < res {
				res = childRes
			}
		}
		return res
	}

	dp := make([]int64, n)
	dp[0] = 0

	var root *Node
	insert(&root, lo, hi, Line{-2 * h[0], dp[0] + h[0]*h[0] - sumW[0]})

	for i := 1; i < n; i++ {
		minVal := query(root, lo, hi, h[i])
		var prevSumW int64
		if i > 0 {
			prevSumW = sumW[i-1]
		}
		dp[i] = h[i]*h[i] + prevSumW + minVal
		insert(&root, lo, hi, Line{-2 * h[i], dp[i] + h[i]*h[i] - sumW[i]})
	}

	return dp[n-1]
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
