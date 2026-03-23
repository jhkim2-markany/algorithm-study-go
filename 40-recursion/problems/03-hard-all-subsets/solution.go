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
func generateSubsets(n int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
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

	_ = sort.Slice // 패키지 사용 보장
}
