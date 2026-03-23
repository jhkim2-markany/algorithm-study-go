package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// segmentedSieve는 세그먼트 체를 이용하여 [l, r] 구간의 소수 개수를 구한다.
//
// [매개변수]
//   - l: 구간의 시작값
//   - r: 구간의 끝값
//
// [반환값]
//   - int: [l, r] 구간의 소수 개수
//
// [알고리즘 힌트]
//   1. √r까지의 소수를 기본 에라토스테네스의 체로 구한다.
//   2. [l, r] 구간 크기의 불리언 배열을 true로 초기화한다.
//   3. 각 소수 p에 대해 구간 내 p의 배수 시작점을 계산하여 합성수를 제거한다.
//   4. l이 1인 경우 1은 소수가 아니므로 별도 처리한다.
//   5. true로 남은 원소의 수를 세어 반환한다.
func segmentedSieve(l, r int64) int {
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

	smallPrimes := []int64{}
	for i := int64(2); i <= limit; i++ {
		if smallIsPrime[i] {
			smallPrimes = append(smallPrimes, i)
		}
	}

	size := r - l + 1
	seg := make([]bool, size)
	for i := range seg {
		seg[i] = true
	}

	if l <= 1 {
		seg[1-l] = false
	}

	for _, p := range smallPrimes {
		start := ((l + p - 1) / p) * p
		if start == p {
			start += p
		}
		for j := start; j <= r; j += p {
			seg[j-l] = false
		}
	}

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

	var l, r int64
	fmt.Fscan(reader, &l, &r)

	fmt.Fprintln(writer, segmentedSieve(l, r))
}
