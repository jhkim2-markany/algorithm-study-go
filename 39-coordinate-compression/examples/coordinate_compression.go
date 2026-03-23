package main

import (
	"fmt"
	"sort"
)

// 좌표 압축 - 정렬 + 중복 제거 + 이진 탐색 기반
// 시간 복잡도: O(N log N)
// 공간 복잡도: O(N)

// compress는 주어진 좌표 배열을 압축하여 0부터 시작하는 인덱스로 변환한다.
// 반환값: 압축된 좌표 배열, 고유 좌표 배열(역매핑용)
func compress(coords []int) ([]int, []int) {
	// 1단계: 좌표 복사 후 정렬
	sorted := make([]int, len(coords))
	copy(sorted, coords)
	sort.Ints(sorted)

	// 2단계: 중복 제거
	unique := []int{sorted[0]}
	for i := 1; i < len(sorted); i++ {
		if sorted[i] != sorted[i-1] {
			unique = append(unique, sorted[i])
		}
	}

	// 3단계: 이진 탐색으로 원래 좌표를 압축된 인덱스로 매핑
	result := make([]int, len(coords))
	for i, v := range coords {
		// sort.SearchInts는 정렬된 배열에서 값의 위치를 이진 탐색으로 찾는다
		result[i] = sort.SearchInts(unique, v)
	}

	return result, unique
}

func main() {
	// 예제 1: 기본 좌표 압축
	coords := []int{100, 5000, 30, 5000, 100}
	fmt.Println("원본 좌표:", coords)

	compressed, unique := compress(coords)
	fmt.Println("고유 좌표:", unique)
	fmt.Println("압축 결과:", compressed)

	// 예제 2: 음수 좌표 포함
	coords2 := []int{-1000000000, 500, 0, 500, 1000000000, -1000000000}
	fmt.Println("\n원본 좌표:", coords2)

	compressed2, unique2 := compress(coords2)
	fmt.Println("고유 좌표:", unique2)
	fmt.Println("압축 결과:", compressed2)

	// 역매핑 예시: 압축된 인덱스로부터 원래 좌표 복원
	fmt.Println("\n역매핑 예시:")
	for i, c := range compressed2 {
		fmt.Printf("  압축 좌표 %d → 원래 좌표 %d (원본: %d)\n", c, unique2[c], coords2[i])
	}
}
