package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Query는 쿼리 정보를 저장하는 구조체이다
type Query struct {
	l, r, idx int
}

// mosDistinctCount는 Mo's Algorithm으로 각 구간 쿼리에 대해 서로 다른 수의 개수를 반환한다.
//
// [매개변수]
//   - a: 정수 배열 (0-indexed)
//   - queries: 쿼리 배열 (각 쿼리는 0-indexed 구간 [l, r]과 원래 인덱스)
//
// [반환값]
//   - []int: 각 쿼리에 대한 서로 다른 수의 개수 (원래 쿼리 순서)
//
// [알고리즘 힌트]
//   - 블록 크기 b = ceil(√n)으로 쿼리를 (L/b, R) 기준으로 정렬한다
//   - 홀짝 최적화: 같은 블록 내에서 짝수 블록은 R 오름차순, 홀수 블록은 R 내림차순
//   - 현재 구간 [curL, curR]을 확장/축소하며 원소를 추가/제거한다
//   - 빈도 맵(cnt)으로 서로 다른 수의 개수(distinct)를 관리한다
//   - 시간 복잡도: O((N + Q)√N)
func mosDistinctCount(a []int, queries []Query) []int {
	n := len(a)
	q := len(queries)

	b := int(math.Ceil(math.Sqrt(float64(n))))
	if b == 0 {
		b = 1
	}

	sort.Slice(queries, func(i, j int) bool {
		bi, bj := queries[i].l/b, queries[j].l/b
		if bi != bj {
			return bi < bj
		}
		if bi%2 == 0 {
			return queries[i].r < queries[j].r
		}
		return queries[i].r > queries[j].r
	})

	cnt := make(map[int]int)
	distinct := 0
	ans := make([]int, q)

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
			delete(cnt, val)
		}
	}

	curL, curR := 0, -1

	for _, qr := range queries {
		l, r := qr.l, qr.r

		for curR < r {
			curR++
			add(a[curR])
		}
		for curL > l {
			curL--
			add(a[curL])
		}
		for curR > r {
			remove(a[curR])
			curR--
		}
		for curL < l {
			remove(a[curL])
			curL++
		}

		ans[qr.idx] = distinct
	}

	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var q int
	fmt.Fscan(reader, &q)

	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].l, &queries[i].r)
		queries[i].l--
		queries[i].r--
		queries[i].idx = i
	}

	ans := mosDistinctCount(a, queries)

	for _, v := range ans {
		fmt.Fprintln(writer, v)
	}
}
