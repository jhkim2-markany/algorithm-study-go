package main

import (
	"fmt"
	"sort"
)

// 중간에서 만나기 (Meet in the Middle) - 부분집합 합 존재 판정
// 시간 복잡도: O(2^(N/2) * N)
// 공간 복잡도: O(2^(N/2))

// enumSubsetSums는 arr의 모든 부분집합 합을 열거하여 반환한다
func enumSubsetSums(arr []int) []int {
	n := len(arr)
	size := 1 << n // 2^n
	sums := make([]int, 0, size)
	for mask := 0; mask < size; mask++ {
		s := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				s += arr[i] // 비트가 켜진 원소를 합산
			}
		}
		sums = append(sums, s)
	}
	return sums
}

// meetInTheMiddle은 arr에서 합이 target인 부분집합이 존재하는지 판정한다
func meetInTheMiddle(arr []int, target int) bool {
	n := len(arr)
	if n == 0 {
		return target == 0
	}

	// 1단계: 배열을 반으로 분할
	half := n / 2
	left := arr[:half]
	right := arr[half:]

	// 2단계: 각 절반의 모든 부분집합 합 열거
	sumA := enumSubsetSums(left)
	sumB := enumSubsetSums(right)

	// 3단계: sumB를 정렬하여 이분 탐색 준비
	sort.Ints(sumB)

	// 4단계: sumA의 각 원소에 대해 sumB에서 보완값 탐색
	for _, a := range sumA {
		need := target - a
		// 이분 탐색으로 need가 sumB에 존재하는지 확인
		idx := sort.SearchInts(sumB, need)
		if idx < len(sumB) && sumB[idx] == need {
			return true // 합이 target인 부분집합 발견
		}
	}
	return false
}

func main() {
	// 예시: 배열 [3, 1, 4, 1, 5, 9]에서 합이 10인 부분집합 찾기
	arr := []int{3, 1, 4, 1, 5, 9}
	target := 10

	fmt.Printf("배열: %v\n", arr)
	fmt.Printf("목표 합: %d\n", target)

	if meetInTheMiddle(arr, target) {
		fmt.Println("결과: 합이 목표값인 부분집합이 존재합니다")
	} else {
		fmt.Println("결과: 합이 목표값인 부분집합이 존재하지 않습니다")
	}

	// 추가 예시
	fmt.Println()
	arr2 := []int{1, 2, 3, 4, 5}
	target2 := 100
	fmt.Printf("배열: %v\n", arr2)
	fmt.Printf("목표 합: %d\n", target2)

	if meetInTheMiddle(arr2, target2) {
		fmt.Println("결과: 합이 목표값인 부분집합이 존재합니다")
	} else {
		fmt.Println("결과: 합이 목표값인 부분집합이 존재하지 않습니다")
	}
}
