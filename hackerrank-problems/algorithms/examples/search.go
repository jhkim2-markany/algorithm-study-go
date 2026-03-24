package main

import "fmt"

// 이진 탐색 (Binary Search) - 기본 패턴 예시
// 정렬된 배열에서 특정 값을 효율적으로 찾는 알고리즘이다.
// 탐색 범위를 매번 절반으로 줄여나가므로 선형 탐색보다 훨씬 빠르다.
// 이진 탐색의 전제 조건: 배열이 정렬되어 있어야 한다.
//
// 이진 탐색의 세 가지 기본 패턴:
//   1. 표준 이진 탐색: 정확히 일치하는 값의 인덱스를 찾는다
//   2. Lower Bound: target 이상인 첫 번째 위치를 찾는다
//   3. Upper Bound: target 초과인 첫 번째 위치를 찾는다
//
// 예시 1: 표준 이진 탐색 (Standard Binary Search)
//   - 시간 복잡도: O(log N)
//   - 공간 복잡도: O(1)
//
// 예시 2: Lower Bound (하한 탐색)
//   - 시간 복잡도: O(log N)
//   - 공간 복잡도: O(1)
//
// 예시 3: Upper Bound (상한 탐색)
//   - 시간 복잡도: O(log N)
//   - 공간 복잡도: O(1)
//
// 예시 4: 매개변수 탐색 (Parametric Search) — 나무 자르기 패턴
//   - 시간 복잡도: O(N log H) (N: 나무 수, H: 최대 나무 높이)
//   - 공간 복잡도: O(1)
//   - 특징: "조건을 만족하는 최댓값/최솟값"을 이진 탐색으로 구한다
//     결정 문제(decision problem)로 변환하여 답의 범위를 좁혀나간다
//
// 예시 5: 투 포인터 (Two Pointer) — 정렬된 배열에서 합 찾기
//   - 시간 복잡도: O(N)
//   - 공간 복잡도: O(1)
//   - 특징: 정렬된 배열의 양 끝에서 포인터를 좁혀가며 목표 합을 찾는다

// binarySearch 함수는 정렬된 배열에서 target 값의 인덱스를 찾는다.
// 값이 존재하면 해당 인덱스를, 존재하지 않으면 -1을 반환한다.
// 탐색 범위를 절반씩 줄여가며 중간값과 target을 비교한다.
func binarySearch(arr []int, target int) int {
	// 탐색 범위의 시작과 끝 인덱스
	lo, hi := 0, len(arr)-1

	for lo <= hi {
		// 중간 인덱스 계산 (오버플로 방지를 위해 lo + (hi-lo)/2 사용)
		mid := lo + (hi-lo)/2

		if arr[mid] == target {
			// 값을 찾은 경우
			return mid
		} else if arr[mid] < target {
			// 중간값이 target보다 작으면 오른쪽 절반을 탐색
			lo = mid + 1
		} else {
			// 중간값이 target보다 크면 왼쪽 절반을 탐색
			hi = mid - 1
		}
	}

	// 값을 찾지 못한 경우
	return -1
}

// lowerBound 함수는 정렬된 배열에서 target 이상인 첫 번째 위치를 반환한다.
// 모든 원소가 target보다 작으면 배열 길이(len(arr))를 반환한다.
// C++의 std::lower_bound와 동일한 동작이다.
// 활용 예: "target이 처음 등장하는 위치", "target을 삽입할 위치"
func lowerBound(arr []int, target int) int {
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := lo + (hi-lo)/2

		if arr[mid] < target {
			// 중간값이 target보다 작으면 답이 될 수 없으므로 오른쪽으로 이동
			lo = mid + 1
		} else {
			// 중간값이 target 이상이면 이 위치가 답의 후보
			hi = mid
		}
	}

	// lo == hi 일 때가 target 이상인 첫 번째 위치
	return lo
}

