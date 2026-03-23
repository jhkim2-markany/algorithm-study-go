package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// enumAllSums는 arr의 모든 부분집합 합을 열거한다
func enumAllSums(arr []int) []int {
	n := len(arr)
	sums := make([]int, 0, 1<<n)
	for mask := 0; mask < (1 << n); mask++ {
		s := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				s += arr[i] // 선택된 원소 합산
			}
		}
		sums = append(sums, s)
	}
	return sums
}

// abs는 정수의 절댓값을 반환한다
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기와 목표값
	var n, s int
	fmt.Fscan(reader, &n, &s)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 배열을 반으로 분할
	half := n / 2
	left := arr[:half]
	right := arr[half:]

	// 각 절반의 부분집합 합 열거
	sumA := enumAllSums(left)
	sumB := enumAllSums(right)

	// sumB를 정렬하여 이분 탐색 준비
	sort.Ints(sumB)

	// 가장 가까운 합 탐색
	bestDiff := int(1e18)
	bestSum := 0

	for _, a := range sumA {
		need := s - a
		// sumB에서 need 이상인 첫 위치를 이분 탐색
		idx := sort.SearchInts(sumB, need)

		// idx 위치의 값 확인 (need 이상인 가장 작은 값)
		if idx < len(sumB) {
			total := a + sumB[idx]
			if abs(total-s) < bestDiff {
				bestDiff = abs(total - s)
				bestSum = total
			}
		}

		// idx-1 위치의 값 확인 (need 미만인 가장 큰 값)
		if idx > 0 {
			total := a + sumB[idx-1]
			if abs(total-s) < bestDiff {
				bestDiff = abs(total - s)
				bestSum = total
			}
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, bestSum)
}
