package main

import (
	"bufio"
	"fmt"
	"os"
)

// knapsack은 0/1 배낭 문제에서 담을 수 있는 최대 가치를 반환한다.
//
// [매개변수]
//   - n: 물건의 개수
//   - k: 배낭의 용량
//   - weight: 각 물건의 무게 배열 (1-indexed)
//   - value: 각 물건의 가치 배열 (1-indexed)
//
// [반환값]
//   - int: 배낭에 담을 수 있는 최대 가치
func knapsack(n, k int, weight, value []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n, k int
	fmt.Fscan(reader, &n, &k)

	weight := make([]int, n+1)
	value := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &weight[i], &value[i])
	}

	// 핵심 함수 호출
	result := knapsack(n, k, weight, value)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
