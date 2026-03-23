package main

import "fmt"

// 확률론 (Probability Theory) - 확률 계산과 기댓값
// 확률, 기댓값, 조건부 확률, 확률 DP의 기본 구현을 보여준다.
// 단순 확률 계산: O(N)
// 확률 DP: O(N × M)

// diceProb: 주사위를 n번 던져서 합이 정확히 target이 될 확률을 구한다
// 확률 DP를 사용하여 각 상태의 확률을 전파한다
func diceProb(n, target int) float64 {
	// dp[i][j] = 주사위를 i번 던져서 합이 j가 될 확률
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, target+1)
	}
	// 초기 상태: 0번 던져서 합이 0일 확률은 1
	dp[0][0] = 1.0

	// 각 주사위를 던질 때마다 1~6의 눈이 나올 확률은 1/6
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			for face := 1; face <= 6; face++ {
				if j >= face && dp[i-1][j-face] > 0 {
					dp[i][j] += dp[i-1][j-face] / 6.0
				}
			}
		}
	}

	return dp[n][target]
}

// expectedTosses: 공정한 동전을 던져서 앞면이 나올 때까지의 기대 횟수를 구한다
// 기하 분포의 기댓값: E[X] = 1/p (p = 앞면 확률)
func expectedTosses(p float64) float64 {
	// 기하 분포: 성공 확률이 p일 때 첫 성공까지의 기대 시행 횟수
	return 1.0 / p
}

// expectedValueLinear: 기댓값의 선형성을 활용한 예시
// n개의 카드 중 k개를 뽑을 때, 뽑힌 카드 값의 합의 기댓값을 구한다
func expectedValueLinear(cards []int, k int) float64 {
	n := len(cards)
	if n == 0 || k == 0 {
		return 0
	}

	// 기댓값의 선형성: E[합] = E[X₁] + E[X₂] + ... + E[Xₙ]
	// 각 카드가 뽑힐 확률 = k/n
	totalExpected := 0.0
	probSelected := float64(k) / float64(n)
	for _, card := range cards {
		// 각 카드의 기댓값 기여 = 카드 값 × 뽑힐 확률
		totalExpected += float64(card) * probSelected
	}
	return totalExpected
}

// couponCollector: 쿠폰 수집가 문제의 기댓값을 구한다
// n종류의 쿠폰을 모두 모으기 위해 필요한 기대 구매 횟수
func couponCollector(n int) float64 {
	// 기댓값의 선형성 활용
	// i종류를 모은 상태에서 새로운 종류를 얻을 확률 = (n-i)/n
	// 새로운 종류를 얻기까지의 기대 횟수 = n/(n-i)
	expected := 0.0
	for i := 0; i < n; i++ {
		// i종류를 이미 모았을 때, 새로운 종류를 얻기까지의 기대 횟수
		expected += float64(n) / float64(n-i)
	}
	return expected
}

// randomWalkProb: 1차원 랜덤 워크에서 위치 0에서 시작하여
// steps번 이동 후 위치 target에 있을 확률을 구한다
// 각 단계에서 +1 또는 -1로 이동 (확률 각 1/2)
func randomWalkProb(steps, target int) float64 {
	// target에 도달하려면 오른쪽 이동 r번, 왼쪽 이동 l번이 필요
	// r + l = steps, r - l = target
	// r = (steps + target) / 2
	if (steps+target)%2 != 0 || target > steps || target < -steps {
		return 0 // 도달 불가능
	}

	r := (steps + target) / 2
	// 확률 = C(steps, r) × (1/2)^steps
	// C(steps, r)을 직접 계산한다
	prob := 1.0
	for i := 0; i < r; i++ {
		prob *= float64(steps-i) / float64(i+1)
	}
	// (1/2)^steps를 곱한다
	for i := 0; i < steps; i++ {
		prob /= 2.0
	}
	return prob
}

func main() {
	// === 주사위 확률 DP 예시 ===
	fmt.Println("=== 주사위 확률 DP ===")
	n, target := 2, 7
	prob := diceProb(n, target)
	fmt.Printf("주사위 %d번 던져서 합이 %d가 될 확률: %.6f\n", n, target, prob)

	n2, target2 := 3, 10
	prob2 := diceProb(n2, target2)
	fmt.Printf("주사위 %d번 던져서 합이 %d가 될 확률: %.6f\n", n2, target2, prob2)

	// === 기하 분포 기댓값 예시 ===
	fmt.Println("\n=== 기하 분포 기댓값 ===")
	fmt.Printf("공정한 동전 앞면까지 기대 횟수: %.1f\n", expectedTosses(0.5))
	fmt.Printf("확률 1/6 사건까지 기대 횟수: %.1f\n", expectedTosses(1.0/6.0))

	// === 기댓값의 선형성 예시 ===
	fmt.Println("\n=== 기댓값의 선형성 ===")
	cards := []int{1, 2, 3, 4, 5}
	k := 3
	ev := expectedValueLinear(cards, k)
	fmt.Printf("카드 %v 중 %d장 뽑을 때 합의 기댓값: %.2f\n", cards, k, ev)

	// === 쿠폰 수집가 문제 ===
	fmt.Println("\n=== 쿠폰 수집가 문제 ===")
	for _, cn := range []int{3, 5, 10} {
		fmt.Printf("%d종류 쿠폰 모으기 기대 횟수: %.2f\n", cn, couponCollector(cn))
	}

	// === 랜덤 워크 확률 ===
	fmt.Println("\n=== 1차원 랜덤 워크 ===")
	steps := 6
	for t := -6; t <= 6; t += 2 {
		p := randomWalkProb(steps, t)
		if p > 0 {
			fmt.Printf("%d번 이동 후 위치 %d에 있을 확률: %.4f\n", steps, t, p)
		}
	}
}
