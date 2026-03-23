package main

import (
	"bufio"
	"fmt"
	"os"
)

// 최소 버텍스 커버 - 쾨니그 정리를 활용하여 최소 카메라 수를 구한다
// 이분 그래프에서 최대 매칭 = 최소 버텍스 커버 (쾨니그 정리)

var (
	adj     [][]int // 행별 연결된 열 목록 (장애물이 없는 칸)
	matchR  []int   // 열별 매칭된 행 (-1이면 미매칭)
	visited []bool  // DFS 방문 여부
)

// dfs: 행 u에서 증가 경로를 찾는다
func dfs(u int) bool {
	for _, v := range adj[u] {
		if visited[v] {
			continue
		}
		visited[v] = true

		// 열 v가 미매칭이거나, 현재 매칭된 행이 다른 열로 이동 가능하면 매칭
		if matchR[v] == -1 || dfs(matchR[v]) {
			matchR[v] = u
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 행 수 N, 열 수 M, 장애물 수 K 입력
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	// 장애물 위치를 기록한다
	blocked := make(map[[2]int]bool)
	for i := 0; i < k; i++ {
		var r, c int
		fmt.Fscan(reader, &r, &c)
		blocked[[2]int{r - 1, c - 1}] = true
	}

	// 이분 그래프 구성: 행을 왼쪽, 열을 오른쪽으로 설정
	// 장애물이 없는 칸 (r, c)에 대해 행 r과 열 c를 간선으로 연결한다
	adj = make([][]int, n)
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if !blocked[[2]int{r, c}] {
				adj[r] = append(adj[r], c)
			}
		}
	}

	// 열 매칭 배열 초기화
	matchR = make([]int, m)
	for i := range matchR {
		matchR[i] = -1
	}

	// 헝가리안 알고리즘으로 최대 매칭을 구한다
	result := 0
	for i := 0; i < n; i++ {
		visited = make([]bool, m)
		if dfs(i) {
			result++
		}
	}

	// 쾨니그 정리: 최대 매칭 = 최소 버텍스 커버 = 최소 카메라 수
	fmt.Fprintln(writer, result)
}
