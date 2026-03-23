package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// generateSubsets는 1부터 n까지의 원소로 만들 수 있는 모든 부분집합을 생성한다.
//
// [매개변수]
//   - n: 집합의 크기 ({1, 2, ..., n})
//
// [반환값]
//   - [][]int: 모든 부분집합의 배열 (크기순, 사전순 정렬)
//
// [알고리즘 힌트]
//   1. 재귀 함수에서 start 인덱스부터 n까지 원소를 하나씩 선택한다.
//   2. 현재 상태(선택된 원소들)를 결과에 추가한다.
//   3. 선택 후 재귀 호출하고, 백트래킹으로 선택을 취소한다.
//   4. 결과를 크기순, 사전순으로 정렬한다.
func generateSubsets(n int) [][]int {
	var results [][]int

	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		subset := make([]int, len(current))
		copy(subset, current)
		results = append(results, subset)

		for i := start; i <= n; i++ {
			current = append(current, i)
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrack(1, []int{})

	sort.Slice(results, func(i, j int) bool {
		if len(results[i]) != len(results[j]) {
			return len(results[i]) < len(results[j])
		}
		for k := 0; k < len(results[i]); k++ {
			if results[i][k] != results[j][k] {
				return results[i][k] < results[j][k]
			}
		}
		return false
	})

	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	results := generateSubsets(n)

	fmt.Fprintln(writer, len(results))
	for _, subset := range results {
		if len(subset) == 0 {
			fmt.Fprintln(writer)
			continue
		}
		for j, v := range subset {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
