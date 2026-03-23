package main

import (
	"bufio"
	"fmt"
	"os"
)

// nCrMod는 이항 계수 C(n, r) mod MOD를 반환한다.
// 팩토리얼과 역팩토리얼을 전처리하여 각 쿼리를 O(1)에 처리한다.
//
// [매개변수]
//   - n: 전체 원소 수
//   - r: 선택할 원소 수
//   - fact: 팩토리얼 배열 (fact[i] = i! mod MOD)
//   - invFact: 역팩토리얼 배열 (invFact[i] = (i!)^(-1) mod MOD)
//
// [반환값]
//   - int64: C(n, r) mod MOD
//
// [알고리즘 힌트]
//
//	C(n, r) = n! / (r! * (n-r)!) = fact[n] * invFact[r] * invFact[n-r] mod MOD
//	역팩토리얼은 페르마 소정리로 fact[MAXN-1]의 역원을 구한 뒤 역순으로 계산한다.
func nCrMod(n, r int, fact, invFact []int64) int64 {
	const MOD = 1000000007
	if r < 0 || r > n {
		return 0
	}
	return fact[n] % MOD * invFact[r] % MOD * invFact[n-r] % MOD
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	const MOD = 1000000007
	const MAXN = 1000001

	// 팩토리얼 전처리
	fact := make([]int64, MAXN)
	fact[0] = 1
	for i := 1; i < MAXN; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}

	// 빠른 거듭제곱
	modPow := func(a, b, m int64) int64 {
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

	// 역팩토리얼 전처리
	invFact := make([]int64, MAXN)
	invFact[MAXN-1] = modPow(fact[MAXN-1], MOD-2, MOD)
	for i := MAXN - 2; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % MOD
	}

	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var n, r int
		fmt.Fscan(reader, &n, &r)
		fmt.Fprintln(writer, nCrMod(n, r, fact, invFact))
	}
}
