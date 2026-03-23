package main

import "fmt"

// KMP 문자열 검색 알고리즘 구현
// 시간 복잡도: O(N + M) (N: 텍스트 길이, M: 패턴 길이)
// 공간 복잡도: O(M) (실패 함수 배열)

// computeFailure 함수는 패턴의 실패 함수(부분 일치 테이블)를 계산한다
func computeFailure(pattern string) []int {
	m := len(pattern)
	fail := make([]int, m)
	fail[0] = 0 // 첫 번째 위치의 실패 함수 값은 항상 0

	// j: 현재까지 일치한 접두사의 길이
	j := 0
	for i := 1; i < m; i++ {
		// 불일치 시 실패 함수를 따라가며 일치 위치를 찾는다
		for j > 0 && pattern[i] != pattern[j] {
			j = fail[j-1]
		}
		// 현재 문자가 일치하면 일치 길이를 증가시킨다
		if pattern[i] == pattern[j] {
			j++
		}
		fail[i] = j
	}
	return fail
}

// kmpSearch 함수는 텍스트에서 패턴이 등장하는 모든 시작 위치를 반환한다
func kmpSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return nil
	}

	// 실패 함수를 미리 계산한다
	fail := computeFailure(pattern)
	result := []int{}

	// j: 패턴에서 현재 비교 중인 위치
	j := 0
	for i := 0; i < n; i++ {
		// 불일치 시 실패 함수를 이용해 패턴 위치를 조정한다
		for j > 0 && text[i] != pattern[j] {
			j = fail[j-1]
		}
		// 현재 문자가 일치하면 패턴 포인터를 전진시킨다
		if text[i] == pattern[j] {
			j++
		}
		// 패턴 전체가 일치하면 시작 위치를 기록한다
		if j == m {
			result = append(result, i-m+1)
			// 다음 매칭을 위해 실패 함수를 이용해 위치를 조정한다
			j = fail[j-1]
		}
	}
	return result
}

// bruteForceSearch 함수는 브루트포스로 패턴을 검색한다 (비교용)
func bruteForceSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	result := []int{}

	// 텍스트의 각 위치에서 패턴을 하나씩 비교한다
	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// KMP 검색 예제
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	fmt.Println("=== KMP 문자열 검색 ===")
	fmt.Printf("텍스트:  %s\n", text)
	fmt.Printf("패턴:    %s\n", pattern)

	// 실패 함수 출력
	fail := computeFailure(pattern)
	fmt.Printf("실패 함수: %v\n", fail)

	// KMP 검색 결과
	positions := kmpSearch(text, pattern)
	fmt.Printf("KMP 매칭 위치: %v\n", positions)

	// 브루트포스 검색 결과 (비교)
	fmt.Println("\n=== 브루트포스 검색 (비교) ===")
	bfPositions := bruteForceSearch(text, pattern)
	fmt.Printf("브루트포스 매칭 위치: %v\n", bfPositions)

	// 다중 매칭 예제
	fmt.Println("\n=== 다중 매칭 예제 ===")
	text2 := "AAAAAA"
	pattern2 := "AA"
	positions2 := kmpSearch(text2, pattern2)
	fmt.Printf("텍스트: %s, 패턴: %s\n", text2, pattern2)
	fmt.Printf("매칭 위치: %v (총 %d개)\n", positions2, len(positions2))
}
