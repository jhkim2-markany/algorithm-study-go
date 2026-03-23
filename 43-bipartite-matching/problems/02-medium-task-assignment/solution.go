package main

import (
	"bufio"
	"fmt"
	"os"
)

// 업무 배정 최적화 - 이분 매칭으로 최대 프로젝트 배정과 배정 결과를 구한다

var (
	adj     [][]int // 직원별 수행 가능한 프로젝트 목록
	matchR  []int   // 프로젝트별 배정된 직원 (-1이면 미배정)
	visited []bool  // DFS 방문 여부
)

// dfs: 직원 u에서 증가 경로를 찾는다
func dfs(u int) bool {
	for _, v := range adj[u] {
		if visited[v] {
			continue
		}
		visited[v] = true

		// 프로젝트 v가 미배정이거나, 현재 담당자가 다른 프로젝트로 이동 가능하면 매칭
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

	// 직원 수 N, 프로젝트 수 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 각 직원의 수행 가능한 프로젝트 목록 입력
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

	// 프로젝트 매칭 배열 초기화
	matchR = make([]int, m)
	for i := range matchR {
		matchR[i] = -1
	}

	// 각 직원에 대해 증가 경로를 찾아 매칭을 확장한다
	result := 0
	for i := 0; i < n; i++ {
		visited = make([]bool, m)
		if dfs(i) {
			result++
		}
	}

	// 최대 매칭 수 출력
	fmt.Fprintln(writer, result)

	// 각 프로젝트의 배정 결과 출력 (1-indexed 직원 번호, 미배정이면 0)
	for i := 0; i < m; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		if matchR[i] == -1 {
			fmt.Fprint(writer, 0)
		} else {
			fmt.Fprint(writer, matchR[i]+1)
		}
	}
	fmt.Fprintln(writer)
}
