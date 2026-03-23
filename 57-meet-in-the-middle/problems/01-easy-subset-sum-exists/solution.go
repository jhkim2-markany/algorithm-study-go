package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// enumSums는 arr의 모든 부분집합 합을 열거한다
func enumSums(arr []int) []int {
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기와 목표 합
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
	sumA := enumSums(left)
	sumB := enumSums(right)

	// sumB를 정렬하여 이분 탐색 준비
	sort.Ints(sumB)

	// sumA의 각 원소에 대해 보완값을 이분 탐색
	found := false
	for _, a := range sumA {
		need := s - a
		idx := sort.SearchInts(sumB, need)
		if idx < len(sumB) && sumB[idx] == need {
			found = true
			break
		}
	}

	// 결과 출력
	if found {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
