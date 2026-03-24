package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 머지 소트 트리: 각 노드가 정렬된 배열을 저장하는 세그먼트 트리
var tree [][]int

// buildTree는 머지 소트 트리를 구성한다.
func buildTree(arr []int, node, start, end int) {
	if start == end {
		tree[node] = []int{arr[start]}
		return
	}
	mid := (start + end) / 2
	buildTree(arr, 2*node, start, mid)
	buildTree(arr, 2*node+1, mid+1, end)
	// 두 자식의 정렬된 배열을 병합
	tree[node] = mergeSorted(tree[2*node], tree[2*node+1])
}

// mergeSorted는 두 정렬된 배열을 병합한다.
func mergeSorted(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// countLessOrEqual은 구간 [l, r]에서 val 이하인 원소의 개수를 반환한다.
func countLessOrEqual(node, start, end, l, r, val int) int {
	if r < start || end < l {
		return 0
	}
	if l <= start && end <= r {
		// 이분 탐색으로 val 이하인 원소 개수 계산
		return sort.SearchInts(tree[node], val+1)
	}
	mid := (start + end) / 2
	return countLessOrEqual(2*node, start, mid, l, r, val) +
		countLessOrEqual(2*node+1, mid+1, end, l, r, val)
}

// kthMinimum은 구간 [l, r]에서 k번째로 작은 원소를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - queries: 쿼리 목록 (각 쿼리는 [l, r, k])
//
// [반환값]
//   - []int: 각 쿼리에 대한 결과
//
// [알고리즘 힌트]
//
//	머지 소트 트리 + 이분 탐색을 사용한다.
//	머지 소트 트리의 각 노드는 해당 구간의 정렬된 배열을 저장한다.
//	k번째 원소를 찾기 위해 값에 대해 이분 탐색하고,
//	각 후보 값에 대해 구간 내 해당 값 이하인 원소 수를 센다.
func kthMinimum(arr []int, queries [][]int) []int {
	n := len(arr)

	// 머지 소트 트리 구성
	tree = make([][]int, 4*n)
	buildTree(arr, 1, 0, n-1)

	// 정렬된 고유 값 목록 (이분 탐색용)
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	var results []int
	for _, q := range queries {
		l, r, k := q[0], q[1], q[2]

		// 값에 대해 이분 탐색: k번째로 작은 원소 찾기
		lo, hi := 0, n-1
		for lo < hi {
			mid := (lo + hi) / 2
			cnt := countLessOrEqual(1, 0, n-1, l, r, sorted[mid])
			if cnt >= k {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		results = append(results, sorted[lo])
	}

	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 쿼리 개수 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 쿼리 입력
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var l, r, k int
		fmt.Fscan(reader, &l, &r, &k)
		queries[i] = []int{l, r, k}
	}

	// 핵심 함수 호출
	results := kthMinimum(arr, queries)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
