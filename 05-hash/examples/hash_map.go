package main

import "fmt"

// 해시맵 활용 예시 - Go의 내장 map을 이용한 다양한 해시 연산
// 시간 복잡도: O(N) (N: 데이터 개수)
// 공간 복잡도: O(N)

// countFrequency 함수는 슬라이스에서 각 원소의 등장 횟수를 센다.
func countFrequency(nums []int) map[int]int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	return freq
}

// twoSum 함수는 배열에서 합이 target이 되는 두 원소의 인덱스를 반환한다.
// 해시맵을 이용하여 O(N)에 해결한다.
func twoSum(nums []int, target int) (int, int) {
	// 값 → 인덱스 매핑을 저장하는 해시맵
	seen := make(map[int]int)

	for i, num := range nums {
		// target에서 현재 값을 뺀 보수(complement)를 계산
		complement := target - num
		// 보수가 이미 해시맵에 존재하면 답을 찾은 것
		if j, ok := seen[complement]; ok {
			return j, i
		}
		// 현재 값과 인덱스를 해시맵에 저장
		seen[num] = i
	}

	return -1, -1 // 답이 없는 경우
}

// findDuplicates 함수는 슬라이스에서 중복된 원소를 찾아 반환한다.
// 해시셋을 이용하여 O(N)에 해결한다.
func findDuplicates(nums []int) []int {
	// 이미 본 원소를 저장하는 해시셋
	seen := make(map[int]bool)
	duplicates := []int{}

	for _, num := range nums {
		if seen[num] {
			duplicates = append(duplicates, num)
		} else {
			seen[num] = true
		}
	}

	return duplicates
}

func main() {
	// 예시 1: 빈도수 세기
	fmt.Println("=== 빈도수 세기 ===")
	nums := []int{1, 3, 2, 3, 1, 3, 2, 1, 1}
	freq := countFrequency(nums)
	fmt.Printf("배열: %v\n", nums)
	for key, count := range freq {
		fmt.Printf("  %d → %d회\n", key, count)
	}

	// 예시 2: Two Sum
	fmt.Println("\n=== Two Sum ===")
	arr := []int{2, 7, 11, 15}
	target := 9
	i, j := twoSum(arr, target)
	fmt.Printf("배열: %v, 목표: %d\n", arr, target)
	fmt.Printf("결과: 인덱스 %d와 %d (값: %d + %d = %d)\n", i, j, arr[i], arr[j], target)

	// 예시 3: 중복 원소 찾기
	fmt.Println("\n=== 중복 원소 찾기 ===")
	nums2 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	dups := findDuplicates(nums2)
	fmt.Printf("배열: %v\n", nums2)
	fmt.Printf("중복 원소: %v\n", dups)

	// 예시 4: 키 존재 여부 확인
	fmt.Println("\n=== 키 존재 여부 확인 ===")
	m := map[string]int{"사과": 3, "바나나": 5, "포도": 2}
	fmt.Printf("맵: %v\n", m)
	if val, ok := m["사과"]; ok {
		fmt.Printf("  '사과' 존재: 값 = %d\n", val)
	}
	if _, ok := m["딸기"]; !ok {
		fmt.Println("  '딸기' 존재하지 않음")
	}
}
