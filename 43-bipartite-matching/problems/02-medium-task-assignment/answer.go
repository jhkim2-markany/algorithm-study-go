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
//
// [알고리즘 힌트]
//   1. 프로젝트 매칭 배열(matchR)을 -1로 초기화한다.
//   2. 각 직원에 대해 DFS로 증가 경로를 탐색하여 매칭을 확장한다.
//   3. 매칭 완료 후 matchR 배열을 1-indexed 직원 번호로 변환하여 반환한다.
func taskAssignment(n, m int, adj [][]int) (int, []int) {
	matchR := make([]int, m)
	for i := range matchR {
		matchR[i] = -1
	}

	var visited []bool

	var dfs func(u int) bool
	dfs = func(u int) bool {
		for _, v := range adj[u] {
			if visited[v] {
				continue
			}
			visited[v] = true
			if matchR[v] == -1 || dfs(matchR[v]) {
				matchR[v] = u
				return true
			}
		}
		return false
	}

	result := 0
	for i := 0; i < n; i++ {
		visited = make([]bool, m)
		if dfs(i) {
			result++
		}
	}

	assignment := make([]int, m)
	for i := 0; i < m; i++ {
		if matchR[i] == -1 {
			assignment[i] = 0
		} else {
			assignment[i] = matchR[i] + 1
		}
	}

	return result, assignment
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
