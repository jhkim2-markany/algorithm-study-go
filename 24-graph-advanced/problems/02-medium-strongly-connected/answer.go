package main

import (
	"bufio"
	"fmt"
	"os"
)

// countSCC는 코사라주 알고리즘으로 강한 연결 요소의 개수를 구한다.
//
// [매개변수]
//   - adj: 인접 리스트로 표현된 방향 그래프 (1-indexed)
//   - radj: 역방향 인접 리스트 (1-indexed)
//   - n: 정점의 수
//
// [반환값]
//   - int: 강한 연결 요소(SCC)의 개수
//
// [알고리즘 힌트]
//
//	코사라주 알고리즘을 사용한다.
//	1단계: 원본 그래프에서 DFS를 수행하며 완료 순서를 기록한다.
//	2단계: 역방향 그래프에서 완료 순서의 역순으로 DFS를 수행한다.
//	각 DFS 호출이 하나의 SCC를 구성한다.
func countSCC(adj, radj [][]int, n int) int {
	visited := make([]bool, n+1)
	order := []int{}

	// 1단계: 원본 그래프에서 DFS, 완료 순서 기록
	var dfs1 func(u int)
	dfs1 = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs1(v)
			}
		}
		order = append(order, u)
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs1(i)
		}
	}

	// 2단계: 역방향 그래프에서 완료 순서의 역순으로 DFS
	for i := range visited {
		visited[i] = false
	}

	var dfs2 func(u int)
	dfs2 = func(u int) {
		visited[u] = true
		for _, v := range radj[u] {
			if !visited[v] {
				dfs2(v)
			}
		}
	}

	sccCount := 0
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		if !visited[u] {
			dfs2(u)
			sccCount++
		}
	}

	return sccCount
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화 (원본 + 역방향)
	adj := make([][]int, n+1)
	radj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []int{}
		radj[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		radj[v] = append(radj[v], u)
	}

	// 핵심 함수 호출
	result := countSCC(adj, radj, n)

	fmt.Fprintln(writer, result)
}
