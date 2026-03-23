package main

import (
	"bufio"
	"fmt"
	"os"
)

// cutTree는 나무들을 높이 h로 잘랐을 때 m 이상의 나무를 얻을 수 있는
// 최대 절단 높이를 반환한다.
//
// [매개변수]
//   - trees: 각 나무의 높이 배열
//   - m: 필요한 나무 길이의 합
//
// [반환값]
//   - int: 조건을 만족하는 최대 절단 높이
//
// [알고리즘 힌트]
//
//	파라메트릭 서치: 절단 높이를 이진 탐색으로 찾는다.
//	결정 함수: 높이 mid로 잘랐을 때 얻는 나무 총합 >= m인가?
//	조건 만족 시 lo = mid + 1 (더 높은 높이 탐색),
//	불만족 시 hi = mid - 1 (높이를 낮춤).
//
//	시간복잡도: O(N log H), H는 최대 나무 높이
func cutTree(trees []int, m int) int {
	maxH := 0
	for _, h := range trees {
		if h > maxH {
			maxH = h
		}
	}

	lo, hi := 0, maxH
	result := 0

	for lo <= hi {
		mid := (lo + hi) / 2
		total := 0
		for _, h := range trees {
			if h > mid {
				total += h - mid
			}
		}
		if total >= m {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 나무 수 N과 필요한 나무 길이 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 각 나무의 높이 입력
	trees := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &trees[i])
	}

	// 핵심 함수 호출
	result := cutTree(trees, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
