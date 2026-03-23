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
func build(arr []int64, node, start, end int) {
	// 여기에 코드를 작성하세요
}

// pushDown은 lazy 값을 자식 노드에 전파한다.
//
// [매개변수]
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
func pushDown(node, start, end int) {
	// 여기에 코드를 작성하세요
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
func updateRange(node, start, end, l, r int, val int64) {
	// 여기에 코드를 작성하세요
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
func query(node, start, end, l, r int) int64 {
	// 여기에 코드를 작성하세요
	return 0
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
