package main

import "fmt"

// 분할 정복 기본 구현 - 병합 정렬과 최대 부분 배열 합
// 시간 복잡도: O(N log N)
// 공간 복잡도: O(N)

// mergeSort 함수는 배열을 병합 정렬로 정렬한다
func mergeSort(arr []int) []int {
	// 기저 조건: 길이가 1 이하이면 이미 정렬된 상태
	if len(arr) <= 1 {
		return arr
	}

	// 분할: 배열을 중간 지점에서 둘로 나눈다
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// 결합: 정렬된 두 배열을 합친다
	return merge(left, right)
}

// merge 함수는 정렬된 두 배열을 하나의 정렬된 배열로 합친다
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// 두 배열을 비교하며 작은 값부터 추가
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 남은 원소 추가
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// maxSubarraySum 함수는 분할 정복으로 최대 부분 배열 합을 구한다
func maxSubarraySum(arr []int, lo, hi int) int {
	// 기저 조건: 원소가 하나이면 그 값을 반환
	if lo == hi {
		return arr[lo]
	}

	mid := (lo + hi) / 2

	// 정복: 왼쪽과 오른쪽 부분의 최대 부분 배열 합
	leftMax := maxSubarraySum(arr, lo, mid)
	rightMax := maxSubarraySum(arr, mid+1, hi)

	// 결합: 중간을 걸치는 최대 부분 배열 합 계산
	crossMax := maxCrossingSum(arr, lo, mid, hi)

	// 세 값 중 최댓값 반환
	return max3(leftMax, rightMax, crossMax)
}

// maxCrossingSum 함수는 중간 지점을 걸치는 최대 부분 배열 합을 구한다
func maxCrossingSum(arr []int, lo, mid, hi int) int {
	// 중간에서 왼쪽으로 확장하며 최대 합 계산
	leftSum := arr[mid]
	sum := arr[mid]
	for i := mid - 1; i >= lo; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// 중간에서 오른쪽으로 확장하며 최대 합 계산
	rightSum := arr[mid+1]
	sum = arr[mid+1]
	for i := mid + 2; i <= hi; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

// max3 함수는 세 정수 중 최댓값을 반환한다
func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= c {
		return b
	}
	return c
}

func main() {
	// 병합 정렬 예제
	fmt.Println("=== 병합 정렬 ===")
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("정렬 전:", arr)
	sorted := mergeSort(arr)
	fmt.Println("정렬 후:", sorted)

	// 최대 부분 배열 합 예제
	fmt.Println("\n=== 최대 부분 배열 합 ===")
	arr2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("배열:", arr2)
	result := maxSubarraySum(arr2, 0, len(arr2)-1)
	fmt.Println("최대 부분 배열 합:", result)
}
