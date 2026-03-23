package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N, K, M 입력
	var n, k, m int
	fmt.Fscan(reader, &n, &k, &m)

	// K개의 소수 입력
	primes := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &primes[i])
	}

	// 이항 계수 C(n, r) 계산용 테이블
	comb := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		comb[i] = make([]int, k+1)
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}

	// S[j] = 크기가 j인 모든 부분집합에 대해 교집합 크기의 합
	// S[j] = Σ |A_{i1} ∩ A_{i2} ∩ ... ∩ A_{ij}| (모든 j-원소 부분집합)
	s := make([]int, k+1)
	s[0] = n // 공집합에 대응하는 값은 전체 원소 수

	// 비트마스크로 모든 비어있지 않은 부분집합을 열거한다
	for mask := 1; mask < (1 << k); mask++ {
		// 선택된 소수들의 곱을 계산한다
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

		// 선택된 집합의 수와 교집합 크기를 기록한다
		cnt := bits.OnesCount(uint(mask))
		s[cnt] += n / product
	}

	// 정확히 m개의 집합에 속하는 수의 개수를 계산한다
	// exactly[m] = Σ(j=m to k) (-1)^(j-m) × C(j, m) × S[j]
	result := 0
	for j := m; j <= k; j++ {
		term := comb[j][m] * s[j]
		if (j-m)%2 == 0 {
			result += term
		} else {
			result -= term
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, result)
}
