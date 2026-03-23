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
//   - houses: 집의 좌표 배열
//   - c: 설치할 공유기 수
//
// [반환값]
//   - int: 인접 공유기 간 최소 거리의 최댓값
func installRouter(houses []int, c int) int {
	// 여기에 코드를 작성하세요
	return 0
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
