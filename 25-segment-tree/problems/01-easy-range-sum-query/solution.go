package main

import (
	"bufio"
	"fmt"
	"os"
)

// 세그먼트 트리 배열
var tree []int64
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

// update 함수는 인덱스 idx의 값을 val로 변경한다
func update(node, start, end, idx int, val int64) {
	if start == end {
		// 리프 노드에 도달하면 값 갱신
		tree[node] = val
		return
	}
	mid := (start + end) / 2
	if idx <= mid {
		update(2*node, start, mid, idx, val)
	} else {
		update(2*node+1, mid+1, end, idx, val)
	}
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
	// 부분적으로 겹치는 경우
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

	// 세그먼트 트리 구축
	tree = make([]int64, 4*n)
	build(arr, 1, 0, n-1)

	// 연산 처리
	for q := 0; q < m; q++ {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			// 점 갱신: i번째 원소를 v로 변경 (1-indexed → 0-indexed)
			var i int
			var v int64
			fmt.Fscan(reader, &i, &v)
			update(1, 0, n-1, i-1, v)
		} else {
			// 구간 합 질의 (1-indexed → 0-indexed)
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, query(1, 0, n-1, l-1, r-1))
		}
	}
}
