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

	// 에라토스테네스의 체로 소수를 구한다
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// √N까지 순회하며 배수를 제거한다
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// i*i부터 i의 배수를 모두 제거한다
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 소수를 오름차순으로 출력한다
	first := true
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			if !first {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, i)
			first = false
		}
	}
	fmt.Fprintln(writer)
}
