package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// 쿼리 구조체: 구간 정보와 원래 순서를 저장한다
type Query struct {
	left  int
	right int
	index int
}

var (
	blockSize int
	cnt       []int
	distinct  int
)

// 원소를 구간에 추가한다
func add(val int) {
	cnt[val]++
	if cnt[val] == 1 {
		distinct++
	}
}

// 원소를 구간에서 제거한다
func remove(val int) {
	cnt[val]--
	if cnt[val] == 0 {
		distinct--
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 쿼리 수 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 배열 입력 (1-indexed로 사용)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 쿼리 입력
	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].left, &queries[i].right)
		queries[i].index = i
	}

	// 블록 크기 설정
	blockSize = int(math.Sqrt(float64(n)))
	if blockSize == 0 {
		blockSize = 1
	}

	// Mo's 알고리즘 정렬: 블록 번호 → R 오름차순
	sort.Slice(queries, func(i, j int) bool {
		bi := queries[i].left / blockSize
		bj := queries[j].left / blockSize
		if bi != bj {
			return bi < bj
		}
		return queries[i].right < queries[j].right
	})

	// 빈도 배열 초기화
	maxVal := 0
	for i := 1; i <= n; i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	cnt = make([]int, maxVal+1)
	distinct = 0

	// 결과 배열
	answers := make([]int, q)

	// 현재 구간 초기화
	curL, curR := 1, 0

	// Mo's 알고리즘으로 쿼리 처리
	for _, query := range queries {
		l, r := query.left, query.right

		// 포인터를 목표 구간으로 이동한다
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

		// 현재 구간의 서로 다른 수 개수를 기록한다
		answers[query.index] = distinct
	}

	// 원래 순서대로 결과 출력
	for _, ans := range answers {
		fmt.Fprintln(writer, ans)
	}
}
