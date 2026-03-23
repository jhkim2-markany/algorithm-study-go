package main

import "fmt"

// 이진 탐색 - 정렬된 배열에서 값을 찾는 기본 예시
// 시간 복잡도: O(log N)
// 공간 복잡도: O(1)

// binarySearch 함수는 정렬된 배열에서 target의 인덱스를 반환한다.
// 값이 없으면 -1을 반환한다.
func binarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if arr[mid] == target {
			// 목표 값을 찾은 경우
			return mid
		} else if arr[mid] < target {
			// 목표 값이 오른쪽 절반에 있음
			lo = mid + 1
		} else {
			// 목표 값이 왼쪽 절반에 있음
			hi = mid - 1
		}
	}

	// 값을 찾지 못한 경우
	return -1
}

// lowerBound 함수는 target 이상인 첫 번째 위치를 반환한다.
// 모든 원소가 target보다 작으면 len(arr)을 반환한다.
func lowerBound(arr []int, target int) int {
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := (lo + hi) / 2

		if arr[mid] < target {
			// 아직 target보다 작으므로 오른쪽으로 이동
			lo = mid + 1
		} else {
			// target 이상인 위치를 찾았으므로 왼쪽 범위 축소
			hi = mid
		}
	}

	return lo
}

// upperBound 함수는 target을 초과하는 첫 번째 위치를 반환한다.
// 모든 원소가 target 이하이면 len(arr)을 반환한다.
func upperBound(arr []int, target int) int {
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := (lo + hi) / 2

		if arr[mid] <= target {
			// target 이하이므로 오른쪽으로 이동
			lo = mid + 1
		} else {
			// target 초과인 위치를 찾았으므로 왼쪽 범위 축소
			hi = mid
		}
	}

	return lo
}

func main() {
	// 정렬된 배열 준비
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}

	// 기본 이진 탐색
	fmt.Println("=== 기본 이진 탐색 ===")
	fmt.Printf("배열: %v\n", arr)

	target := 7
	idx := binarySearch(arr, target)
	fmt.Printf("값 %d의 인덱스: %d\n", target, idx)

	target = 6
	idx = binarySearch(arr, target)
	fmt.Printf("값 %d의 인덱스: %d (없음)\n", target, idx)

	// Lower Bound / Upper Bound
	fmt.Println("\n=== Lower Bound / Upper Bound ===")
	arr2 := []int{1, 2, 2, 2, 3, 3, 5, 7}
	fmt.Printf("배열: %v\n", arr2)

	target = 2
	lb := lowerBound(arr2, target)
	ub := upperBound(arr2, target)
	fmt.Printf("값 %d: lower_bound=%d, upper_bound=%d\n", target, lb, ub)
	fmt.Printf("값 %d의 개수: %d\n", target, ub-lb)

	target = 3
	lb = lowerBound(arr2, target)
	ub = upperBound(arr2, target)
	fmt.Printf("값 %d: lower_bound=%d, upper_bound=%d\n", target, lb, ub)
	fmt.Printf("값 %d의 개수: %d\n", target, ub-lb)

	// 존재하지 않는 값
	target = 4
	lb = lowerBound(arr2, target)
	ub = upperBound(arr2, target)
	fmt.Printf("값 %d: lower_bound=%d, upper_bound=%d (없음, 개수=%d)\n", target, lb, ub, ub-lb)
}
