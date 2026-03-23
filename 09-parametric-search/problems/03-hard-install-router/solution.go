package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

	// 좌표를 오름차순으로 정렬한다
	sort.Ints(houses)

	// 파라메트릭 서치: 최소 거리의 최댓값을 이진 탐색으로 찾는다
	// 탐색 범위: 1부터 (최대 좌표 - 최소 좌표)까지
	lo, hi := 1, houses[n-1]-houses[0]
	result := 0

	for lo <= hi {
		mid := (lo + hi) / 2

		// 결정 함수: 최소 거리를 mid 이상으로 유지하며 C개 이상 설치할 수 있는가?
		count := 1           // 첫 번째 집에 무조건 설치
		lastPos := houses[0] // 마지막으로 설치한 위치

		for i := 1; i < n; i++ {
			if houses[i]-lastPos >= mid {
				// 현재 집과 마지막 설치 위치의 거리가 mid 이상이면 설치
				count++
				lastPos = houses[i]
			}
		}

		if count >= c {
			// 조건 만족: 더 큰 거리도 가능한지 확인
			result = mid
			lo = mid + 1
		} else {
			// 조건 불만족: 거리를 줄인다
			hi = mid - 1
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, result)
}
