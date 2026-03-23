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
func factorial(n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

// fibonacci는 메모이제이션 재귀를 이용하여 n번째 피보나치 수를 계산한다.
//
// [매개변수]
//   - n: 피보나치 수열의 인덱스 (0-indexed)
//   - memo: 이미 계산된 결과를 저장하는 맵
//
// [반환값]
//   - int: n번째 피보나치 수
func fibonacci(n int, memo map[int]int) int {
	// 여기에 코드를 작성하세요
	return 0
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
