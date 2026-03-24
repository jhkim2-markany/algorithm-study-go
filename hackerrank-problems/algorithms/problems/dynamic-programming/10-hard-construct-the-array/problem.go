package main

import (
	"bufio"
	"fmt"
	"os"
)

// constructArray는 조건을 만족하는 배열의 개수를 반환한다.
//
// [매개변수]
//   - n: 배열의 길이
//   - k: 원소의 최댓값
//   - x: 마지막 원소의 값
//
// [반환값]
//   - int: 조건을 만족하는 배열의 수 (mod 10^9+7)
func constructArray(n, k, x int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: n, k, x
	var n, k, x int
	fmt.Fscan(reader, &n, &k, &x)

	// 핵심 함수 호출 및 결과 출력
	result := constructArray(n, k, x)
	fmt.Fprintln(writer, result)
}
