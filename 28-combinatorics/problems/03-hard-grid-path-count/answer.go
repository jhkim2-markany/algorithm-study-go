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

// power는 모듈러 거듭제곱 base^exp mod m 을 계산한다
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

// precompute는 팩토리얼과 역팩토리얼을 전처리한다
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

// pathCount는 (r1,c1)에서 (r2,c2)까지의 경로 수를 구한다
func pathCount(r1, c1, r2, c2 int) int64 {
	dr := r2 - r1
	dc := c2 - c1
	if dr < 0 || dc < 0 {
		return 0
	}
	return fact[dr+dc] % mod * invFact[dr] % mod * invFact[dc] % mod
}

// obstacle은 장애물 좌표를 저장하는 구조체이다
type obstacle struct {
	r, c int
}

// gridPathCount는 장애물이 있는 격자에서 (1,1)에서 (n,m)까지의 경로 수를 구한다.
//
// [매개변수]
//   - n: 격자의 행 수
//   - m: 격자의 열 수
//   - obs: 장애물 좌표 목록 (행, 열 순으로 정렬됨)
//
// [반환값]
//   - int64: 장애물을 피하는 경로 수 (mod 1000000007)
//
// [알고리즘 힌트]
//
//	포함-배제 원리를 사용한다.
//	도착점을 장애물 목록 끝에 추가하고,
//	dp[i] = (1,1)에서 obs[i]까지 장애물을 하나도 지나지 않는 경로 수.
//	전체 경로 수에서 중간 장애물을 반드시 지나는 경로 수를 빼준다.
//	마지막 원소(도착점)의 dp 값이 정답이다.
func gridPathCount(n, m int, obs []obstacle) int64 {
	// 도착점을 장애물 목록 끝에 추가
	obs = append(obs, obstacle{n, m})

	dp := make([]int64, len(obs))

	for i := 0; i < len(obs); i++ {
		dp[i] = pathCount(1, 1, obs[i].r, obs[i].c)

		for j := 0; j < i; j++ {
			if obs[j].r <= obs[i].r && obs[j].c <= obs[i].c {
				sub := dp[j] * pathCount(obs[j].r, obs[j].c, obs[i].r, obs[i].c) % mod
				dp[i] = (dp[i] - sub%mod + mod) % mod
			}
		}
	}

	return dp[len(obs)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 팩토리얼 전처리
	precompute()

	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	obs := make([]obstacle, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &obs[i].r, &obs[i].c)
	}

	// 장애물을 행, 열 순으로 정렬
	sort.Slice(obs, func(i, j int) bool {
		if obs[i].r != obs[j].r {
			return obs[i].r < obs[j].r
		}
		return obs[i].c < obs[j].c
	})

	// 핵심 함수 호출
	result := gridPathCount(n, m, obs)

	fmt.Fprintln(writer, result)
}
