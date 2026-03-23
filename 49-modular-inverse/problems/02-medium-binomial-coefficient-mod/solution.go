package main

import (
	"bufio"
	"fmt"
	"os"
)

// 이항 계수 C(N, R) mod M 계산
// 팩토리얼과 역팩토리얼을 전처리하여 각 쿼리를 O(1)에 처리한다

const MOD = 1000000007
const MAXN = 1000001

var (
	fact    [MAXN]int64 // fact[i] = i! mod MOD
	invFact [MAXN]int64 // invFact[i] = (i!)^(-1) mod MOD
)

// modPow는 빠른 거듭제곱으로 a^b mod m을 계산한다
func modPow(a, b, m int64) int64 {
	a %= m
	result := int64(1)
	for b > 0 {
		if b%2 == 1 {
			result = result * a % m
		}
		b /= 2
		a = a * a % m
	}
	return result
}

// init은 팩토리얼과 역팩토리얼 배열을 전처리한다
func init() {
	// 팩토리얼 계산: fact[i] = i! mod MOD
	fact[0] = 1
	for i := 1; i < MAXN; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}

	// 역팩토리얼 계산: fact[MAXN-1]의 역원을 구한 뒤 역순으로 계산한다
	invFact[MAXN-1] = modPow(fact[MAXN-1], MOD-2, MOD)
	for i := MAXN - 2; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % MOD
	}
}

// nCr은 이항 계수 C(n, r) mod MOD를 반환한다
func nCr(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	// C(n, r) = n! / (r! * (n-r)!) = fact[n] * invFact[r] * invFact[n-r]
	return fact[n] % MOD * invFact[r] % MOD * invFact[n-r] % MOD
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 쿼리 수
	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)

		// 전처리된 팩토리얼/역팩토리얼로 O(1)에 계산
		fmt.Fprintln(writer, nCr(n, r))
	}
}
