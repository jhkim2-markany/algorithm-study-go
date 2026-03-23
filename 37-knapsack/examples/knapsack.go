package main

import (
	"fmt"
	"sort"
)

// 배낭 문제 (Knapsack Problem) - 0/1 배낭, 완전 배낭, 분할 배낭
// 제한된 용량 내에서 물건들의 가치 합을 최대화한다.
// 0/1 배낭 시간 복잡도: O(N × W)
// 완전 배낭 시간 복잡도: O(N × W)
// 분할 배낭 시간 복잡도: O(N log N)

// knapsack01: 0/1 배낭 문제 (1차원 DP 최적화)
// 각 물건을 넣거나 넣지 않는 두 가지 선택만 가능하다
func knapsack01(weights, values []int, capacity int) int {
	n := len(weights)
	// dp[w] = 용량 w 이하로 담을 때의 최대 가치
	dp := make([]int, capacity+1)

	for i := 0; i < n; i++ {
		// 용량을 역순으로 순회하여 각 물건을 최대 1번만 사용한다
		for w := capacity; w >= weights[i]; w-- {
			if dp[w-weights[i]]+values[i] > dp[w] {
				dp[w] = dp[w-weights[i]] + values[i]
			}
		}
	}
	return dp[capacity]
}

// knapsackUnbounded: 완전 배낭 문제
// 각 물건을 무한히 사용할 수 있다
func knapsackUnbounded(weights, values []int, capacity int) int {
	n := len(weights)
	// dp[w] = 용량 w 이하로 담을 때의 최대 가치
	dp := make([]int, capacity+1)

	for i := 0; i < n; i++ {
		// 용량을 정순으로 순회하여 같은 물건을 여러 번 선택 가능하게 한다
		for w := weights[i]; w <= capacity; w++ {
			if dp[w-weights[i]]+values[i] > dp[w] {
				dp[w] = dp[w-weights[i]] + values[i]
			}
		}
	}
	return dp[capacity]
}

// Item: 분할 배낭용 물건 구조체
type Item struct {
	weight int
	value  int
}

// knapsackFractional: 분할 배낭 문제
// 물건을 쪼개서 일부만 담을 수 있다 (그리디)
func knapsackFractional(items []Item, capacity int) float64 {
	// 단위 무게당 가치 기준으로 내림차순 정렬한다
	sort.Slice(items, func(i, j int) bool {
		ratioI := float64(items[i].value) / float64(items[i].weight)
		ratioJ := float64(items[j].value) / float64(items[j].weight)
		return ratioI > ratioJ
	})

	totalValue := 0.0
	remaining := float64(capacity)

	for _, item := range items {
		if remaining <= 0 {
			break
		}
		if float64(item.weight) <= remaining {
			// 물건 전체를 넣는다
			totalValue += float64(item.value)
			remaining -= float64(item.weight)
		} else {
			// 남은 용량만큼만 쪼개서 넣는다
			fraction := remaining / float64(item.weight)
			totalValue += float64(item.value) * fraction
			remaining = 0
		}
	}
	return totalValue
}

func main() {
	// === 0/1 배낭 예시 ===
	fmt.Println("=== 0/1 배낭 문제 ===")
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 7
	fmt.Printf("물건 무게: %v\n", weights)
	fmt.Printf("물건 가치: %v\n", values)
	fmt.Printf("배낭 용량: %d\n", capacity)
	fmt.Printf("최대 가치: %d\n", knapsack01(weights, values, capacity))

	// === 완전 배낭 예시 ===
	fmt.Println("\n=== 완전 배낭 문제 ===")
	weights2 := []int{2, 3, 4}
	values2 := []int{3, 4, 5}
	capacity2 := 10
	fmt.Printf("물건 무게: %v\n", weights2)
	fmt.Printf("물건 가치: %v\n", values2)
	fmt.Printf("배낭 용량: %d\n", capacity2)
	fmt.Printf("최대 가치: %d\n", knapsackUnbounded(weights2, values2, capacity2))

	// === 분할 배낭 예시 ===
	fmt.Println("\n=== 분할 배낭 문제 ===")
	items := []Item{
		{10, 60},
		{20, 100},
		{30, 120},
	}
	capacity3 := 50
	fmt.Printf("물건: ")
	for _, item := range items {
		fmt.Printf("(무게=%d, 가치=%d) ", item.weight, item.value)
	}
	fmt.Printf("\n배낭 용량: %d\n", capacity3)
	fmt.Printf("최대 가치: %.2f\n", knapsackFractional(items, capacity3))
}
