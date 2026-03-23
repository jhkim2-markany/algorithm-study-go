package main

import "fmt"

// 소인수분해 (Prime Factorization)
// 시행 나눗셈과 SPF 체 기반 소인수분해를 구현한다.

// 시행 나눗셈 기반 소인수분해
// 주어진 수 n을 소인수분해하여 (소인수, 지수) 쌍의 목록을 반환한다.
// 시간 복잡도: O(√N)
// 공간 복잡도: O(log N) (소인수 개수)
func trialDivision(n int) [][2]int {
	factors := [][2]int{}

	// 2로 나누어 떨어지는 만큼 나눈다
	if n%2 == 0 {
		cnt := 0
		for n%2 == 0 {
			n /= 2
			cnt++
		}
		factors = append(factors, [2]int{2, cnt})
	}

	// 3부터 홀수만 탐색하며 소인수를 찾는다
	for d := 3; d*d <= n; d += 2 {
		if n%d == 0 {
			cnt := 0
			for n%d == 0 {
				n /= d
				cnt++
			}
			factors = append(factors, [2]int{d, cnt})
		}
	}

	// 남은 수가 1보다 크면 그 자체가 소인수이다
	if n > 1 {
		factors = append(factors, [2]int{n, 1})
	}

	return factors
}

// 최소 소인수(SPF) 체 구축
// spf[i]는 i의 최소 소인수를 저장한다.
// 시간 복잡도: O(N log log N)
// 공간 복잡도: O(N)
func buildSPF(n int) []int {
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		spf[i] = i // 초기값은 자기 자신
	}

	// 에라토스테네스의 체와 유사하게 최소 소인수를 기록한다
	for i := 2; i*i <= n; i++ {
		if spf[i] == i { // i가 소수인 경우
			for j := i * i; j <= n; j += i {
				if spf[j] == j { // 아직 갱신되지 않은 경우만
					spf[j] = i
				}
			}
		}
	}
	return spf
}

// SPF 배열을 이용한 소인수분해
// 전처리된 SPF 배열을 사용하여 O(log N)에 소인수분해한다.
func factorizeWithSPF(x int, spf []int) [][2]int {
	factors := [][2]int{}
	for x > 1 {
		p := spf[x]
		cnt := 0
		// 같은 소인수를 모두 나눈다
		for x%p == 0 {
			x /= p
			cnt++
		}
		factors = append(factors, [2]int{p, cnt})
	}
	return factors
}

// 소인수분해 결과를 문자열로 출력하는 헬퍼 함수
func formatFactors(n int, factors [][2]int) string {
	result := fmt.Sprintf("%d = ", n)
	for i, f := range factors {
		if i > 0 {
			result += " × "
		}
		if f[1] == 1 {
			result += fmt.Sprintf("%d", f[0])
		} else {
			result += fmt.Sprintf("%d^%d", f[0], f[1])
		}
	}
	return result
}

func main() {
	// 1. 시행 나눗셈 기반 소인수분해
	fmt.Println("=== 시행 나눗셈 기반 소인수분해 ===")
	testNums := []int{12, 60, 97, 360, 1000000007}
	for _, n := range testNums {
		factors := trialDivision(n)
		fmt.Println(formatFactors(n, factors))
	}

	// 2. SPF 체 기반 소인수분해
	fmt.Println("\n=== SPF 체 기반 소인수분해 (N ≤ 100) ===")
	spf := buildSPF(100)

	// SPF 배열 일부 출력
	fmt.Print("SPF[2..20]: ")
	for i := 2; i <= 20; i++ {
		fmt.Printf("%d ", spf[i])
	}
	fmt.Println()

	// SPF를 이용한 소인수분해
	spfTestNums := []int{12, 30, 60, 84, 97}
	for _, n := range spfTestNums {
		factors := factorizeWithSPF(n, spf)
		fmt.Println(formatFactors(n, factors))
	}

	// 3. 약수의 개수 계산 (소인수분해 활용)
	fmt.Println("\n=== 약수의 개수 (소인수분해 활용) ===")
	for _, n := range []int{12, 60, 360} {
		factors := trialDivision(n)
		// 약수의 개수 = (e1+1) × (e2+1) × ... × (ek+1)
		divisorCount := 1
		for _, f := range factors {
			divisorCount *= (f[1] + 1)
		}
		fmt.Printf("%d의 약수의 개수: %d\n", n, divisorCount)
	}
}
