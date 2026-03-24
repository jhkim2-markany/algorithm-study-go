package main

import (
	"bufio"
	"fmt"
	"os"
)

// minimumLoss는 매입가 - 매도가의 최솟값(양수)을 반환한다.
// 매입은 매도보다 먼저 일어나야 하며, 매입가 > 매도가이다.
//
// [매개변수]
//   - price: 연도별 집값 배열
//
// [반환값]
//   - int64: 최소 손실
func minimumLoss(price []int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	price := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &price[i])
	}

	fmt.Fprintln(writer, minimumLoss(price))
}
