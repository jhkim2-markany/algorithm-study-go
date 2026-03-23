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

	// N 입력
	var n int
	fmt.Fscan(reader, &n)

	// 에라토스테네스의 체: 소수 여부 배열 초기화
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// 2부터 √N까지 순회하며 배수를 제거한다
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// i의 배수를 합성수로 표시 (i*i부터 시작)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 소수 목록 수집
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	// 소수 개수 출력
	fmt.Fprintln(writer, len(primes))

	// 소수 목록 출력
	for i, p := range primes {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, p)
	}
	fmt.Fprintln(writer)
}
