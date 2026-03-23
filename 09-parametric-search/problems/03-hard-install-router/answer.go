package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// installRouter는 집 좌표 배열에서 c개의 공유기를 설치할 때
// 인접 공유기 간 최소 거리의 최댓값을 반환한다.
//
// [매개변수]
//   - houses: 집의 좌표 배열 (오름차순 정렬됨)
//   - c: 설치할 공유기 수
//
// [반환값]
//   - int: 인접 공유기 간 최소 거리의 최댓값
//
// [알고리즘 힌트]
//
//	파라메트릭 서치: 최소 거리를 이진 탐색으로 찾는다.
//	결정 함수: 최소 거리를 mid 이상으로 유지하며 c개 이상 설치 가능한가?
//	첫 번째 집에 무조건 설치하고, 이후 거리가 mid 이상인 집에만 설치한다.
//	조건 만족 시 lo = mid + 1, 불만족 시 hi = mid - 1.
//
//	시간복잡도: O(N log D), D는 최대 좌표 차이
func installRouter(houses []int, c int) int {
	n := len(houses)
	lo, hi := 1, houses[n-1]-houses[0]
	result := 0

	for lo <= hi {
		mid := (lo + hi) / 2
		count := 1
		lastPos := houses[0]

		for i := 1; i < n; i++ {
			if houses[i]-lastPos >= mid {
				count++
				lastPos = houses[i]
			}
		}

		if count >= c {
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

	// 집의 수 N과 공유기 수 C 입력
	var n, c int
	fmt.Fscan(reader, &n, &c)

	// 각 집의 좌표 입력
	houses := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &houses[i])
	}

	// 좌표 정렬
	sort.Ints(houses)

	// 핵심 함수 호출
	result := installRouter(houses, c)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
