package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// sherlockAndMinimax는 범위 [p, q]에서 min(|arr[i] - M|)을 최대화하는 M을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - p: 범위 시작
//   - q: 범위 끝
//
// [반환값]
//   - int: 조건을 만족하는 M
//
// [알고리즘 힌트]
//
//	정렬 후 인접 원소의 중간점과 범위 양 끝점을 후보로 평가한다.
//	각 후보에 대해 이진 탐색으로 가장 가까운 원소와의 거리를 계산한다.
func sherlockAndMinimax(arr []int, p int, q int) int {
	// 배열 정렬
	sort.Ints(arr)
	n := len(arr)

	// 후보 M에 대해 min(|arr[i] - M|)을 계산하는 함수
	minDist := func(m int) int {
		// 이진 탐색으로 m에 가장 가까운 원소 찾기
		idx := sort.SearchInts(arr, m)
		dist := math.MaxInt64
		if idx < n {
			d := arr[idx] - m
			if d < dist {
				dist = d
			}
		}
		if idx > 0 {
			d := m - arr[idx-1]
			if d < dist {
				dist = d
			}
		}
		return dist
	}

	bestM := p
	bestDist := minDist(p)

	// 범위 끝점 Q 확인
	d := minDist(q)
	if d > bestDist || (d == bestDist && q < bestM) {
		bestDist = d
		bestM = q
	}

	// 인접한 두 원소의 중간점을 후보로 평가
	for i := 0; i < n-1; i++ {
		mid := (arr[i] + arr[i+1]) / 2
		// 중간점이 범위 [p, q]에 포함되는지 확인
		if mid >= p && mid <= q {
			d := minDist(mid)
			if d > bestDist || (d == bestDist && mid < bestM) {
				bestDist = d
				bestM = mid
			}
		}
	}

	return bestM
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 범위 입력
	var p, q int
	fmt.Fscan(reader, &p, &q)

	// 핵심 함수 호출 및 결과 출력
	result := sherlockAndMinimax(arr, p, q)
	fmt.Fprintln(writer, result)
}
