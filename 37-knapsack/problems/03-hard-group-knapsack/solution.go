package main

import (
	"bufio"
	"fmt"
	"os"
)

// Group은 그룹 내 물건들의 무게와 가치를 저장한다.
type Group struct {
	weights []int
	values  []int
}

// groupKnapsack은 그룹 배낭 문제의 최대 가치를 반환한다.
//
// [매개변수]
//   - n: 그룹의 수
//   - k: 배낭의 용량
//   - groups: 각 그룹의 물건 정보 배열 (길이 n)
//
// [반환값]
//   - int: 각 그룹에서 최대 1개씩 선택하여 담을 수 있는 최대 가치
func groupKnapsack(n, k int, groups []Group) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	groups := make([]Group, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(reader, &m)
		groups[i].weights = make([]int, m)
		groups[i].values = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &groups[i].weights[j], &groups[i].values[j])
		}
	}

	fmt.Fprintln(writer, groupKnapsack(n, k, groups))
}
