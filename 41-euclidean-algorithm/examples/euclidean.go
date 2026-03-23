package main

import "fmt"

// 유클리드 호제법 - GCD, LCM, 확장 유클리드 알고리즘, 모듈러 역원
// 시간 복잡도: O(log(min(a, b)))
// 공간 복잡도: O(1) (반복), O(log(min(a, b))) (재귀)

// gcd - 기본 유클리드 호제법 (반복)
func gcd(a, b int) int {
	// 나머지가 0이 될 때까지 반복
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// gcdRecursive - 유클리드 호제법 (재귀)
func gcdRecursive(a, b int) int {
	// 기저 조건: 나머지가 0이면 a가 GCD
	if b == 0 {
		return a
	}
	// GCD(a, b) = GCD(b, a mod b)
	return gcdRecursive(b, a%b)
}

// lcm - 최소공배수 계산
// 오버플로 방지를 위해 a/GCD * b 순서로 계산
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// extGCD - 확장 유클리드 알고리즘
// ax + by = GCD(a, b)를 만족하는 g, x, y를 반환
func extGCD(a, b int) (int, int, int) {
	// 기저 조건: b = 0이면 GCD = a, x = 1, y = 0
	if b == 0 {
		return a, 1, 0
	}
	// 재귀적으로 GCD(b, a%b)에 대해 계수를 구함
	g, x1, y1 := extGCD(b, a%b)
	// 역추적: x = y1, y = x1 - (a/b) * y1
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}

// modInverse - 모듈러 역원 계산
// a * x ≡ 1 (mod m)을 만족하는 x를 반환
// GCD(a, m) = 1일 때만 역원이 존재
func modInverse(a, m int) int {
	g, x, _ := extGCD(a, m)
	// 서로소가 아니면 역원이 존재하지 않음
	if g != 1 {
		return -1
	}
	// 음수일 수 있으므로 양수로 변환
	return (x%m + m) % m
}

func main() {
	// 기본 GCD 예제
	fmt.Println("=== 기본 유클리드 호제법 ===")
	fmt.Printf("GCD(48, 18) = %d\n", gcd(48, 18))
	fmt.Printf("GCD(56, 98) = %d\n", gcd(56, 98))
	fmt.Printf("GCD(101, 103) = %d (서로소)\n", gcd(101, 103))

	// GCD 과정 시각화
	fmt.Println("\n=== GCD(48, 18) 계산 과정 ===")
	a, b := 48, 18
	for b != 0 {
		fmt.Printf("GCD(%d, %d) → %d mod %d = %d\n", a, b, a, b, a%b)
		a, b = b, a%b
	}
	fmt.Printf("GCD = %d\n", a)

	// LCM 예제
	fmt.Println("\n=== 최소공배수 ===")
	fmt.Printf("LCM(12, 18) = %d\n", lcm(12, 18))
	fmt.Printf("LCM(4, 6) = %d\n", lcm(4, 6))

	// 확장 유클리드 알고리즘 예제
	fmt.Println("\n=== 확장 유클리드 알고리즘 ===")
	g, x, y := extGCD(35, 15)
	fmt.Printf("35*(%d) + 15*(%d) = %d (GCD)\n", x, y, g)
	g, x, y = extGCD(161, 28)
	fmt.Printf("161*(%d) + 28*(%d) = %d (GCD)\n", x, y, g)

	// 모듈러 역원 예제
	fmt.Println("\n=== 모듈러 역원 ===")
	inv := modInverse(3, 7)
	fmt.Printf("3의 mod 7 역원 = %d (검증: 3*%d mod 7 = %d)\n", inv, inv, (3*inv)%7)
	inv = modInverse(10, 17)
	fmt.Printf("10의 mod 17 역원 = %d (검증: 10*%d mod 17 = %d)\n", inv, inv, (10*inv)%17)

	// 여러 수의 GCD/LCM
	fmt.Println("\n=== 여러 수의 GCD/LCM ===")
	nums := []int{12, 18, 24, 36}
	g = nums[0]
	l := nums[0]
	for i := 1; i < len(nums); i++ {
		g = gcd(g, nums[i])
		l = lcm(l, nums[i])
	}
	fmt.Printf("GCD(12, 18, 24, 36) = %d\n", g)
	fmt.Printf("LCM(12, 18, 24, 36) = %d\n", l)
}
