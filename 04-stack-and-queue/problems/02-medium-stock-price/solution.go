package main

import (
	"bufio"
	"fmt"
	"os"
)

// stockPrice는 각 날짜의 주식 가격이 떨어지지 않은 기간(일수)을 반환한다.
//
// [매개변수]
//   - prices: N일 동안의 주식 가격 배열
//
// [반환값]
//   - []int: 각 날짜별 가격이 떨어지지 않은 기간(일수) 배열
func stockPrice(prices []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 날짜 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 주식 가격 입력
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &prices[i])
	}

	// 핵심 함수 호출
	answer := stockPrice(prices)

	// 결과 출력
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, answer[i])
	}
	fmt.Fprintln(writer)
}
