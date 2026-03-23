package main

import (
	"bufio"
	"fmt"
	"os"
)

// 세그먼트 트리와 lazy 배열
var tree []int64
var lazy []int64
var n int

// build 함수는 세그먼트 트리를 구축한다
func build(arr []int64, node, start, end int) {
	if start == end {
		// 리프 노드: 원본 배열 값 저장
		tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	build(arr, 2*node, start, mid)
	build(arr, 2*node+1, mid+1, end)
	// 내부 노드: 자식의 합
	tree[node] = tree[2*node] + tree[2*node+1]
}

// pushDown 함수는 lazy 값을 자식 노드에 전파한다
func pushDown(node, start, end int) {
	if lazy[node] != 0 {
		mid := (start + end) / 2
		// 왼쪽 자식에 lazy 값 전파
		tree[2*node] += lazy[node] * int64(mid-start+1)
		lazy[2*node] += lazy[node]
		// 오른쪽 자식에 lazy 값 전파
		tree[2*node+1] += lazy[node] * int64(end-mid)
		lazy[2*node+1] += lazy[node]
		// 현재 노드의 lazy 값 초기화
		lazy[node] = 0
	}
}

// updateRange 함수는 구간 [l, r]에 val을 더한다
func updateRange(node, start, end, l, r int, val int64) {
	// 구간이 겹치지 않는 경우
	if r < start || end < l {
		return
	}
	// 구간이 완전히 포함되는 경우: lazy 값 저장
	if l <= start && end <= r {
		tree[node] += val * int64(end-start+1)
		lazy[node] += val
		return
	}
	// 부분적으로 겹치는 경우: lazy 전파 후 자식에 재귀
	pushDown(node, start, end)
	mid := (start + end) / 2
	updateRange(2*node, start, mid, l, r, val)
	updateRange(2*node+1, mid+1, end, l, r, val)
	// 부모 노드 재계산
	tree[node] = tree[2*node] + tree[2*node+1]
}

// query 함수는 구간 [l, r]의 합을 반환한다
func query(node, start, end, l, r int) int64 {
	// 구간이 겹치지 않는 경우
	if r < start || end < l {
		return 0
	}
	// 구간이 완전히 포함되는 경우
	if l <= start && end <= r {
		return tree[node]
	}
	// 부분적으로 겹치는 경우: lazy 전파 후 자식에 재귀
	pushDown(node, start, end)
	mid := (start + end) / 2
	return query(2*node, start, mid, l, r) + query(2*node+1, mid+1, end, l, r)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m int
	fmt.Fscan(reader, &n, &m)

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 세그먼트 트리 및 lazy 배열 초기화
	tree = make([]int64, 4*n)
	lazy = make([]int64, 4*n)
	build(arr, 1, 0, n-1)

	// 연산 처리
	for q := 0; q < m; q++ {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			// 구간 갱신: [l, r]에 v를 더한다 (1-indexed → 0-indexed)
			var l, r int
			var v int64
			fmt.Fscan(reader, &l, &r, &v)
			updateRange(1, 0, n-1, l-1, r-1, v)
		} else {
			// 구간 합 질의 (1-indexed → 0-indexed)
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, query(1, 0, n-1, l-1, r-1))
		}
	}
}
