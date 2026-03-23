package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 구간 [L, R]
	var l, r int
	fmt.Fscan(reader, &l, &r)

	// SPF(최소 소인수) 체를 구축한다
	spf := make([]int, r+1)
	for i := 2; i <= r; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= r; i++ {
		if spf[i] == i { // i가 소수인 경우
			for j := i * i; j <= r; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	// 각 수의 소인수 합을 SPF 배열로 계산하고 구간 합을 구한다
	totalSum := int64(0)
	for i := l; i <= r; i++ {
		x := i
		pfSum := 0
		// SPF를 이용하여 소인수분해하며 소인수 합을 계산한다
		for x > 1 {
			p := spf[x]
			for x%p == 0 {
				pfSum += p // 같은 소인수가 여러 번이면 그 횟수만큼 더한다
				x /= p
			}
		}
		totalSum += int64(pfSum)
	}

	// 결과 출력
	fmt.Fprintln(writer, totalSum)
}
