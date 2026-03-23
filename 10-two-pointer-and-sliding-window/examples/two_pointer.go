package main

import "fmt"

// 투 포인터 - 정렬된 배열에서 두 수의 합 찾기
// 정렬된 배열에서 합이 target인 두 원소의 인덱스를 찾는다.
// 시간 복잡도: O(N)
// 공간 복잡도: O(1)

// twoSumSorted 함수는 정렬된 배열에서 합이 target인 두 원소의 인덱스를 반환한다.
// 찾지 못하면 (-1, -1)을 반환한다.
func twoSumSorted(arr []int, target int) (int, int) {
	// 왼쪽 포인터는 배열의 시작, 오른쪽 포인터는 배열의 끝에서 시작
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			// 합이 목표값과 같으면 결과 반환
			return left, right
		} else if sum < target {
			// 합이 목표값보다 작으면 왼쪽 포인터를 오른쪽으로 이동 (더 큰 값 선택)
			left++
		} else {
			// 합이 목표값보다 크면 오른쪽 포인터를 왼쪽으로 이동 (더 작은 값 선택)
			right--
		}
	}

	// 조건을 만족하는 쌍이 없는 경우
	return -1, -1
}

// removeDuplicates 함수는 정렬된 배열에서 중복을 제거하고 고유 원소의 개수를 반환한다.
// 같은 방향 투 포인터 예시이다.
func removeDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// slow 포인터는 고유 원소의 마지막 위치를 추적
	slow := 0

	// fast 포인터가 배열을 순회하며 새로운 값을 발견하면 slow 위치에 복사
	for fast := 1; fast < len(arr); fast++ {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
	}

	// 고유 원소의 개수 반환
	return slow + 1
}

func main() {
	// 반대 방향 투 포인터 예시: 두 수의 합
	arr := []int{1, 3, 5, 7, 9, 11}
	target := 12
	fmt.Printf("배열: %v, 목표값: %d\n", arr, target)

	i, j := twoSumSorted(arr, target)
	if i != -1 {
		fmt.Printf("인덱스: (%d, %d), 값: (%d, %d)\n", i, j, arr[i], arr[j])
	} else {
		fmt.Println("조건을 만족하는 쌍이 없습니다")
	}

	// 같은 방향 투 포인터 예시: 중복 제거
	fmt.Println("\n--- 중복 제거 ---")
	arr2 := []int{1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Printf("원본 배열: %v\n", arr2)

	count := removeDuplicates(arr2)
	fmt.Printf("고유 원소 개수: %d\n", count)
	fmt.Printf("결과 배열: %v\n", arr2[:count])
}
