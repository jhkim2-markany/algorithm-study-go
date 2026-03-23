package main

import (
	"bufio"
	"fmt"
	"os"
)

// findSumPair는 배열에서 합이 target인 두 수를 찾아 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - target: 두 수의 합이 되어야 하는 목표값
//
// [반환값]
//   - int, int: 합이 target인 두 수 (첫 번째로 발견된 쌍)
//   - bool: 조건을 만족하는 쌍을 찾았으면 true, 없으면 false
//
// [알고리즘 힌트]
//
//	브루트포스 접근: 2중 반복문으로 모든 (i, j) 쌍을 탐색한다 (i < j).
//	arr[i] + arr[j] == target이면 해당 쌍을 반환한다.
//	시간복잡도: O(N²), 공간복잡도: O(1)
//
//	예시: arr=[1, 3, 5, 7], target=8
//	  → (1,7) 또는 (3,5) 중 먼저 발견되는 쌍 반환
func findSumPair(arr []int, target int) (int, int, bool) {
	// 2중 반복문으로 모든 쌍을 탐색
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return arr[i], arr[j], true
			}
		}
	}
	return 0, 0, false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 목표 합 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	a, b, found := findSumPair(arr, m)

	// 결과 출력
	if found {
		fmt.Fprintf(writer, "%d %d\n", a, b)
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
