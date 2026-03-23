package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	writer  *bufio.Writer
	n       int
	results [][]int
)

// 재귀적으로 부분집합을 생성하는 함수
// start: 현재 선택할 수 있는 최소 원소
// current: 현재까지 선택한 원소들
func generateSubsets(start int, current []int) {
	// 현재 상태를 결과에 추가 (복사본 저장)
	subset := make([]int, len(current))
	copy(subset, current)
	results = append(results, subset)

	// start부터 n까지 원소를 하나씩 선택하여 재귀 호출
	for i := start; i <= n; i++ {
		// 원소 i를 선택
		current = append(current, i)
		// 다음 원소부터 재귀 탐색 (중복 방지를 위해 i+1부터)
		generateSubsets(i+1, current)
		// 백트래킹: 선택 취소
		current = current[:len(current)-1]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 집합의 크기
	fmt.Fscan(reader, &n)

	// 모든 부분집합 생성
	results = [][]int{}
	generateSubsets(1, []int{})

	// 크기순, 사전순 정렬
	sort.Slice(results, func(i, j int) bool {
		if len(results[i]) != len(results[j]) {
			return len(results[i]) < len(results[j])
		}
		// 크기가 같으면 원소를 사전순 비교
		for k := 0; k < len(results[i]); k++ {
			if results[i][k] != results[j][k] {
				return results[i][k] < results[j][k]
			}
		}
		return false
	})

	// 총 개수 출력
	fmt.Fprintln(writer, len(results))

	// 각 부분집합 출력
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
