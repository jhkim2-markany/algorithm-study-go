package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Mo's Algorithm으로 구간 내 서로 다른 수의 개수를 구한다
// 시간 복잡도: O((N + Q)√N)

// Query는 쿼리 정보를 저장하는 구조체이다
type Query struct {
	l, r, idx int // 구간 [l, r], 원래 쿼리 인덱스
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 입력: 쿼리
	var q int
	fmt.Fscan(reader, &q)

	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].l, &queries[i].r)
		queries[i].l-- // 0-indexed로 변환
		queries[i].r--
		queries[i].idx = i
	}

	// 블록 크기 결정
	b := int(math.Ceil(math.Sqrt(float64(n))))
	if b == 0 {
		b = 1
	}

	// Mo's Algorithm: 쿼리를 (L/B, R) 기준으로 정렬한다
	sort.Slice(queries, func(i, j int) bool {
		bi, bj := queries[i].l/b, queries[j].l/b
		if bi != bj {
			return bi < bj
		}
		// 같은 블록이면 R 기준 정렬 (홀짝 최적화)
		if bi%2 == 0 {
			return queries[i].r < queries[j].r
		}
		return queries[i].r > queries[j].r
	})

	// 빈도 배열과 서로 다른 수의 개수
	cnt := make(map[int]int)
	distinct := 0
	ans := make([]int, q)

	// 원소 추가 함수
	add := func(val int) {
		cnt[val]++
		if cnt[val] == 1 {
			distinct++
		}
	}

	// 원소 제거 함수
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

		// 구간을 확장/축소하여 [l, r]로 맞춘다
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

	// 원래 쿼리 순서대로 출력
	for _, v := range ans {
		fmt.Fprintln(writer, v)
	}
}
