package main

import (
	"bufio"
	"fmt"
	"os"
)

// sieveOfEratosthenes는 2 이상 n 이하의 모든 소수를 반환한다.
//
// [매개변수]
//   - n: 소수를 구할 상한값
//
// [반환값]
//   - []int: 2 이상 n 이하의 소수 배열 (오름차순)
//
// [알고리즘 힌트]
//
//	에라토스테네스의 체: 크기 n+1의 bool 배열을 만들고 모두 true로 초기화한다.
//	2부터 √n까지 순회하며, 소수인 i의 배수를 모두 false로 표시한다.
//	i*i부터 시작하면 중복 제거를 줄일 수 있다.
//	남은 true인 인덱스가 소수이다.
//
//	시간복잡도: O(N log log N), 공간복잡도: O(N)
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

	// N 입력
	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	primes := sieveOfEratosthenes(n)

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
