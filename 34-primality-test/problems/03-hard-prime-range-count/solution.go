package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 구간 체 (Segmented Sieve)를 이용한 소수 구간 카운팅
// [L, R] 범위의 소수를 효율적으로 구한다.

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, r int64
	fmt.Fscan(reader, &l, &r)

	// 1단계: √R 이하의 소수를 에라토스테네스의 체로 구한다
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

	// 작은 소수 목록을 수집한다
	smallPrimes := make([]int64, 0)
	for i := int64(2); i <= limit; i++ {
		if sieve[i] {
			smallPrimes = append(smallPrimes, i)
		}
	}

	// 2단계: 구간 체로 [L, R] 범위의 소수를 판정한다
	size := r - l + 1
	// isPrime[i]는 (l + i)가 소수인지를 나타낸다
	isPrime := make([]bool, size)
	for i := int64(0); i < size; i++ {
		isPrime[i] = true
	}

	// L이 1이면 1은 소수가 아니다
	if l == 1 {
		isPrime[0] = false
	}

	// 각 작은 소수 p에 대해 [L, R] 범위의 배수를 제거한다
	for _, p := range smallPrimes {
		// p의 배수 중 L 이상인 가장 작은 수를 구한다
		start := ((l + p - 1) / p) * p
		// p 자체는 소수이므로 제거하지 않는다
		if start == p {
			start += p
		}
		// p의 배수를 합성수로 표시한다
		for j := start; j <= r; j += p {
			isPrime[j-l] = false
		}
	}

	// 3단계: 소수의 개수를 센다
	count := 0
	for i := int64(0); i < size; i++ {
		if isPrime[i] {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}
