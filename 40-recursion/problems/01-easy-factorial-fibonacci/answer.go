package main

import (
	"bufio"
	"fmt"
	"os"
)

// factorial은 재귀를 이용하여 n!을 계산한다.
//
// [매개변수]
//   - n: 팩토리얼을 구할 음이 아닌 정수
//
// [반환값]
//   - int: n!의 값
//
// [알고리즘 힌트]
//   1. 기저 조건: n ≤ 1이면 1을 반환한다.
//   2. 재귀: n * factorial(n-1)을 반환한다.
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// fibonacci는 메모이제이션 재귀를 이용하여 n번째 피보나치 수를 계산한다.
//
// [매개변수]
//   - n: 피보나치 수열의 인덱스 (0-indexed)
//   - memo: 이미 계산된 결과를 저장하는 맵
//
// [반환값]
//   - int: n번째 피보나치 수
//
// [알고리즘 힌트]
//   1. 캐시에 이미 계산된 값이 있으면 바로 반환한다.
//   2. 기저 조건: n ≤ 0이면 0, n = 1이면 1을 반환한다.
//   3. fibonacci(n-1) + fibonacci(n-2)를 계산하고 memo에 저장한다.
func fibonacci(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	result := fibonacci(n-1, memo) + fibonacci(n-2, memo)
	memo[n] = result
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)

	fmt.Fprintln(writer, factorial(n))

	memo := make(map[int]int)
	fmt.Fprintln(writer, fibonacci(m, memo))
}
