package main

import "fmt"

// 병합 정렬 - 분할 정복을 이용한 정렬 알고리즘
// 시간 복잡도: O(N log N) (최선, 평균, 최악 모두 동일)
// 공간 복잡도: O(N) (임시 배열 필요)
// 안정 정렬이다.

// mergeSort 함수는 정수 슬라이스를 오름차순으로 정렬한다.
func mergeSort(arr []int) []int {
	n := len(arr)
	// 원소가 1개 이하이면 이미 정렬된 상태
	if n <= 1 {
		return arr
	}

	// 배열을 절반으로 분할
	mid := n / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// 두 정렬된 배열을 병합
	return merge(left, right)
}

// merge 함수는 두 정렬된 슬라이스를 하나의 정렬된 슬라이스로 병합한다.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// 두 배열의 앞쪽 원소를 비교하여 작은 것부터 결과에 추가
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

func main() {
	// 정렬할 배열 준비
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("정렬 전:", arr)

	// 병합 정렬 수행
	sorted := mergeSort(arr)
	fmt.Println("정렬 후:", sorted)

	// 역순 배열 테스트
	reversed := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("\n역순 배열:", reversed)
	sorted2 := mergeSort(reversed)
	fmt.Println("정렬 후:", sorted2)
}
