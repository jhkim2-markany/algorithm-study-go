package main

import "fmt"

// 조합 (Combination) - 파스칼의 삼각형을 이용한 이항 계수 계산
// n개에서 r개를 순서 없이 선택하는 경우의 수를 구한다.
// 시간 복잡도: O(N²) 전처리, O(1) 쿼리
// 공간 복잡도: O(N²)

const maxN = 1001

// dp 배열에 파스칼의 삼각형을 저장한다
var dp [maxN][maxN]int64

// 파스칼의 삼각형을 구축한다
func buildPascal(n int) {
	for i := 0; i <= n; i++ {
		dp[i][0] = 1 // nC0 = 1
		dp[i][i] = 1 // nCn = 1
		for j := 1; j < i; j++ {
			// 점화식: C(n, r) = C(n-1, r-1) + C(n-1, r)
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
}

// nCr 값을 반환한다
func comb(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	return dp[n][r]
}

func main() {
	// 파스칼의 삼각형 전처리
	buildPascal(10)

	// 예시 1: 5C2 = 10
	fmt.Printf("C(5, 2) = %d\n", comb(5, 2))

	// 예시 2: 10C3 = 120
	fmt.Printf("C(10, 3) = %d\n", comb(10, 3))

	// 예시 3: 6C0 = 1
	fmt.Printf("C(6, 0) = %d\n", comb(6, 0))

	// 예시 4: 파스칼의 삼각형 출력 (0~5행)
	fmt.Println("\n파스칼의 삼각형:")
	for i := 0; i <= 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Println()
	}
}
