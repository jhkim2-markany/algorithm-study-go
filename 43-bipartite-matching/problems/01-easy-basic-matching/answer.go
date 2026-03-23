package main

import (
	"bufio"
	"fmt"
	"os"
)

// maxMatching은 이분 그래프에서 헝가리안 알고리즘으로 최대 매칭 수를 구한다.
//
// [매개변수]
//   - n: 왼쪽 정점(학생)의 수
//   - m: 오른쪽 정점(동아리)의 수
//   - adj: 왼쪽 정점별 연결된 오른쪽 정점 목록 (0-indexed)
//
// [반환값]
//   - int: 최대 매칭 수
//
// [알고리즘 힌트]
//   1. 오른쪽 정점의 매칭 배열(matchR)을 -1로 초기화한다.
//   2. 각 왼쪽 정점에 대해 DFS로 증가 경로를 탐색한다.
//   3. DFS에서 오른쪽 정점이 미매칭이거나, 현재 매칭 상대가 다른 경로를 찾을 수 있으면 매칭한다.
//   4. 증가 경로를 찾을 때마다 매칭 수를 1 증가시킨다.
func maxMatching(n, m int, adj [][]int) int {
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
	return result
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

	fmt.Fprintln(writer, maxMatching(n, m, adj))
}
