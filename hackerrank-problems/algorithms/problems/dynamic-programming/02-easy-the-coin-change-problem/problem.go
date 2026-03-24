package main

import (
	"bufio"
	"fmt"
	"os"
)

// coinChange는 주어진 동전으로 금액 n을 만드는 방법의 수를 반환한다.
//
// [매개변수]
//   - n: 목표 금액
//   - coins: 동전 액면가 배열
//
// [반환값]
//   - int64: 금액 n을 만드는 방법의 수
func coinChange(n int, coins []int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 금액과 동전 종류 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 동전 액면가 입력
	coins := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &coins[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := coinChange(n, coins)
	fmt.Fprintln(writer, result)
}
