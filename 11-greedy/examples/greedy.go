package main

import (
	"fmt"
	"sort"
)

// 그리디 알고리즘 기본 예시
// 1) 거스름돈 문제: 최소 동전 수로 거스름돈 만들기
// 2) 활동 선택 문제: 최대한 많은 활동 선택하기

// coinChange 함수는 주어진 금액을 최소 동전 수로 거슬러 준다.
// 동전 단위가 배수 관계일 때 그리디가 최적해를 보장한다.
// 시간 복잡도: O(K) (K: 동전 종류 수)
// 공간 복잡도: O(1)
func coinChange(amount int, coins []int) (int, []int) {
	// 큰 동전부터 사용하기 위해 내림차순 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	totalCount := 0
	used := []int{} // 사용된 동전 기록

	for _, coin := range coins {
		// 현재 동전으로 거슬러 줄 수 있는 최대 개수 계산
		count := amount / coin
		if count > 0 {
			totalCount += count
			amount -= coin * count
			for i := 0; i < count; i++ {
				used = append(used, coin)
			}
		}
	}

	return totalCount, used
}

// Activity 구조체는 활동의 시작 시간과 종료 시간을 나타낸다.
type Activity struct {
	start int
	end   int
	name  string
}

// activitySelection 함수는 겹치지 않는 최대 활동 수를 선택한다.
// 종료 시간이 빠른 활동부터 선택하는 그리디 전략을 사용한다.
// 시간 복잡도: O(N log N) (정렬 포함)
// 공간 복잡도: O(N)
func activitySelection(activities []Activity) []Activity {
	// 종료 시간 기준으로 오름차순 정렬
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].end < activities[j].end
	})

	selected := []Activity{activities[0]} // 첫 번째 활동 선택
	lastEnd := activities[0].end

	for i := 1; i < len(activities); i++ {
		// 현재 활동의 시작 시간이 마지막 선택 활동의 종료 시간 이후인 경우 선택
		if activities[i].start >= lastEnd {
			selected = append(selected, activities[i])
			lastEnd = activities[i].end
		}
	}

	return selected
}

func main() {
	// === 거스름돈 문제 예시 ===
	fmt.Println("=== 거스름돈 문제 ===")
	coins := []int{500, 100, 50, 10}
	amount := 1260

	count, used := coinChange(amount, coins)
	fmt.Printf("거슬러 줄 금액: %d원\n", amount)
	fmt.Printf("동전 종류: %v\n", coins)
	fmt.Printf("최소 동전 수: %d개\n", count)
	fmt.Printf("사용된 동전: %v\n", used)

	// === 활동 선택 문제 예시 ===
	fmt.Println("\n=== 활동 선택 문제 ===")
	activities := []Activity{
		{1, 4, "A"},
		{3, 5, "B"},
		{0, 6, "C"},
		{5, 7, "D"},
		{3, 9, "E"},
		{5, 9, "F"},
		{6, 10, "G"},
		{8, 11, "H"},
		{8, 12, "I"},
		{2, 14, "J"},
		{12, 16, "K"},
	}

	fmt.Println("전체 활동 목록:")
	for _, a := range activities {
		fmt.Printf("  활동 %s: [%d, %d)\n", a.name, a.start, a.end)
	}

	selected := activitySelection(activities)
	fmt.Printf("\n선택된 활동 수: %d개\n", len(selected))
	fmt.Println("선택된 활동:")
	for _, a := range selected {
		fmt.Printf("  활동 %s: [%d, %d)\n", a.name, a.start, a.end)
	}
}
