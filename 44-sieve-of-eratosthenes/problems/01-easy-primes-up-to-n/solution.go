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
func sieveOfEratosthenes(n int) []int {
	// 여기에 코드를 작성하세요
	return nil
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
