package main

import "fmt"

// 슬라이딩 윈도우 - 연속 구간 탐색 기법
// 고정 크기 윈도우와 가변 크기 윈도우 두 가지 유형을 보여준다.
// 시간 복잡도: O(N)
// 공간 복잡도: O(1) ~ O(K)

// maxSumFixedWindow 함수는 크기 K인 연속 부분 배열의 최대 합을 반환한다.
// 고정 크기 슬라이딩 윈도우 예시이다.
func maxSumFixedWindow(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return 0
	}

	// 첫 번째 윈도우의 합 계산
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum := windowSum

	// 윈도우를 한 칸씩 오른쪽으로 밀면서 최대 합 갱신
	for i := k; i < n; i++ {
		// 새로 들어오는 원소를 더하고, 빠지는 원소를 뺀다
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// minLengthSubarraySum 함수는 합이 target 이상인 연속 부분 배열의 최소 길이를 반환한다.
// 가변 크기 슬라이딩 윈도우 예시이다.
// 조건을 만족하는 부분 배열이 없으면 0을 반환한다.
func minLengthSubarraySum(arr []int, target int) int {
	n := len(arr)
	minLen := n + 1 // 불가능한 큰 값으로 초기화
	windowSum := 0
	left := 0

	// right 포인터를 확장하며 윈도우에 원소 추가
	for right := 0; right < n; right++ {
		windowSum += arr[right]

		// 윈도우 합이 target 이상이면 left를 축소하며 최소 길이 갱신
		for windowSum >= target {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}
			windowSum -= arr[left]
			left++
		}
	}

	// 조건을 만족하는 부분 배열이 없는 경우
	if minLen == n+1 {
		return 0
	}
	return minLen
}

func main() {
	// 고정 크기 슬라이딩 윈도우 예시
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	fmt.Printf("배열: %v\n", arr)
	fmt.Printf("윈도우 크기 %d인 연속 부분 배열의 최대 합: %d\n", k, maxSumFixedWindow(arr, k))

	// 가변 크기 슬라이딩 윈도우 예시
	fmt.Println("\n--- 가변 크기 윈도우 ---")
	arr2 := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Printf("배열: %v, 목표 합: %d\n", arr2, target)

	result := minLengthSubarraySum(arr2, target)
	if result > 0 {
		fmt.Printf("합이 %d 이상인 최소 길이 부분 배열: %d\n", target, result)
	} else {
		fmt.Println("조건을 만족하는 부분 배열이 없습니다")
	}

	// 추가 예시: 조건을 만족하지 않는 경우
	fmt.Println("\n--- 조건 불만족 예시 ---")
	arr3 := []int{1, 1, 1, 1}
	target2 := 100
	fmt.Printf("배열: %v, 목표 합: %d\n", arr3, target2)
	result2 := minLengthSubarraySum(arr3, target2)
	if result2 > 0 {
		fmt.Printf("최소 길이: %d\n", result2)
	} else {
		fmt.Println("조건을 만족하는 부분 배열이 없습니다")
	}
}
