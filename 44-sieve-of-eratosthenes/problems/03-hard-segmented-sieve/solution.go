package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 세그먼트 체 (Segmented Sieve)
// [L, R] 구간의 소수를 효율적으로 구한다.
func segmentedSieve(l, r int64) int {
	// √R까지의 소수를 기본 체로 구한다
	limit := int64(math.Sqrt(float64(r))) + 1
	smallIsPrime := make([]bool, limit+1)
	for i := int64(2); i <= limit; i++ {
		smallIsPrime[i] = true
	}
	for i := int64(2); i*i <= limit; i++ {
		if smallIsPrime[i] {
			for j := i * i; j <= limit; j += i {
				smallIsPrime[j] = false
			}
		}
	}

	// 소수 목록을 수집한다
	smallPrimes := []int64{}
	for i := int64(2); i <= limit; i++ {
		if smallIsPrime[i] {
			smallPrimes = append(smallPrimes, i)
		}
	}

	// [L, R] 구간의 소수 여부 배열을 생성한다
	size := r - l + 1
	seg := make([]bool, size)
	for i := range seg {
		seg[i] = true
	}

	// L이 1인 경우 1은 소수가 아니므로 제외한다
	if l <= 1 {
		seg[1-l] = false
	}

	// 각 소수의 배수를 구간에서 제거한다
	for _, p := range smallPrimes {
		// 구간 내 p의 배수 시작점을 구한다
		start := ((l + p - 1) / p) * p
		if start == p {
			// p 자체는 소수이므로 건너뛴다
			start += p
		}
		for j := start; j <= r; j += p {
			seg[j-l] = false
		}
	}

	// 소수 개수를 센다
	count := 0
	for i := int64(0); i < size; i++ {
		if seg[i] && l+i >= 2 {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// L, R 입력
	var l, r int64
	fmt.Fscan(reader, &l, &r)

	// 세그먼트 체로 구간 내 소수 개수를 구한다
	result := segmentedSieve(l, r)
	fmt.Fprintln(writer, result)
}
