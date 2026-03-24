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
//
// [알고리즘 힌트]
//
//	각 고객의 완성 시각(order + prep)을 계산하고, 이를 기준으로 정렬한다.
//	동일 시각이면 고객 번호가 작은 순서로 정렬한다.
func jimOrders(orders [][]int) []int {
	n := len(orders)

	// (완성시각, 고객번호) 쌍을 생성
	type customer struct {
		finishTime int
		index      int
	}
	customers := make([]customer, n)
	for i := 0; i < n; i++ {
		customers[i] = customer{
			finishTime: orders[i][0] + orders[i][1],
			index:      i + 1, // 1-indexed
		}
	}

	// 완성 시각 기준 오름차순 정렬, 동일하면 고객 번호 오름차순
	sort.Slice(customers, func(a, b int) bool {
		if customers[a].finishTime == customers[b].finishTime {
			return customers[a].index < customers[b].index
		}
		return customers[a].finishTime < customers[b].finishTime
	})

	// 정렬된 순서대로 고객 번호 추출
	result := make([]int, n)
	for i, c := range customers {
		result[i] = c.index
	}

	return result
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
}
