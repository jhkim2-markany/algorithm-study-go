package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

// exactlyMSets는 포함 배제의 원리를 이용하여 정확히 m개의 집합에 속하는 원소의 수를 구한다.
//
// [매개변수]
//   - n: 전체 원소 범위 (1부터 n)
//   - primes: 소수 배열 (각 소수 p에 대해 집합 A_p = {p의 배수})
//   - m: 정확히 속해야 하는 집합의 수
//
// [반환값]
//   - int: 정확히 m개의 집합에 속하는 원소의 수
func exactlyMSets(n int, primes []int, m int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k, m int
	fmt.Fscan(reader, &n, &k, &m)

	primes := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &primes[i])
	}

	fmt.Fprintln(writer, exactlyMSets(n, primes, m))

	_ = bits.OnesCount // 패키지 사용 보장
}
