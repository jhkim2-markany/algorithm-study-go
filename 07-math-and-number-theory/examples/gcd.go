package main

import "fmt"

// 유클리드 호제법을 이용한 GCD/LCM 계산 예시
// GCD 시간 복잡도: O(log(min(a, b)))
// 공간 복잡도: O(1) (반복문 방식)

// gcd 함수는 두 수의 최대공약수를 반환한다
func gcd(a, b int) int {
	// b가 0이 될 때까지 나머지 연산을 반복한다
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm 함수는 두 수의 최소공배수를 반환한다
func lcm(a, b int) int {
	// LCM = a * b / GCD(a, b)
	// 오버플로 방지를 위해 나누기를 먼저 수행한다
	return a / gcd(a, b) * b
}

func main() {
	// GCD 계산 예시
	fmt.Println("=== GCD (최대공약수) ===")
	pairs := [][2]int{{12, 8}, {48, 18}, {100, 75}, {17, 13}}
	for _, p := range pairs {
		fmt.Printf("GCD(%d, %d) = %d\n", p[0], p[1], gcd(p[0], p[1]))
	}

	// LCM 계산 예시
	fmt.Println("\n=== LCM (최소공배수) ===")
	for _, p := range pairs {
		fmt.Printf("LCM(%d, %d) = %d\n", p[0], p[1], lcm(p[0], p[1]))
	}

	// 여러 수의 GCD/LCM 계산
	fmt.Println("\n=== 여러 수의 GCD/LCM ===")
	nums := []int{12, 18, 24}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = gcd(result, nums[i])
	}
	fmt.Printf("GCD(%v) = %d\n", nums, result)

	result = nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}
	fmt.Printf("LCM(%v) = %d\n", nums, result)
}