// upperBound 함수는 정렬된 배열에서 target 초과인 첫 번째 위치를 반환한다.
// 모든 원소가 target 이하이면 배열 길이(len(arr))를 반환한다.
// C++의 std::upper_bound와 동일한 동작이다.
// 활용 예: "target이 마지막으로 등장하는 위치 + 1", "target의 개수 = upperBound - lowerBound"
func upperBound(arr []int, target int) int {
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := lo + (hi-lo)/2

		if arr[mid] <= target {
			// 중간값이 target 이하이면 답이 될 수 없으므로 오른쪽으로 이동
			lo = mid + 1
		} else {
			// 중간값이 target 초과이면 이 위치가 답의 후보
			hi = mid
		}
	}

	// lo == hi 일 때가 target 초과인 첫 번째 위치
	return lo
}

// countOccurrences 함수는 정렬된 배열에서 target의 등장 횟수를 구한다.
// lowerBound와 upperBound의 차이를 이용한다.
func countOccurrences(arr []int, target int) int {
	return upperBound(arr, target) - lowerBound(arr, target)
}

// parametricSearch 함수는 나무 자르기 패턴의 매개변수 탐색을 수행한다.
// 절단기 높이 h로 나무를 잘랐을 때, 잘린 나무 조각의 합이 k 이상이 되는
// 최대 높이 h를 이진 탐색으로 구한다.
// logs: 각 나무의 높이, k: 필요한 나무 길이
// 반환값: 조건을 만족하는 절단기의 최대 높이
func parametricSearch(logs []int, k int) int {
	// 탐색 범위: 0 ~ 가장 높은 나무
	lo, hi := 0, 0
	for _, h := range logs {
		if h > hi {
			hi = h
		}
	}

	// 이진 탐색: "높이 mid로 잘랐을 때 k 이상을 얻을 수 있는가?"
	for lo <= hi {
		mid := lo + (hi-lo)/2

		// 높이 mid로 잘랐을 때 얻는 나무 총량 계산
		total := 0
		for _, h := range logs {
			if h > mid {
				total += h - mid
			}
		}

		if total >= k {
			// 조건을 만족하면 더 높은 높이를 시도 (최댓값을 찾으므로)
			lo = mid + 1
		} else {
			// 조건을 만족하지 못하면 높이를 낮춘다
			hi = mid - 1
		}
	}

	// lo-1이 아닌 hi가 조건을 만족하는 최대 높이
	return hi
}

// twoPointerSum 함수는 정렬된 배열에서 합이 target인 두 원소의 인덱스를 찾는다.
// 양 끝에서 포인터를 좁혀가며 O(N) 시간에 해결한다.
// arr: 오름차순 정렬된 배열, target: 목표 합
// 반환값: (왼쪽 인덱스, 오른쪽 인덱스), 찾지 못하면 (-1, -1)
func twoPointerSum(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			// 합이 목표와 일치
			return left, right
		} else if sum < target {
			// 합이 부족하면 왼쪽 포인터를 오른쪽으로 이동 (더 큰 값 선택)
			left++
		} else {
			// 합이 초과하면 오른쪽 포인터를 왼쪽으로 이동 (더 작은 값 선택)
			right--
		}
	}

	// 합이 target인 쌍을 찾지 못한 경우
	return -1, -1
}

