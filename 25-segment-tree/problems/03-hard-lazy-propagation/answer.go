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

// build는 세그먼트 트리를 구축한다.
//
// [매개변수]
//   - arr: 원본 배열
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
//
// [알고리즘 힌트]
//
//	리프 노드에는 원본 배열 값을 저장하고,
//	내부 노드에는 자식 노드의 합을 저장한다.
func build(arr []int64, node, start, end int) {
	if start == end {
		tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	build(arr, 2*node, start, mid)
	build(arr, 2*node+1, mid+1, end)
	tree[node] = tree[2*node] + tree[2*node+1]
}

// pushDown은 lazy 값을 자식 노드에 전파한다.
//
// [매개변수]
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
//
// [알고리즘 힌트]
//
//	lazy 값이 0이 아니면 자식 노드에 전파한다.
//	각 자식의 tree 값에 lazy × 구간 크기를 더하고,
//	자식의 lazy 값에도 누적한 후 현재 lazy를 0으로 초기화한다.
func pushDown(node, start, end int) {
	if lazy[node] != 0 {
		mid := (start + end) / 2
		tree[2*node] += lazy[node] * int64(mid-start+1)
		lazy[2*node] += lazy[node]
		tree[2*node+1] += lazy[node] * int64(end-mid)
		lazy[2*node+1] += lazy[node]
		lazy[node] = 0
	}
}

// updateRange는 구간 [l, r]에 val을 더한다.
//
// [매개변수]
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
//   - l: 갱신 구간 왼쪽 끝
//   - r: 갱신 구간 오른쪽 끝
//   - val: 더할 값
//
// [알고리즘 힌트]
//
//	구간이 겹치지 않으면 반환하고,
//	완전히 포함되면 lazy 값을 저장한다.
//	부분적으로 겹치면 lazy를 전파한 후 자식에 재귀한다.
func updateRange(node, start, end, l, r int, val int64) {
	if r < start || end < l {
		return
	}
	if l <= start && end <= r {
		tree[node] += val * int64(end-start+1)
		lazy[node] += val
		return
	}
	pushDown(node, start, end)
	mid := (start + end) / 2
	updateRange(2*node, start, mid, l, r, val)
	updateRange(2*node+1, mid+1, end, l, r, val)
	tree[node] = tree[2*node] + tree[2*node+1]
}

// query는 구간 [l, r]의 합을 반환한다.
//
// [매개변수]
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
//   - l: 질의 구간 왼쪽 끝
//   - r: 질의 구간 오른쪽 끝
//
// [반환값]
//   - int64: 구간 [l, r]의 합
//
// [알고리즘 힌트]
//
//	구간이 겹치지 않으면 0을 반환하고,
//	완전히 포함되면 현재 노드 값을 반환한다.
//	부분적으로 겹치면 lazy를 전파한 후 자식에 재귀하여 합산한다.
func query(node, start, end, l, r int) int64 {
	if r < start || end < l {
		return 0
	}
	if l <= start && end <= r {
		return tree[node]
	}
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
			var l, r int
			var v int64
			fmt.Fscan(reader, &l, &r, &v)
			updateRange(1, 0, n-1, l-1, r-1, v)
		} else {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, query(1, 0, n-1, l-1, r-1))
		}
	}
}
