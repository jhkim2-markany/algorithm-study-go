package main

import (
	"bufio"
	"fmt"
	"os"
)

// 다리 놓기 (Building Bridges)
//
// dp[i] = min(dp[j] + (h[i] - h[j])² + cost[i] - cost[j])
// 여기서 cost[i]는 1~i 기둥 중 i를 제외한 나머지의 무게 합 (누적 제거 비용)
//
// 정리하면:
// dp[i] = min(dp[j] + h[i]² - 2·h[i]·h[j] + h[j]²) + sumW[i] - sumW[j] - w[i]
//       (sumW[i] = w[1]+...+w[i], 제거 비용 = sumW[n] - 남긴 기둥 무게 합)
//
// 더 간단하게:
// dp[i] = i번 기둥까지 도달하는 최소 비용 (i번 기둥은 반드시 남김)
// dp[i] = min(dp[j] + (h[i]-h[j])² + (sumW[i-1] - sumW[j]))  (j < i)
//       = min(dp[j] + h[i]² - 2·h[i]·h[j] + h[j]² + sumW[i-1] - sumW[j])
//
// 일차 함수 형태:
// dp[i] = h[i]² + sumW[i-1] + min((-2·h[j])·h[i] + (dp[j] + h[j]² - sumW[j]))
//
// f_j(x) = (-2·h[j])·x + (dp[j] + h[j]² - sumW[j])
// 쿼리 x = h[i]
//
// 기울기 -2·h[j]가 단조가 아닐 수 있으므로 (h가 정렬되지 않음)
// → 직선을 기울기 순으로 정렬하여 추가하고, 이분 탐색으로 쿼리한다
// 또는 Li Chao Tree를 사용한다
//
// 여기서는 기울기 정렬 + 이분 탐색 CHT를 사용한다

// Line은 일차 함수 y = m*x + b를 나타낸다
type Line struct {
	m, b int64
}

func (l Line) eval(x int64) int64 {
	return l.m*x + l.b
}

// CHT는 이분 탐색 기반 볼록 껍질 트릭이다
type CHT struct {
	lines []Line
}

// bad는 직선 b가 a와 c 사이에서 불필요한지 판정한다
func bad(a, b, c Line) bool {
	return float64(c.b-a.b)*float64(a.m-b.m) <= float64(b.b-a.b)*float64(a.m-c.m)
}

// AddLine은 기울기 감소 순서로 직선을 추가한다
func (cht *CHT) AddLine(m, b int64) {
	newLine := Line{m, b}
	for len(cht.lines) >= 2 {
		n := len(cht.lines)
		if bad(cht.lines[n-2], cht.lines[n-1], newLine) {
			cht.lines = cht.lines[:n-1]
		} else {
			break
		}
	}
	cht.lines = append(cht.lines, newLine)
}

// Query는 이분 탐색으로 x에서의 최솟값을 반환한다
func (cht *CHT) Query(x int64) int64 {
	lo, hi := 0, len(cht.lines)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if cht.lines[mid].eval(x) <= cht.lines[mid+1].eval(x) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return cht.lines[lo].eval(x)
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

	// 누적 무게 합
	sumW := make([]int64, n)
	sumW[0] = w[0]
	for i := 1; i < n; i++ {
		sumW[i] = sumW[i-1] + w[i]
	}

	// dp[i] = i번 기둥을 남기고 도달하는 최소 비용
	// dp[i] = h[i]² + (sumW[i-1] if i>0 else 0) + min_j(f_j(h[i]))
	// f_j(x) = (-2·h[j])·x + (dp[j] + h[j]² - sumW[j])
	//
	// 기울기가 단조가 아니므로, 인덱스를 기울기(= -2·h[j]) 감소 순으로 정렬하여 처리
	// → 오프라인으로 처리: 모든 j의 직선을 기울기 순 정렬 후 추가, 쿼리도 정렬

	const INF = int64(1e18)
	dp := make([]int64, n)

	// dp[0] = 0 (첫 번째 기둥은 반드시 남김, 제거 비용 없음)
	dp[0] = 0

	// 오프라인 처리를 위해 인덱스를 기울기 순으로 정렬
	// 하지만 dp[j]가 순차적으로 결정되므로 오프라인이 어렵다
	// → Li Chao Tree 사용

	// Li Chao Tree 구현
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
		leftBetter := newLine.eval(l) < nd.line.eval(l)
		midBetter := newLine.eval(mid) < nd.line.eval(mid)

		if midBetter {
			// 새 직선이 mid에서 더 좋으므로 교체
			nd.line, newLine = newLine, nd.line
			leftBetter = !leftBetter
		}

		// 이제 nd.line이 mid에서 더 좋고, newLine은 한쪽에서 더 좋을 수 있다
		if l+1 >= r {
			return
		}
		if leftBetter {
			// newLine이 왼쪽에서 더 좋을 수 있다
			insert(&nd.left, l, mid, newLine)
		} else {
			// newLine이 오른쪽에서 더 좋을 수 있다
			insert(&nd.right, mid, r, newLine)
		}
	}

	var queryLiChao func(node *Node, l, r, x int64) int64
	queryLiChao = func(node *Node, l, r, x int64) int64 {
		if node == nil || !node.hasLine {
			return INF
		}
		mid := (l + r) / 2
		res := node.line.eval(x)
		if x < mid {
			childRes := queryLiChao(node.left, l, mid, x)
			if childRes < res {
				res = childRes
			}
		} else {
			childRes := queryLiChao(node.right, mid, r, x)
			if childRes < res {
				res = childRes
			}
		}
		return res
	}

	var root *Node

	// j=0 직선 추가
	insert(&root, lo, hi, Line{-2 * h[0], dp[0] + h[0]*h[0] - sumW[0]})

	for i := 1; i < n; i++ {
		// 쿼리: x = h[i]
		minVal := queryLiChao(root, lo, hi, h[i])

		var prevSumW int64
		if i > 0 {
			prevSumW = sumW[i-1]
		}
		dp[i] = h[i]*h[i] + prevSumW + minVal

		// 직선 추가: f_i(x) = (-2·h[i])·x + (dp[i] + h[i]² - sumW[i])
		insert(&root, lo, hi, Line{-2 * h[i], dp[i] + h[i]*h[i] - sumW[i]})
	}

	// 답: dp[n-1] (마지막 기둥은 반드시 남김)
	// 전체 무게에서 남긴 기둥 무게를 빼는 것은 이미 점화식에 포함됨
	fmt.Fprintln(writer, dp[n-1])
}