func main() {
	// === 표준 이진 탐색 ===
	fmt.Println("=== 표준 이진 탐색 (Standard Binary Search) ===")
	fmt.Println()

	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Printf("배열: %v\n", arr)
	fmt.Println()

	targets := []int{7, 13, 4, 19, 0}
	for _, target := range targets {
		idx := binarySearch(arr, target)
		if idx != -1 {
			fmt.Printf("  탐색(%2d) → 인덱스 %d에서 발견\n", target, idx)
		} else {
			fmt.Printf("  탐색(%2d) → 존재하지 않음\n", target)
		}
	}

	// === Lower Bound (하한 탐색) ===
	fmt.Println("\n=== Lower Bound (target 이상인 첫 번째 위치) ===")
	fmt.Println()

	arr2 := []int{1, 2, 2, 4, 4, 4, 7, 8, 9}
	fmt.Printf("배열: %v\n", arr2)
	fmt.Println()

	lbTargets := []int{2, 4, 5, 1, 10}
	for _, target := range lbTargets {
		pos := lowerBound(arr2, target)
		fmt.Printf("  lowerBound(%2d) → 위치 %d", target, pos)
		if pos < len(arr2) {
			fmt.Printf(" (arr[%d] = %d)\n", pos, arr2[pos])
		} else {
			fmt.Println(" (배열 끝을 넘어감)")
		}
	}

	// === Upper Bound (상한 탐색) ===
	fmt.Println("\n=== Upper Bound (target 초과인 첫 번째 위치) ===")
	fmt.Println()

	fmt.Printf("배열: %v\n", arr2)
	fmt.Println()

	ubTargets := []int{2, 4, 5, 9, 0}
	for _, target := range ubTargets {
		pos := upperBound(arr2, target)
		fmt.Printf("  upperBound(%2d) → 위치 %d", target, pos)
		if pos < len(arr2) {
			fmt.Printf(" (arr[%d] = %d)\n", pos, arr2[pos])
		} else {
			fmt.Println(" (배열 끝을 넘어감)")
		}
	}

	// === Lower Bound + Upper Bound 활용: 등장 횟수 세기 ===
	fmt.Println("\n=== 등장 횟수 세기 (upperBound - lowerBound) ===")
	fmt.Println()

	fmt.Printf("배열: %v\n", arr2)
	fmt.Println()

	countTargets := []int{2, 4, 7, 5}
	for _, target := range countTargets {
		lb := lowerBound(arr2, target)
		ub := upperBound(arr2, target)
		cnt := countOccurrences(arr2, target)
		fmt.Printf("  값 %d: lowerBound=%d, upperBound=%d → 등장 횟수=%d\n",
			target, lb, ub, cnt)
	}

	// === 매개변수 탐색 (나무 자르기) ===
	fmt.Println("\n=== 매개변수 탐색 (Parametric Search) — 나무 자르기 ===")
	fmt.Println()

	logs := []int{20, 15, 10, 17}
	k := 7
	fmt.Printf("나무 높이: %v\n", logs)
	fmt.Printf("필요한 나무 길이: %d\n", k)
	h := parametricSearch(logs, k)
	fmt.Printf("절단기 최대 높이: %d\n", h)
	fmt.Println("설명: 높이 15로 자르면 (20-15)+(17-15) = 7, 조건 충족")

	fmt.Println()
	logs2 := []int{4, 42, 40, 26, 46}
	k2 := 20
	fmt.Printf("나무 높이: %v\n", logs2)
	fmt.Printf("필요한 나무 길이: %d\n", k2)
	h2 := parametricSearch(logs2, k2)
	fmt.Printf("절단기 최대 높이: %d\n", h2)

	// === 투 포인터 (정렬된 배열에서 합 찾기) ===
	fmt.Println("\n=== 투 포인터 (Two Pointer) — 정렬된 배열에서 합 찾기 ===")
	fmt.Println()

	sortedArr := []int{1, 2, 3, 5, 8, 11, 15}
	fmt.Printf("배열: %v\n", sortedArr)
	fmt.Println()

	tpTargets := []int{10, 16, 100}
	for _, t := range tpTargets {
		l, r := twoPointerSum(sortedArr, t)
		if l != -1 {
			fmt.Printf("  합이 %3d인 쌍: arr[%d]=%d + arr[%d]=%d\n",
				t, l, sortedArr[l], r, sortedArr[r])
		} else {
			fmt.Printf("  합이 %3d인 쌍: 존재하지 않음\n", t)
		}
	}
}
