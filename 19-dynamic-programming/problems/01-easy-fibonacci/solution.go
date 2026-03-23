package main

import (
	"bufio"
	"fmt"
	"os"
)

// fibonacci는 N번째 피보나치 수를 반환한다.
//
// [매개변수]
//   - n: 구하고자 하는 피보나치 수의 인덱스 (0-indexed)
//
// [반환값]
//   - int: N번째 피보나치 수
func fibonacci(n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	result := fibonacci(n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
