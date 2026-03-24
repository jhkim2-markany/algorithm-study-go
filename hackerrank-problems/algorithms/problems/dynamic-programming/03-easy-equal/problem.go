package main

import (
	"bufio"
	"fmt"
	"os"
)

// equal은 모든 동료가 같은 수의 초콜릿을 갖도록 하는 최소 연산 횟수를 반환한다.
//
// [매개변수]
//   - arr: 각 동료의 초콜릿 수 배열
//
// [반환값]
//   - int: 최소 연산 횟수
func equal(arr []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		// 동료 수 입력
		var n int
		fmt.Fscan(reader, &n)

		// 초콜릿 수 배열 입력
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := equal(arr)
		fmt.Fprintln(writer, result)
	}
}
