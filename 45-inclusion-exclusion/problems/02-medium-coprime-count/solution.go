package main

import (
	"bufio"
	"fmt"
	"os"
)

// 오일러 피 함수: 포함 배제의 원리를 이용하여 계산한다
// φ(n) = n × Π(1 - 1/p) (p는 n의 소인수)
func eulerPhi(n int64) int64 {
	result := n
	temp := n

	// n의 소인수를 구하며 포함 배제를 적용한다
	for p := int64(2); p*p <= temp; p++ {
		if temp%p == 0 {
			// p는 n의 소인수이다
			// result = result × (1 - 1/p) = result - result/p
			result -= result / p
			// temp에서 p를 모두 나눈다
			for temp%p == 0 {
				temp /= p
			}
		}
	}

	// temp가 1보다 크면 남은 소인수가 하나 있다
	if temp > 1 {
		result -= result / temp
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		// N 입력
		var n int64
		fmt.Fscan(reader, &n)

		// 오일러 피 함수로 서로소인 수의 개수를 구한다
		fmt.Fprintln(writer, eulerPhi(n))
	}
}
