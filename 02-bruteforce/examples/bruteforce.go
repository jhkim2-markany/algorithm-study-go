package main

import "fmt"

// 브루트포스 기본 예시 - 부분집합 탐색
// 주어진 배열에서 합이 목표값과 같은 부분집합을 찾는다.
// 시간 복잡도: O(2^N) (N: 배열 크기)
// 공간 복잡도: O(N) (재귀 호출 스택)

// findSubsetSum 함수는 재귀적으로 모든 부분집합을 탐색하여
// 합이 target과 같은 부분집합을 찾는다.
func findSubsetSum(arr []int, idx int, currentSum int, target int, selected []bool) bool {
	// 기저 조건: 모든 원소를 검사한 경우
	if idx == len(arr) {
		if currentSum == target {
			return true
		}
		return false
	}

	// 현재 원소를 선택하는 경우
	selected[idx] = true
	if findSubsetSum(arr, idx+1, currentSum+arr[idx], target, selected) {
		return true
	}

	// 현재 원소를 선택하지 않는 경우
	selected[idx] = false
	if findSubsetSum(arr, idx+1, currentSum, target, selected) {
		return true
	}

	return false
}

// printSubset 함수는 선택된 원소들을 출력한다.
func printSubset(arr []int, selected []bool) {
	fmt.Print("{ ")
	first := true
	for i, s := range selected {
		if s {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print(arr[i])
			first = false
		}
	}
	fmt.Println(" }")
}

func main() {
	// 예시 1: 부분집합의 합 찾기
	arr := []int{3, 1, 4, 1, 5}
	target := 8

	fmt.Printf("배열: %v\n", arr)
	fmt.Printf("목표 합: %d\n", target)

	selected := make([]bool, len(arr))
	if findSubsetSum(arr, 0, 0, target, selected) {
		fmt.Print("합이 목표값인 부분집합: ")
		printSubset(arr, selected)
	} else {
		fmt.Println("조건을 만족하는 부분집합이 없습니다")
	}

	// 예시 2: 비트마스크를 이용한 모든 부분집합 출력
	fmt.Println("\n--- 비트마스크로 모든 부분집합 나열 ---")
	small := []int{1, 2, 3}
	n := len(small)

	// 0부터 2^N - 1까지 순회하며 각 비트가 원소 선택 여부를 나타냄
	for mask := 0; mask < (1 << n); mask++ {
		fmt.Print("{ ")
		first := true
		for i := 0; i < n; i++ {
			// i번째 비트가 1이면 해당 원소를 선택
			if mask&(1<<i) != 0 {
				if !first {
					fmt.Print(", ")
				}
				fmt.Print(small[i])
				first = false
			}
		}
		fmt.Println(" }")
	}
}
