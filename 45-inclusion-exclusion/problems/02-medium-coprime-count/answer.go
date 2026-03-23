package main

import (
	"bufio"
	"fmt"
	"os"
)

// eulerPhi는 오일러 피 함수를 계산한다.
// n 이하의 양의 정수 중 n과 서로소인 수의 개수를 반환한다.
//
// [매개변수]
//   - n: 양의 정수
//
// [반환값]
//   - int64: φ(n) 값 (n과 서로소인 수의 개수)
//
// [알고리즘 힌트]
//   1. result = n으로 초기화한다.
//   2. 2부터 √n까지 순회하며 n의 소인수 p를 찾는다.
//   3. 소인수 p를 찾으면 result = result - result/p를 적용한다 (포함 배제).
//   4. n에서 p를 모두 나눈다.
//   5. 순회 후 남은 값이 1보다 크면 마지막 소인수에 대해 같은 처리를 한다.
func eulerPhi(n int64) int64 {
	result := n
	temp := n

	for p := int64(2); p*p <= temp; p++ {
		if temp%p == 0 {
			result -= result / p
			for temp%p == 0 {
				temp /= p
			}
		}
	}

	if temp > 1 {
		result -= result / temp
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscan(reader, &n)
		fmt.Fprintln(writer, eulerPhi(n))
	}
}
