package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// countPrimesInRange는 [l, r] 범위의 소수 개수를 반환한다.
//
// [매개변수]
//   - l: 범위의 시작 (1 이상)
//   - r: 범위의 끝 (l 이상)
//
// [반환값]
//   - int: [l, r] 범위에 포함된 소수의 개수
//
// [알고리즘 힌트]
//
//	구간 체 (Segmented Sieve)를 사용한다.
//	1단계: √R 이하의 소수를 에라토스테네스의 체로 구한다.
//	2단계: 작은 소수들로 [L, R] 범위의 합성수를 제거한다.
//	시간복잡도: O(√R * log(log(R)) + (R-L+1) * log(log(R)))
func countPrimesInRange(l, r int64) int {
	limit := int64(math.Sqrt(float64(r))) + 1
	sieve := make([]bool, limit+1)
	for i := int64(2); i <= limit; i++ {
		sieve[i] = true
	}
	for i := int64(2); i*i <= limit; i++ {
		if sieve[i] {
			for j := i * i; j <= limit; j += i {
				sieve[j] = false
			}
		}
	}

	smallPrimes := make([]int64, 0)
	for i := int64(2); i <= limit; i++ {
		if sieve[i] {
			smallPrimes = append(smallPrimes, i)
		}
	}

	size := r - l + 1
	isPrime := make([]bool, size)
	for i := int64(0); i < size; i++ {
		isPrime[i] = true
	}

	if l == 1 {
		isPrime[0] = false
	}

	for _, p := range smallPrimes {
		start := ((l + p - 1) / p) * p
		if start == p {
			start += p
		}
		for j := start; j <= r; j += p {
			isPrime[j-l] = false
		}
	}

	count := 0
	for i := int64(0); i < size; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, r int64
	fmt.Fscan(reader, &l, &r)

	fmt.Fprintln(writer, countPrimesInRange(l, r))
}
