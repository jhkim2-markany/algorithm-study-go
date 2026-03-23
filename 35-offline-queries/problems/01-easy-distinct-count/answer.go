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

// distinctCount는 각 쿼리 구간 [l, r]에서 서로 다른 수의 개수를 반환한다.
//
// [매개변수]
//   - n: 배열의 크기
//   - arr: 1-indexed 배열 (길이 n+1, arr[0]은 미사용)
//   - queries: 쿼리 목록 (각 쿼리는 left, right, index를 포함)
//
// [반환값]
//   - []int: 각 쿼리에 대한 서로 다른 수의 개수 (원래 순서)
//
// [알고리즘 힌트]
//
//	Mo's 알고리즘을 사용한다.
//	블록 크기 √N으로 쿼리를 정렬한 뒤 포인터를 이동하며 처리한다.
//	빈도 배열로 서로 다른 수의 개수를 관리한다.
//	시간복잡도: O((N+Q) * √N)
func distinctCount(n int, arr []int, queries []Query) []int {
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
		return queries[i].right < queries[j].right
	})

	maxVal := 0
	for i := 1; i <= n; i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	cnt := make([]int, maxVal+1)
	distinct := 0

	add := func(val int) {
		cnt[val]++
		if cnt[val] == 1 {
			distinct++
		}
	}
	remove := func(val int) {
		cnt[val]--
		if cnt[val] == 0 {
			distinct--
		}
	}

	answers := make([]int, len(queries))
	curL, curR := 1, 0

	for _, query := range queries {
		l, r := query.left, query.right

		for curR < r {
			curR++
			add(arr[curR])
		}
		for curL > l {
			curL--
			add(arr[curL])
		}
		for curR > r {
			remove(arr[curR])
			curR--
		}
		for curL < l {
			remove(arr[curL])
			curL++
		}

		answers[query.index] = distinct
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

	answers := distinctCount(n, arr, queries)
	for _, ans := range answers {
		fmt.Fprintln(writer, ans)
	}
}
