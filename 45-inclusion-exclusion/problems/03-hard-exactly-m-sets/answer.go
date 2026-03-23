package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

// exactlyMSets는 포함 배제의 원리를 이용하여 정확히 m개의 집합에 속하는 원소의 수를 구한다.
//
// [매개변수]
//   - n: 전체 원소 범위 (1부터 n)
//   - primes: 소수 배열 (각 소수 p에 대해 집합 A_p = {p의 배수})
//   - m: 정확히 속해야 하는 집합의 수
//
// [반환값]
//   - int: 정확히 m개의 집합에 속하는 원소의 수
//
// [알고리즘 힌트]
//   1. 이항 계수 테이블 C(n, r)을 미리 계산한다.
//   2. 비트마스크로 모든 부분집합을 열거하여 S[j] (크기 j인 부분집합의 교집합 크기 합)를 구한다.
//   3. 정확히 m개 공식: exactly[m] = Σ(j=m to k) (-1)^(j-m) × C(j, m) × S[j]를 적용한다.
func exactlyMSets(n int, primes []int, m int) int {
	k := len(primes)

	// 이항 계수 테이블
	comb := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		comb[i] = make([]int, k+1)
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}

	// S[j] 계산
	s := make([]int, k+1)
	s[0] = n

	for mask := 1; mask < (1 << k); mask++ {
		product := 1
		overflow := false
		for i := 0; i < k; i++ {
			if mask&(1<<i) != 0 {
				product *= primes[i]
				if product > n {
					overflow = true
					break
				}
			}
		}

		if overflow {
			continue
		}

		cnt := bits.OnesCount(uint(mask))
		s[cnt] += n / product
	}

	// 정확히 m개 공식 적용
	result := 0
	for j := m; j <= k; j++ {
		term := comb[j][m] * s[j]
		if (j-m)%2 == 0 {
			result += term
		} else {
			result -= term
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k, m int
	fmt.Fscan(reader, &n, &k, &m)

	primes := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &primes[i])
	}

	fmt.Fprintln(writer, exactlyMSets(n, primes, m))
}
