package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// jimOrders는 버거 완성 순서대로 고객 번호를 반환한다.
//
// [매개변수]
//   - orders: 각 고객의 [주문시각, 준비시간] 배열
//
// [반환값]
//   - []int: 완성 순서대로 정렬된 고객 번호 (1-indexed)
func jimOrders(orders [][]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 고객 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 주문 정보 입력
	orders := make([][]int, n)
	for i := 0; i < n; i++ {
		orders[i] = make([]int, 2)
		fmt.Fscan(reader, &orders[i][0], &orders[i][1])
	}

	// 핵심 함수 호출 및 결과 출력
	result := jimOrders(orders)
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Slice
}
