package main

import (
	"fmt"
	"sort"
)

// LIS (최장 증가 부분 수열) - 이분 탐색 기반 O(N log N) 풀이
// tails 배열을 유지하며, tails[k]는 길이 k+1인 증가 부분 수열의 마지막 원소 최솟값이다.
// 시간 복잡도: O(N log N)
// 공간 복잡도: O(N)

// lisLength는 수열 a의 LIS 길이를 O(N log N)에 구한다
func lisLength(a []int) int {
	tails := []int{} // tails[k] = 길이 k+1인 증가 부분 수열의 마지막 원소 최솟값

	for _, x := range a {
		// tails에서 x 이상인 첫 번째 위치를 이분 탐색으로 찾는다
		pos := sort.SearchInts(tails, x)
		if pos == len(tails) {
			// x가 tails의 모든 원소보다 크면 뒤에 추가
			tails = append(tails, x)
		} else {
			// 해당 위치의 값을 x로 교체하여 더 작은 값을 유지
			tails[pos] = x
		}
	}

	return len(tails)
}

func main() {
	// 예시 수열
	a := []int{10, 20, 10, 30, 20, 50}

	fmt.Println("수열:", a)
	fmt.Println("LIS 길이:", lisLength(a))
	// 출력: LIS 길이: 4 (LIS 예시: [10, 20, 30, 50])

	// 추가 예시
	b := []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}
	fmt.Println("\n수열:", b)
	fmt.Println("LIS 길이:", lisLength(b))
	// 출력: LIS 길이: 6

	// 엣지 케이스: 역순 정렬
	c := []int{5, 4, 3, 2, 1}
	fmt.Println("\n수열:", c)
	fmt.Println("LIS 길이:", lisLength(c))
	// 출력: LIS 길이: 1

	// 엣지 케이스: 이미 정렬
	d := []int{1, 2, 3, 4, 5}
	fmt.Println("\n수열:", d)
	fmt.Println("LIS 길이:", lisLength(d))
	// 출력: LIS 길이: 5
}
