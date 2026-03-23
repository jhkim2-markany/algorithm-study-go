package main

import "fmt"

// 재귀 함수 기본 예제 - 팩토리얼, 피보나치, 꼬리 재귀, 메모이제이션
// 시간 복잡도: 함수마다 다름 (본문 주석 참고)
// 공간 복잡도: O(재귀 깊이)

// 팩토리얼 - 기본 재귀
// 시간 복잡도: O(N)
func factorial(n int) int {
	// 기저 조건: 0! = 1, 1! = 1
	if n <= 1 {
		return 1
	}
	// 재귀 단계: n! = n * (n-1)!
	return n * factorial(n-1)
}

// 팩토리얼 - 꼬리 재귀 버전
// 재귀 호출이 함수의 마지막 연산이므로 꼬리 재귀이다
func factorialTail(n int, acc int) int {
	// 기저 조건
	if n <= 1 {
		return acc
	}
	// 꼬리 재귀: 누적값을 매개변수로 전달
	return factorialTail(n-1, n*acc)
}

// 피보나치 - 기본 재귀 (비효율적, O(2^N))
func fibNaive(n int) int {
	// 기저 조건: F(0) = 0, F(1) = 1
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 이진 재귀: 두 번의 재귀 호출
	return fibNaive(n-1) + fibNaive(n-2)
}

// 피보나치 - 메모이제이션 적용 (효율적, O(N))
func fibMemo(n int, memo map[int]int) int {
	// 캐시에 결과가 있으면 즉시 반환
	if val, ok := memo[n]; ok {
		return val
	}
	// 기저 조건
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 결과를 계산하고 캐시에 저장
	result := fibMemo(n-1, memo) + fibMemo(n-2, memo)
	memo[n] = result
	return result
}

// 거듭제곱 - 재귀적 분할 정복
// 시간 복잡도: O(log N)
func power(base, exp int) int {
	// 기저 조건: x^0 = 1
	if exp == 0 {
		return 1
	}
	// 짝수 지수: x^n = (x^(n/2))^2
	if exp%2 == 0 {
		half := power(base, exp/2)
		return half * half
	}
	// 홀수 지수: x^n = x * x^(n-1)
	return base * power(base, exp-1)
}

func main() {
	// 팩토리얼 예제
	fmt.Println("=== 팩토리얼 ===")
	fmt.Printf("5! = %d (기본 재귀)\n", factorial(5))
	fmt.Printf("5! = %d (꼬리 재귀)\n", factorialTail(5, 1))

	// 피보나치 예제
	fmt.Println("\n=== 피보나치 ===")
	fmt.Printf("F(10) = %d (기본 재귀)\n", fibNaive(10))

	memo := make(map[int]int)
	fmt.Printf("F(10) = %d (메모이제이션)\n", fibMemo(10, memo))
	fmt.Printf("F(40) = %d (메모이제이션)\n", fibMemo(40, memo))

	// 거듭제곱 예제
	fmt.Println("\n=== 거듭제곱 ===")
	fmt.Printf("2^10 = %d\n", power(2, 10))
	fmt.Printf("3^5 = %d\n", power(3, 5))

	// 호출 스택 시각화 예제
	fmt.Println("\n=== 호출 스택 시각화 (팩토리얼 4) ===")
	factorialTrace(4, 0)
}

// 호출 스택을 시각적으로 보여주는 팩토리얼 함수
func factorialTrace(n int, depth int) int {
	// 들여쓰기로 호출 깊이를 표현
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	fmt.Printf("%sfactorial(%d) 호출\n", indent, n)

	if n <= 1 {
		fmt.Printf("%sfactorial(%d) = 1 반환 (기저 조건)\n", indent, n)
		return 1
	}

	result := n * factorialTrace(n-1, depth+1)
	fmt.Printf("%sfactorial(%d) = %d 반환\n", indent, n, result)
	return result
}
