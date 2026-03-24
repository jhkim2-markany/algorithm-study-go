package main

import (
	"fmt"
	"sort"
)

// 그리디 알고리즘 (Greedy Algorithm) - 기본 패턴 예시
// 매 단계에서 현재 상황에서 가장 좋은 선택(지역 최적)을 반복하여 전체 최적해를 구하는 방법이다.
// 그리디 알고리즘이 최적해를 보장하려면 두 가지 조건이 필요하다:
//   1. 탐욕적 선택 속성 (Greedy Choice Property): 지역 최적 선택이 전역 최적해로 이어진다
//   2. 최적 부분 구조 (Optimal Substructure): 부분 문제의 최적해가 전체 최적해를 구성한다
//
// 예시 1: 활동 선택 문제 (Activity Selection)
//   - 시간 복잡도: O(N log N) (정렬) + O(N) (선택) = O(N log N)
//   - 공간 복잡도: O(N) (활동 목록 저장)
//
// 예시 2: 거스름돈 문제 (Coin Change - Greedy)
//   - 시간 복잡도: O(N log N) (동전 정렬) + O(N) (탐색) = O(N log N)
//   - 공간 복잡도: O(N) (사용된 동전 목록)
//
// 예시 3: 분할 가능 배낭 문제 (Fractional Knapsack)
//   - 시간 복잡도: O(N log N) (정렬)
//   - 공간 복잡도: O(N) (물건 목록 저장)

// Activity 구조체는 하나의 활동을 나타낸다.
// Start는 시작 시간, End는 종료 시간이다.
type Activity struct {
	Start int
	End   int
}

// activitySelection 함수는 서로 겹치지 않는 최대 활동 수를 선택한다.
// 종료 시간이 빠른 순서로 정렬한 뒤, 이전 활동의 종료 시간 이후에
// 시작하는 활동만 선택하는 그리디 전략을 사용한다.
func activitySelection(activities []Activity) []Activity {
	// 종료 시간 기준으로 오름차순 정렬
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].End < activities[j].End
	})

	// 첫 번째 활동은 항상 선택
	selected := []Activity{activities[0]}
	lastEnd := activities[0].End

	// 나머지 활동 중 겹치지 않는 것만 선택
	for i := 1; i < len(activities); i++ {
		if activities[i].Start >= lastEnd {
			selected = append(selected, activities[i])
			lastEnd = activities[i].End
		}
	}

	return selected
}

// coinChangeGreedy 함수는 그리디 방식으로 거스름돈에 필요한 최소 동전 수를 구한다.
// 가장 큰 동전부터 최대한 사용하는 전략이다.
// 주의: 이 방식은 동전 단위가 특정 조건(예: 1, 5, 10, 50, 100, 500원)을
// 만족할 때만 최적해를 보장한다. 임의의 동전 단위에서는 DP가 필요하다.
func coinChangeGreedy(coins []int, amount int) (int, []int) {
	// 동전을 내림차순으로 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	totalCoins := 0
	usedCoins := []int{}
	remaining := amount

	// 큰 동전부터 최대한 사용
	for _, coin := range coins {
		if remaining <= 0 {
			break
		}
		// 현재 동전으로 거슬러 줄 수 있는 최대 개수
		count := remaining / coin
		if count > 0 {
			totalCoins += count
			remaining -= coin * count
			// 사용된 동전 기록
			for k := 0; k < count; k++ {
				usedCoins = append(usedCoins, coin)
			}
		}
	}

	return totalCoins, usedCoins
}

// Item 구조체는 배낭 문제의 물건을 나타낸다.
// Weight는 무게, Value는 가치이다.
type Item struct {
	Weight float64
	Value  float64
}

// fractionalKnapsack 함수는 분할 가능 배낭 문제를 그리디로 해결한다.
// 단위 무게당 가치가 높은 물건부터 넣는 전략이다.
// 물건을 쪼갤 수 있으므로 그리디가 최적해를 보장한다.
func fractionalKnapsack(items []Item, capacity float64) float64 {
	// 단위 무게당 가치 기준으로 내림차순 정렬
	sort.Slice(items, func(i, j int) bool {
		return (items[i].Value / items[i].Weight) > (items[j].Value / items[j].Weight)
	})

	totalValue := 0.0
	remainCap := capacity

	for _, item := range items {
		if remainCap <= 0 {
			break
		}
		if item.Weight <= remainCap {
			// 물건 전체를 넣을 수 있는 경우
			totalValue += item.Value
			remainCap -= item.Weight
		} else {
			// 물건의 일부만 넣는 경우 (분할)
			fraction := remainCap / item.Weight
			totalValue += item.Value * fraction
			remainCap = 0
		}
	}

	return totalValue
}

func main() {
	// === 활동 선택 문제 (Activity Selection) ===
	fmt.Println("=== 활동 선택 문제 (Activity Selection) ===")
	fmt.Println()

	activities := []Activity{
		{Start: 1, End: 4},
		{Start: 3, End: 5},
		{Start: 0, End: 6},
		{Start: 5, End: 7},
		{Start: 3, End: 9},
		{Start: 5, End: 9},
		{Start: 6, End: 10},
		{Start: 8, End: 11},
		{Start: 8, End: 12},
		{Start: 2, End: 14},
		{Start: 12, End: 16},
	}

	fmt.Println("전체 활동 목록:")
	for i, a := range activities {
		fmt.Printf("  활동 %2d: [%2d, %2d)\n", i+1, a.Start, a.End)
	}

	selected := activitySelection(activities)
	fmt.Printf("\n선택된 활동 수: %d\n", len(selected))
	fmt.Println("선택된 활동:")
	for _, a := range selected {
		fmt.Printf("  [%2d, %2d)\n", a.Start, a.End)
	}

	// === 거스름돈 문제 (Coin Change - Greedy) ===
	fmt.Println("\n=== 거스름돈 문제 (Coin Change - Greedy) ===")
	fmt.Println()

	coins := []int{500, 100, 50, 10, 5, 1}
	amount := 1263

	fmt.Printf("동전 종류: %v\n", coins)
	fmt.Printf("거슬러 줄 금액: %d원\n", amount)

	totalCoins, usedCoins := coinChangeGreedy(coins, amount)
	fmt.Printf("필요한 최소 동전 수: %d개\n", totalCoins)
	fmt.Printf("사용된 동전: %v\n", usedCoins)

	// === 분할 가능 배낭 문제 (Fractional Knapsack) ===
	fmt.Println("\n=== 분할 가능 배낭 문제 (Fractional Knapsack) ===")
	fmt.Println()

	items := []Item{
		{Weight: 10, Value: 60},
		{Weight: 20, Value: 100},
		{Weight: 30, Value: 120},
	}
	capacity := 50.0

	fmt.Printf("배낭 용량: %.0f\n", capacity)
	fmt.Println("물건 목록:")
	for i, item := range items {
		fmt.Printf("  물건 %d: 무게=%.0f, 가치=%.0f (단위 가치=%.2f)\n",
			i+1, item.Weight, item.Value, item.Value/item.Weight)
	}

	maxValue := fractionalKnapsack(items, capacity)
	fmt.Printf("\n최대 가치: %.2f\n", maxValue)
}
