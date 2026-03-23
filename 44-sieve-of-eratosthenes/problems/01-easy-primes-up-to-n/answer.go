package main

import (
	"bufio"
	"fmt"
	"os"
)

// sieveOfEratosthenes는 에라토스테네스의 체로 n 이하의 모든 소수를 구한다.
//
// [매개변수]
//   - n: 소수를 구할 상한값
//
// [반환값]
//   - []int: n 이하의 소수 배열 (오름차순)
//
// [알고리즘 힌트]
//   1. 크기 n+1의 불리언 배열을 만들어 2부터 n까지 true로 초기화한다.
//   2. 2부터 √n까지 순회하며, 소수인 i에 대해 i*i부터 i의 배수를 모두 false로 표시한다.
//   3. true로 남은 인덱스를 수집하여 소수 배열로 반환한다.
func sieveOfEratosthenes(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	primes := sieveOfEratosthenes(n)

	for i, p := range primes {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, p)
	}
	fmt.Fprintln(writer)
}
