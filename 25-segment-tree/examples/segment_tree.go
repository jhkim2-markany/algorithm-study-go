package main

import "fmt"

// 세그먼트 트리 - 구간 합 질의와 점 갱신
// 시간 복잡도: 구축 O(N), 질의 O(log N), 갱신 O(log N)
// 공간 복잡도: O(N)

// SegmentTree 구조체는 구간 합을 관리하는 세그먼트 트리이다
type SegmentTree struct {
	tree []int // 트리 배열
	n    int   // 원본 배열의 크기
}

// NewSegmentTree 함수는 주어진 배열로 세그먼트 트리를 구축한다
func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	st := &SegmentTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	// 트리 구축
	st.build(arr, 1, 0, n-1)
	return st
}

// build 함수는 재귀적으로 세그먼트 트리를 구축한다
func (st *SegmentTree) build(arr []int, node, start, end int) {
	if start == end {
		// 리프 노드: 원본 배열의 값을 저장
		st.tree[node] = arr[start]
		return
	}

	mid := (start + end) / 2
	// 왼쪽 자식 구축
	st.build(arr, 2*node, start, mid)
	// 오른쪽 자식 구축
	st.build(arr, 2*node+1, mid+1, end)
	// 내부 노드: 두 자식의 합을 저장
	st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
}

// Query 함수는 구간 [l, r]의 합을 반환한다
func (st *SegmentTree) Query(l, r int) int {
	return st.query(1, 0, st.n-1, l, r)
}

func (st *SegmentTree) query(node, start, end, l, r int) int {
	// 현재 구간이 질의 구간과 전혀 겹치지 않는 경우
	if r < start || end < l {
		return 0 // 합의 항등원
	}
	// 현재 구간이 질의 구간에 완전히 포함되는 경우
	if l <= start && end <= r {
		return st.tree[node]
	}
	// 부분적으로 겹치는 경우: 양쪽 자식에 재귀 질의
	mid := (start + end) / 2
	leftSum := st.query(2*node, start, mid, l, r)
	rightSum := st.query(2*node+1, mid+1, end, l, r)
	return leftSum + rightSum
}

// Update 함수는 인덱스 idx의 값을 val로 변경한다
func (st *SegmentTree) Update(idx, val int) {
	st.update(1, 0, st.n-1, idx, val)
}

func (st *SegmentTree) update(node, start, end, idx, val int) {
	if start == end {
		// 리프 노드에 도달: 값 갱신
		st.tree[node] = val
		return
	}

	mid := (start + end) / 2
	if idx <= mid {
		// 갱신할 인덱스가 왼쪽 구간에 있는 경우
		st.update(2*node, start, mid, idx, val)
	} else {
		// 갱신할 인덱스가 오른쪽 구간에 있는 경우
		st.update(2*node+1, mid+1, end, idx, val)
	}
	// 부모 노드 값 재계산
	st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
}

func main() {
	// 예제 배열
	arr := []int{1, 3, 5, 7, 9, 11}
	fmt.Println("원본 배열:", arr)

	// 세그먼트 트리 구축
	st := NewSegmentTree(arr)

	// 구간 합 질의
	fmt.Printf("구간 [1, 3]의 합: %d\n", st.Query(1, 3)) // 3+5+7 = 15
	fmt.Printf("구간 [0, 5]의 합: %d\n", st.Query(0, 5)) // 1+3+5+7+9+11 = 36
	fmt.Printf("구간 [2, 4]의 합: %d\n", st.Query(2, 4)) // 5+7+9 = 21

	// 점 갱신: 인덱스 2의 값을 10으로 변경
	st.Update(2, 10)
	fmt.Println("\n인덱스 2의 값을 10으로 변경 후:")
	fmt.Printf("구간 [1, 3]의 합: %d\n", st.Query(1, 3)) // 3+10+7 = 20
	fmt.Printf("구간 [0, 5]의 합: %d\n", st.Query(0, 5)) // 1+3+10+7+9+11 = 41
}
