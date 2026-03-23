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
func sieveOfEratosthenes(n int) []int {
	// 여기에 코드를 작성하세요
	return nil
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
