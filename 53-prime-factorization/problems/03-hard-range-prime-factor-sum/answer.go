package main

import (
	"bufio"
	"fmt"
	"os"
)

// rangePrimeFactorSum은 구간 [l, r]에 속하는 각 수의 소인수 합의 총합을 반환한다.
// 소인수가 여러 번 나누어지면 그 횟수만큼 더한다.
//
// [매개변수]
//   - l: 구간 시작 (l >= 2)
//   - r: 구간 끝
//
// [반환값]
//   - int64: 구간 내 모든 수의 소인수 합의 총합
//
// [알고리즘 힌트]
//
//	SPF(최소 소인수) 체를 r까지 구축한 뒤,
//	각 수를 SPF로 반복 나누며 소인수 합을 계산한다.
//	같은 소인수가 여러 번이면 그 횟수만큼 더한다.
func rangePrimeFactorSum(l, r int) int64 {
	// SPF 체 구축
	spf := make([]int, r+1)
	for i := 2; i <= r; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= r; i++ {
		if spf[i] == i {
			for j := i * i; j <= r; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	totalSum := int64(0)
	for i := l; i <= r; i++ {
		x := i
		pfSum := 0
		for x > 1 {
			p := spf[x]
			for x%p == 0 {
				pfSum += p
				x /= p
			}
		}
		totalSum += int64(pfSum)
	}
	return totalSum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, r int
	fmt.Fscan(reader, &l, &r)

	fmt.Fprintln(writer, rangePrimeFactorSum(l, r))
}
