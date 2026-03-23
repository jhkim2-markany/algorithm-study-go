package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const maxVal = 200001
const mod = 1000000007

// 팩토리얼과 역팩토리얼 배열
var fact [maxVal]int64
var invFact [maxVal]int64

// 모듈러 거듭제곱: base^exp mod m 을 계산한다
func power(base, exp, m int64) int64 {
	result := int64(1)
	base %= m
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % m
		}
		exp /= 2
		base = base * base % m
	}
	return result
}

// 팩토리얼과 역팩토리얼을 전처리한다
func precompute() {
	fact[0] = 1
	for i := 1; i < maxVal; i++ {
		fact[i] = fact[i-1] * int64(i) % mod
	}
	invFact[maxVal-1] = power(fact[maxVal-1], mod-2, mod)
	for i := maxVal - 2; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % mod
	}
}

// (r1,c1)에서 (r2,c2)까지 장애물 없이 이동하는 경로 수를 구한다
// 오른쪽 (c2-c1)번, 아래쪽 (r2-r1)번 이동 → C(dr+dc, dr)
func pathCount(r1, c1, r2, c2 int) int64 {
	dr := r2 - r1
	dc := c2 - c1
	if dr < 0 || dc < 0 {
		return 0
	}
	// C(dr+dc, dr) mod p
	return fact[dr+dc] % mod * invFact[dr] % mod * invFact[dc] % mod
}

// 장애물 좌표를 저장하는 구조체
type obstacle struct {
	r, c int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 팩토리얼 전처리
	precompute()

	// 격자 크기와 장애물 수 입력
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	// 장애물 좌표 입력
	obs := make([]obstacle, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &obs[i].r, &obs[i].c)
	}

	// 장애물을 행, 열 순으로 정렬한다
	sort.Slice(obs, func(i, j int) bool {
		if obs[i].r != obs[j].r {
			return obs[i].r < obs[j].r
		}
		return obs[i].c < obs[j].c
	})

	// 도착점을 장애물 목록 끝에 추가한다 (포함-배제 계산용)
	obs = append(obs, obstacle{n, m})

	// dp[i]: (1,1)에서 obs[i]까지 장애물을 하나도 지나지 않는 경로 수
	dp := make([]int64, len(obs))

	for i := 0; i < len(obs); i++ {
		// (1,1)에서 obs[i]까지의 전체 경로 수
		dp[i] = pathCount(1, 1, obs[i].r, obs[i].c)

		// 중간에 다른 장애물을 지나는 경로를 빼준다 (포함-배제)
		for j := 0; j < i; j++ {
			if obs[j].r <= obs[i].r && obs[j].c <= obs[i].c {
				// obs[j]를 반드시 지나는 경로 수를 빼준다
				sub := dp[j] * pathCount(obs[j].r, obs[j].c, obs[i].r, obs[i].c) % mod
				dp[i] = (dp[i] - sub%mod + mod) % mod
			}
		}
	}

	// 마지막 원소(도착점)의 dp 값이 정답이다
	fmt.Fprintln(writer, dp[len(obs)-1])
}
