package main

import (
	"bufio"
	"fmt"
	"os"
)

// 이분 매칭 - 헝가리안 알고리즘으로 최대 매칭을 구한다

var (
	adj     [][]int // 왼쪽 정점의 인접 리스트
	matchR  []int   // 오른쪽 정점의 매칭 상대 (-1이면 미매칭)
	visited []bool  // DFS 방문 여부
)

// dfs: 왼쪽 정점 u에서 증가 경로를 찾는다
func dfs(u int) bool {
	for _, v := range adj[u] {
		if visited[v] {
			continue
		}
		visited[v] = true

		// v가 미매칭이거나 v의 매칭 상대에서 다른 경로를 찾을 수 있으면 매칭
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

	// 학생 수 N, 동아리 수 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 각 학생의 가입 가능한 동아리 목록 입력
	adj = make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(reader, &k)
		adj[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &adj[i][j])
			adj[i][j]-- // 0-indexed로 변환
		}
	}

	// 오른쪽 정점(동아리) 매칭 배열 초기화
	matchR = make([]int, m)
	for i := range matchR {
		matchR[i] = -1
	}

	// 각 학생에 대해 증가 경로를 찾아 매칭 수를 센다
	result := 0
	for i := 0; i < n; i++ {
		visited = make([]bool, m)
		if dfs(i) {
			result++
		}
	}

	// 최대 매칭 수 출력
	fmt.Fprintln(writer, result)
}
