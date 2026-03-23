package main

import (
	"bufio"
	"fmt"
	"os"
)

// maxProfit은 프로젝트와 직원의 의존 관계에서 최대 가중 폐합(Maximum Weight Closure)
// 모델을 이용하여 최대 순이익을 구한다.
//
// [매개변수]
//   - n: 프로젝트 수
//   - m: 직원 수
//   - profit: 각 프로젝트의 이익
//   - salary: 각 직원의 급여
//   - requires: 각 프로젝트가 요구하는 직원 번호 목록 (1-indexed)
//
// [반환값]
//   - int: 최대 순이익 (음수이면 0)
func maxProfit(n, m int, profit, salary []int, requires [][]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	profit := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &profit[i])
	}

	salary := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &salary[i])
	}

	requires := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(reader, &k)
		requires[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &requires[i][j])
		}
	}

	fmt.Fprintln(writer, maxProfit(n, m, profit, salary, requires))
}
