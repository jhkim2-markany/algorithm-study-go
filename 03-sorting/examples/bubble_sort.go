package main

import "fmt"

// 버블 정렬 - 인접한 두 원소를 비교하여 교환하는 정렬 알고리즘
// 시간 복잡도: O(N²) (최선: O(N), 이미 정렬된 경우)
// 공간 복잡도: O(1)
// 안정 정렬이다.

// bubbleSort 함수는 정수 슬라이스를 오름차순으로 정렬한다.
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 교환이 발생했는지 추적하여 조기 종료에 활용
		swapped := false
		for j := 0; j < n-1-i; j++ {
			// 인접한 두 원소를 비교하여 순서가 잘못되면 교환
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// 한 번의 순회에서 교환이 없으면 이미 정렬 완료
		if !swapped {
			break
		}
	}
}

func main() {
	// 정렬할 배열 준비
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("정렬 전:", arr)

	// 버블 정렬 수행
	bubbleSort(arr)
	fmt.Println("정렬 후:", arr)

	// 이미 정렬된 배열 테스트 (최선의 경우 O(N))
	sorted := []int{1, 2, 3, 4, 5}
	fmt.Println("\n이미 정렬된 배열:", sorted)
	bubbleSort(sorted)
	fmt.Println("정렬 후:", sorted)
}
