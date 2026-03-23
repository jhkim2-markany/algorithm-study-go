package main

import (
	"bufio"
	"fmt"
	"os"
)

// minVertexCover는 이분 그래프에서 쾨니그 정리를 활용하여 최소 버텍스 커버 크기를 구한다.
// 격자에서 장애물이 없는 칸을 행-열 이분 그래프로 모델링한다.
//
// [매개변수]
//   - n: 행의 수
//   - m: 열의 수
//   - blocked: 장애물 위치 맵 (0-indexed, [행][열] → true)
//
// [반환값]
//   - int: 최소 버텍스 커버 크기 (= 최대 매칭 수)
//
// [알고리즘 힌트]
//   1. 행을 왼쪽, 열을 오른쪽으로 하는 이분 그래프를 구성한다.
//   2. 장애물이 없는 칸 (r, c)에 대해 행 r과 열 c를 간선으로 연결한다.
//   3. 헝가리안 알고리즘으로 최대 매칭을 구한다.
//   4. 쾨니그 정리에 의해 최대 매칭 = 최소 버텍스 커버이다.
func minVertexCover(n, m int, blocked map[[2]int]bool) int {
	// 이분 그래프 구성
	adj := make([][]int, n)
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if !blocked[[2]int{r, c}] {
				adj[r] = append(adj[r], c)
			}
		}
	}

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

	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	blocked := make(map[[2]int]bool)
	for i := 0; i < k; i++ {
		var r, c int
		fmt.Fscan(reader, &r, &c)
		blocked[[2]int{r - 1, c - 1}] = true
	}

	fmt.Fprintln(writer, minVertexCover(n, m, blocked))
}
