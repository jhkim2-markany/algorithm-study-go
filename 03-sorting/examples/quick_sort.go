package main

import "fmt"

// 퀵 정렬 - 피벗을 기준으로 분할하는 정렬 알고리즘
// 시간 복잡도: O(N log N) (평균), O(N²) (최악, 이미 정렬된 경우)
// 공간 복잡도: O(log N) (재귀 호출 스택)
// 불안정 정렬이다.

// quickSort 함수는 정수 슬라이스를 오름차순으로 정렬한다.
func quickSort(arr []int, low, high int) {
	if low < high {
		// 피벗을 기준으로 배열을 분할하고 피벗의 최종 위치를 반환
		pivotIdx := partition(arr, low, high)

		// 피벗 왼쪽과 오른쪽을 재귀적으로 정렬
		quickSort(arr, low, pivotIdx-1)
		quickSort(arr, pivotIdx+1, high)
	}
}

// partition 함수는 마지막 원소를 피벗으로 선택하고,
// 피벗보다 작은 원소는 왼쪽, 큰 원소는 오른쪽으로 분할한다.
func partition(arr []int, low, high int) int {
	pivot := arr[high] // 마지막 원소를 피벗으로 선택
	i := low - 1       // 피벗보다 작은 원소들의 마지막 인덱스

	for j := low; j < high; j++ {
		// 현재 원소가 피벗보다 작거나 같으면 왼쪽 영역으로 이동
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 피벗을 올바른 위치에 배치
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	// 정렬할 배열 준비
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("정렬 전:", arr)

	// 퀵 정렬 수행
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("정렬 후:", arr)

	// 중복 원소가 있는 배열 테스트
	duplicates := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	fmt.Println("\n중복 원소 배열:", duplicates)
	quickSort(duplicates, 0, len(duplicates)-1)
	fmt.Println("정렬 후:", duplicates)
}
