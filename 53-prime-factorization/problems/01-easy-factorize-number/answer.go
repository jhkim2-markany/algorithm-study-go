package main

import (
	"bufio"
	"fmt"
	"os"
)

// factorize는 자연수 N을 소인수분해하여 소인수 목록을 오름차순으로 반환한다.
// 같은 소인수가 여러 번 나누어지면 그 횟수만큼 포함한다.
//
// [매개변수]
//   - n: 소인수분해할 자연수 (n >= 2)
//
// [반환값]
//   - []int: 소인수 목록 (오름차순, 중복 포함)
//
// [알고리즘 힌트]
//
//	시행 나눗셈: 2부터 √N까지 나누어 떨어지면 반복하여 나눈다.
//	남은 수가 1보다 크면 그 자체가 소인수이다.
func factorize(n int) []int {
	var factors []int
	for d := 2; d*d <= n; d++ {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	factors := factorize(n)
	for _, f := range factors {
		fmt.Fprintln(writer, f)
	}
}
