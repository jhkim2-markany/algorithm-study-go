package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// getMinimumCost는 K명의 친구가 모든 꽃을 구매할 때 최소 비용을 반환한다.
//
// [매개변수]
//   - k: 친구의 수
//   - c: 각 꽃의 원래 가격 배열
//
// [반환값]
//   - int: 최소 총 비용
//
// [알고리즘 힌트]
//
//	비싼 꽃부터 K명에게 라운드 로빈으로 배분한다.
//	비싼 꽃에 낮은 배수를 적용하여 총 비용을 최소화한다.
func getMinimumCost(k int, c []int) int {
	// 가격을 내림차순으로 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(c)))

	total := 0
	for i, price := range c {
		// i번째 꽃을 사는 사람은 이전에 i/k개의 꽃을 산 상태
		multiplier := i/k + 1
		total += multiplier * price
	}

	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N, K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 꽃 가격 배열 입력
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := getMinimumCost(k, c)
	fmt.Fprintln(writer, result)
}
