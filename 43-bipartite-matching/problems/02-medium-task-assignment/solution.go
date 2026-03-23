package main

import (
	"bufio"
	"fmt"
	"os"
)

// taskAssignment은 이분 매칭으로 최대 프로젝트 배정 수와 각 프로젝트의 배정 결과를 구한다.
//
// [매개변수]
//   - n: 직원 수
//   - m: 프로젝트 수
//   - adj: 직원별 수행 가능한 프로젝트 목록 (0-indexed)
//
// [반환값]
//   - int: 최대 매칭 수
//   - []int: 각 프로젝트에 배정된 직원 번호 (1-indexed, 미배정이면 0)
func taskAssignment(n, m int, adj [][]int) (int, []int) {
	// 여기에 코드를 작성하세요
	return 0, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(reader, &k)
		adj[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &adj[i][j])
			adj[i][j]--
		}
	}

	result, assignment := taskAssignment(n, m, adj)
	fmt.Fprintln(writer, result)

	for i := 0; i < m; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, assignment[i])
	}
	fmt.Fprintln(writer)
}
