package main

import (
	"bufio"
	"fmt"
	"os"
)

// knapsack01은 0/1 배낭 문제의 최대 가치를 반환한다.
//
// [매개변수]
//   - n: 물건의 수
//   - k: 배낭의 용량
//   - weights: 각 물건의 무게 배열 (길이 n)
//   - values: 각 물건의 가치 배열 (길이 n)
//
// [반환값]
//   - int: 용량 k 이하로 담을 수 있는 최대 가치
func knapsack01(n, k int, weights, values []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	weights := make([]int, n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &weights[i], &values[i])
	}

	fmt.Fprintln(writer, knapsack01(n, k, weights, values))
}
