package main

import (
	"fmt"
	"math"
	"sort"
)

// 오프라인 쿼리 (Offline Query) - Mo's 알고리즘
// 배열의 구간 쿼리를 정렬하여 효율적으로 처리한다.
// 시간 복잡도: O((N + Q) × √N)
// 공간 복잡도: O(N + Q)

// Query 구조체: 구간 쿼리 정보를 저장한다
type Query struct {
	left  int // 구간 왼쪽 끝 (0-indexed)
	right int // 구간 오른쪽 끝 (0-indexed, 포함)
	index int // 원래 쿼리 순서 (답 복원용)
}

// 블록 크기 (√N)
var blockSize int

// Mo's 알고리즘용 쿼리 정렬
// 1차: L이 속하는 블록 번호 기준 오름차순
// 2차: 같은 블록 내에서 R 기준 오름차순
func sortQueries(queries []Query) {
	sort.Slice(queries, func(i, j int) bool {
		blockI := queries[i].left / blockSize
		blockJ := queries[j].left / blockSize
		if blockI != blockJ {
			return blockI < blockJ
		}
		return queries[i].right < queries[j].right
	})
}

// 구간 내 서로 다른 원소의 개수를 구하는 예제
// cnt[v]: 현재 구간에서 값 v의 등장 횟수
// distinct: 현재 구간의 서로 다른 원소 개수
var cnt []int
var distinct int

// 원소를 현재 구간에 추가한다
func add(val int) {
	cnt[val]++
	if cnt[val] == 1 {
		// 처음 등장하는 원소이면 서로 다른 수 증가
		distinct++
	}
}

// 원소를 현재 구간에서 제거한다
func remove(val int) {
	cnt[val]--
	if cnt[val] == 0 {
		// 더 이상 없는 원소이면 서로 다른 수 감소
		distinct--
	}
}

// Mo's 알고리즘으로 구간별 서로 다른 원소 개수를 구한다
func mosAlgorithm(arr []int, queries []Query) []int {
	n := len(arr)
	q := len(queries)

	// 블록 크기 설정: √N
	blockSize = int(math.Sqrt(float64(n)))
	if blockSize == 0 {
		blockSize = 1
	}

	// 쿼리를 Mo's 순서로 정렬한다
	sortQueries(queries)

	// 빈도 배열 초기화 (값의 범위에 맞게 설정)
	maxVal := 0
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	cnt = make([]int, maxVal+1)
	distinct = 0

	// 결과 배열
	answers := make([]int, q)

	// 현재 구간 [curL, curR]
	curL, curR := 0, -1

	for _, query := range queries {
		l, r := query.left, query.right

		// 오른쪽 포인터를 목표 위치까지 이동한다
		for curR < r {
			curR++
			add(arr[curR])
		}
		// 왼쪽 포인터를 목표 위치까지 이동한다
		for curL > l {
			curL--
			add(arr[curL])
		}
		// 오른쪽 포인터를 축소한다
		for curR > r {
			remove(arr[curR])
			curR--
		}
		// 왼쪽 포인터를 축소한다
		for curL < l {
			remove(arr[curL])
			curL++
		}

		// 현재 구간의 답을 기록한다
		answers[query.index] = distinct
	}

	return answers
}

func main() {
	// 예시 배열
	arr := []int{1, 2, 1, 3, 2, 1, 4, 3}
	fmt.Println("배열:", arr)
	fmt.Println()

	// 구간 쿼리 목록 (0-indexed)
	queries := []Query{
		{left: 0, right: 3, index: 0}, // [1, 2, 1, 3] → 서로 다른 수: 3
		{left: 1, right: 5, index: 1}, // [2, 1, 3, 2, 1] → 서로 다른 수: 3
		{left: 2, right: 7, index: 2}, // [1, 3, 2, 1, 4, 3] → 서로 다른 수: 4
		{left: 0, right: 7, index: 3}, // [1, 2, 1, 3, 2, 1, 4, 3] → 서로 다른 수: 4
		{left: 4, right: 6, index: 4}, // [2, 1, 4] → 서로 다른 수: 3
	}

	// Mo's 알고리즘으로 쿼리 처리
	answers := mosAlgorithm(arr, queries)

	// 결과 출력 (원래 쿼리 순서대로)
	fmt.Println("=== Mo's 알고리즘 결과 ===")
	queryInfos := [][2]int{{0, 3}, {1, 5}, {2, 7}, {0, 7}, {4, 6}}
	for i, ans := range answers {
		fmt.Printf("구간 [%d, %d]: 서로 다른 원소 개수 = %d\n",
			queryInfos[i][0], queryInfos[i][1], ans)
	}
}
