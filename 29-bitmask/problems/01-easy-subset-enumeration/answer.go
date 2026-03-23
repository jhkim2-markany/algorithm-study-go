package main

import (
	"bufio"
	"fmt"
	"os"
)

// enumerateSubsets는 N개 원소의 모든 부분집합을 비트마스크로 열거하여 반환한다.
//
// [매개변수]
//   - n: 원소의 수
//
// [반환값]
//   - [][]int: 각 부분집합에 포함된 원소 목록 (1-based)
//
// [알고리즘 힌트]
//
//	0부터 2^N - 1까지 모든 비트마스크를 순회한다.
//	각 비트마스크에서 i번째 비트가 1이면 원소 i+1이 포함된다.
//	비트 AND 연산 mask & (1 << i)로 포함 여부를 확인한다.
func enumerateSubsets(n int) [][]int {
	total := 1 << n
	result := make([][]int, total)
	for mask := 0; mask < total; mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, i+1)
			}
		}
		result[mask] = subset
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	subsets := enumerateSubsets(n)

	// 결과 출력
	for _, subset := range subsets {
		for i, v := range subset {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
