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

	// 입력: 최대 범위 N, 쿼리 수 Q
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// SPF(최소 소인수) 체를 구축한다
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= n; i++ {
		if spf[i] == i { // i가 소수인 경우
			for j := i * i; j <= n; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	// 각 수의 서로 다른 소인수 개수를 전처리한다
	distinctCount := make([]int, n+1)
	for i := 2; i <= n; i++ {
		x := i
		cnt := 0
		for x > 1 {
			p := spf[x]
			cnt++
			// 같은 소인수를 모두 나눈다
			for x%p == 0 {
				x /= p
			}
		}
		distinctCount[i] = cnt
	}

	// 쿼리 처리: 전처리된 결과를 O(1)에 응답한다
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(reader, &x)
		fmt.Fprintln(writer, distinctCount[x])
	}
}
