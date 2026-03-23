package main

import "fmt"

// 비트마스킹 (Bitmask) - 기본 비트 연산과 집합 조작
// 정수의 이진 표현을 이용하여 집합을 효율적으로 관리한다.
// 시간 복잡도: O(1) (개별 비트 연산), O(2^N) (전체 부분집합 열거)
// 공간 복잡도: O(1)

// setBit: i번째 비트를 1로 설정한다 (원소 추가)
func setBit(mask, i int) int {
	return mask | (1 << i)
}

// checkBit: i번째 비트가 1인지 확인한다 (원소 포함 여부)
func checkBit(mask, i int) bool {
	return mask&(1<<i) != 0
}

// clearBit: i번째 비트를 0으로 해제한다 (원소 제거)
func clearBit(mask, i int) int {
	return mask &^ (1 << i)
}

// toggleBit: i번째 비트를 반전한다 (포함 여부 토글)
func toggleBit(mask, i int) int {
	return mask ^ (1 << i)
}

// countBits: 마스크에서 1인 비트의 개수를 센다 (집합의 크기)
func countBits(mask int) int {
	count := 0
	for mask > 0 {
		count += mask & 1
		mask >>= 1
	}
	return count
}

// printSet: 비트마스크를 집합 형태로 출력한다
func printSet(mask, n int) {
	fmt.Print("{")
	first := true
	for i := 0; i < n; i++ {
		if checkBit(mask, i) {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print(i)
			first = false
		}
	}
	fmt.Println("}")
}

func main() {
	n := 5 // 원소 수: {0, 1, 2, 3, 4}

	// 빈 집합에서 시작
	mask := 0
	fmt.Printf("초기 상태: ")
	printSet(mask, n)

	// 원소 추가: {1, 3, 4}
	mask = setBit(mask, 1)
	mask = setBit(mask, 3)
	mask = setBit(mask, 4)
	fmt.Printf("원소 1, 3, 4 추가 후: ")
	printSet(mask, n)

	// 원소 포함 여부 확인
	fmt.Printf("원소 3 포함 여부: %v\n", checkBit(mask, 3))
	fmt.Printf("원소 2 포함 여부: %v\n", checkBit(mask, 2))

	// 원소 제거: 3을 제거
	mask = clearBit(mask, 3)
	fmt.Printf("원소 3 제거 후: ")
	printSet(mask, n)

	// 원소 토글: 0을 토글 (없으면 추가)
	mask = toggleBit(mask, 0)
	fmt.Printf("원소 0 토글 후: ")
	printSet(mask, n)

	// 집합 크기
	fmt.Printf("집합 크기: %d\n", countBits(mask))

	// 합집합, 교집합 예시
	a := setBit(setBit(setBit(0, 0), 1), 2) // A = {0, 1, 2}
	b := setBit(setBit(setBit(0, 1), 2), 3) // B = {1, 2, 3}
	fmt.Printf("\nA = ")
	printSet(a, n)
	fmt.Printf("B = ")
	printSet(b, n)
	fmt.Printf("A ∪ B = ")
	printSet(a|b, n)
	fmt.Printf("A ∩ B = ")
	printSet(a&b, n)
	fmt.Printf("A △ B = ")
	printSet(a^b, n)

	// 전체 부분집합 열거
	fmt.Printf("\n집합 {0, 1, 2}의 모든 부분집합:\n")
	full := (1 << 3) - 1 // {0, 1, 2} = 0b111 = 7
	// 공집합 포함 모든 부분집합을 열거한다
	for sub := full; sub > 0; sub = (sub - 1) & full {
		fmt.Printf("  ")
		printSet(sub, 3)
	}
	fmt.Printf("  {}\n") // 공집합
}
