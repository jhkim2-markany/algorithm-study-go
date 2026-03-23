package main

import (
	"bufio"
	"fmt"
	"os"
)

// countSubsetSum은 배열의 부분집합 중 합이 target인
// 부분집합의 개수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열 (원소는 양수, 음수, 0 모두 가능)
//   - target: 부분집합의 합이 되어야 하는 목표값
//
// [반환값]
//   - int: 합이 target인 부분집합의 개수 (빈 집합 제외)
//
// [알고리즘 힌트]
//
//	비트마스크를 사용하여 모든 부분집합을 열거한다.
//	N개 원소의 부분집합은 총 2^N개이며, 각 부분집합을
//	0부터 2^N-1까지의 정수(mask)로 표현할 수 있다.
//	mask의 i번째 비트가 1이면 arr[i]를 선택한 것이다.
//	빈 집합(mask=0)은 제외하므로 mask는 1부터 시작한다.
//
//	시간복잡도: O(N × 2^N), N ≤ 20 정도에서 사용 가능
//
//	예시: arr=[1, 2, 3], target=3
//	  부분집합: {1,2}=3 ✓, {3}=3 ✓ → 답: 2
func countSubsetSum(arr []int, target int) int {
	n := len(arr)
	count := 0

	// 비트마스크로 모든 부분집합을 탐색 (빈 집합 제외)
	for mask := 1; mask < (1 << n); mask++ {
		sum := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += arr[i]
			}
		}
		if sum == target {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 목표 합 입력
	var n, s int
	fmt.Fscan(reader, &n, &s)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	count := countSubsetSum(arr, s)

	// 결과 출력
	fmt.Fprintln(writer, count)
}
