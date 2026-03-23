package main

import "fmt"

// 파라메트릭 서치 - 결정 문제로 변환하여 이진 탐색으로 최적해를 찾는 기법
// 예시: N개의 나무에서 높이 H로 잘라 최소 M미터의 나무를 얻을 수 있는 최대 H를 구한다.
// 시간 복잡도: O(N × log(max_height))
// 공간 복잡도: O(1)

// canObtain 함수는 높이 h로 나무를 잘랐을 때 목표량 이상을 얻을 수 있는지 판별한다.
// 이것이 파라메트릭 서치의 핵심인 "결정 함수"이다.
func canObtain(trees []int, h int, target int) bool {
	total := 0
	for _, tree := range trees {
		if tree > h {
			// 나무 높이가 h보다 크면 차이만큼 가져갈 수 있다
			total += tree - h
		}
	}
	// 목표량 이상을 얻을 수 있으면 true
	return total >= target
}

// parametricSearch 함수는 이진 탐색으로 조건을 만족하는 최대 높이를 찾는다.
func parametricSearch(trees []int, target int) int {
	// 탐색 범위 설정: 최소 0, 최대는 가장 높은 나무
	lo, hi := 0, 0
	for _, tree := range trees {
		if tree > hi {
			hi = tree
		}
	}

	// 이진 탐색으로 최적의 높이를 찾는다
	result := 0
	for lo <= hi {
		mid := (lo + hi) / 2

		if canObtain(trees, mid, target) {
			// 조건을 만족하면 더 높은 값도 가능한지 확인
			result = mid
			lo = mid + 1
		} else {
			// 조건을 불만족하면 높이를 낮춘다
			hi = mid - 1
		}
	}

	return result
}

func main() {
	// 예시: 나무 높이가 [20, 15, 10, 17]이고, 7미터가 필요한 경우
	trees := []int{20, 15, 10, 17}
	target := 7

	fmt.Println("=== 파라메트릭 서치 예시: 나무 자르기 ===")
	fmt.Printf("나무 높이: %v\n", trees)
	fmt.Printf("필요한 나무: %d미터\n", target)

	h := parametricSearch(trees, target)
	fmt.Printf("절단 높이: %d\n", h)

	// 검증: 실제로 얻는 나무 양 계산
	total := 0
	for _, tree := range trees {
		if tree > h {
			total += tree - h
		}
	}
	fmt.Printf("얻는 나무: %d미터\n", total)

	// 추가 예시: 더 많은 나무가 필요한 경우
	fmt.Println("\n=== 추가 예시 ===")
	target2 := 20
	h2 := parametricSearch(trees, target2)
	fmt.Printf("필요한 나무: %d미터 → 절단 높이: %d\n", target2, h2)

	total2 := 0
	for _, tree := range trees {
		if tree > h2 {
			total2 += tree - h2
		}
	}
	fmt.Printf("얻는 나무: %d미터\n", total2)
}
