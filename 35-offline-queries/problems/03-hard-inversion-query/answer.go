package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Query는 구간 쿼리 정보를 저장한다.
type Query struct {
	left  int
	right int
	index int
}

// BIT (Binary Indexed Tree) 관련 변수
var bit []int
var bitSize int

func bitUpdate(idx, val int) {
	for ; idx <= bitSize; idx += idx & (-idx) {
		bit[idx] += val
	}
}

func bitQuery(idx int) int {
	sum := 0
	for ; idx > 0; idx -= idx & (-idx) {
		sum += bit[idx]
	}
	return sum
}

// inversionQuery는 각 쿼리 구간 [l, r]에서 역전 쌍의 개수를 반환한다.
//
// [매개변수]
//   - n: 배열의 크기
//   - arr: 1-indexed 배열 (길이 n+1, arr[0]은 미사용)
//   - queries: 쿼리 목록 (각 쿼리는 left, right, index를 포함)
//
// [반환값]
//   - []int: 각 쿼리에 대한 역전 쌍의 개수 (원래 순서)
//
// [알고리즘 힌트]
//
//	Mo's 알고리즘 + BIT를 사용한다.
//	오른쪽 추가: 기존 원소 중 새 원소보다 큰 값의 개수가 역전 쌍 증가분.
//	왼쪽 추가: 기존 원소 중 새 원소보다 작은 값의 개수가 역전 쌍 증가분.
//	홀짝 최적화로 포인터 이동을 줄인다.
//	시간복잡도: O((N+Q) * √N * log(maxVal))
func inversionQuery(n int, arr []int, queries []Query) []int {
	blockSize := int(math.Sqrt(float64(n)))
	if blockSize == 0 {
		blockSize = 1
	}

	sort.Slice(queries, func(i, j int) bool {
		bi := queries[i].left / blockSize
		bj := queries[j].left / blockSize
		if bi != bj {
			return bi < bj
		}
		if bi%2 == 1 {
			return queries[i].right > queries[j].right
		}
		return queries[i].right < queries[j].right
	})

	maxVal := 0
	for i := 1; i <= n; i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	bitSize = maxVal
	bit = make([]int, bitSize+1)
	inversions := 0

	addRight := func(val int) {
		total := bitQuery(bitSize)
		lessOrEqual := bitQuery(val)
		inversions += total - lessOrEqual
		bitUpdate(val, 1)
	}
	removeRight := func(val int) {
		bitUpdate(val, -1)
		total := bitQuery(bitSize)
		lessOrEqual := bitQuery(val)
		inversions -= total - lessOrEqual
	}
	addLeft := func(val int) {
		less := bitQuery(val - 1)
		inversions += less
		bitUpdate(val, 1)
	}
	removeLeft := func(val int) {
		bitUpdate(val, -1)
		less := bitQuery(val - 1)
		inversions -= less
	}

	answers := make([]int, len(queries))
	curL, curR := 1, 0

	for _, qr := range queries {
		l, r := qr.left, qr.right

		for curR < r {
			curR++
			addRight(arr[curR])
		}
		for curL > l {
			curL--
			addLeft(arr[curL])
		}
		for curR > r {
			removeRight(arr[curR])
			curR--
		}
		for curL < l {
			removeLeft(arr[curL])
			curL++
		}

		answers[qr.index] = inversions
	}

	return answers
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].left, &queries[i].right)
		queries[i].index = i
	}

	answers := inversionQuery(n, arr, queries)
	for _, ans := range answers {
		fmt.Fprintln(writer, ans)
	}
}
