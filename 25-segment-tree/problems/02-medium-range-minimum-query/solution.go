package main

import (
	"bufio"
	"fmt"
	"os"
)

// 세그먼트 트리 배열 (구간 최솟값)
var tree []int64
var n int

// min64는 두 int64 값 중 작은 값을 반환한다
func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// build는 구간 최솟값 세그먼트 트리를 구축한다.
//
// [매개변수]
//   - arr: 원본 배열
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
func build(arr []int64, node, start, end int) {
	// 여기에 코드를 작성하세요
}

// update는 인덱스 idx의 값을 val로 변경한다.
//
// [매개변수]
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
//   - idx: 변경할 인덱스
//   - val: 새로운 값
func update(node, start, end, idx int, val int64) {
	// 여기에 코드를 작성하세요
}

// query는 구간 [l, r]의 최솟값을 반환한다.
//
// [매개변수]
//   - node: 현재 노드 번호
//   - start: 구간 시작 인덱스
//   - end: 구간 끝 인덱스
//   - l: 질의 구간 왼쪽 끝
//   - r: 질의 구간 오른쪽 끝
//
// [반환값]
//   - int64: 구간 [l, r]의 최솟값
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

	// 세그먼트 트리 구축
	tree = make([]int64, 4*n)
	build(arr, 1, 0, n-1)

	// 연산 처리
	for q := 0; q < m; q++ {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			var i int
			var v int64
			fmt.Fscan(reader, &i, &v)
			update(1, 0, n-1, i-1, v)
		} else {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, query(1, 0, n-1, l-1, r-1))
		}
	}
}
