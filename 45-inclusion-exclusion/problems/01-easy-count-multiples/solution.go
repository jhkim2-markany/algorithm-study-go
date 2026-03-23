package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

// countMultiples는 포함 배제의 원리를 이용하여 1부터 n까지 주어진 소수 중
// 적어도 하나의 배수인 수의 개수를 구한다.
//
// [매개변수]
//   - n: 범위 상한값
//   - primes: 소수 배열
//
// [반환값]
//   - int: 1부터 n까지 primes 중 적어도 하나의 배수인 수의 개수
func countMultiples(n int, primes []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	primes := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &primes[i])
	}

	fmt.Fprintln(writer, countMultiples(n, primes))

	_ = bits.OnesCount // 패키지 사용 보장
}
