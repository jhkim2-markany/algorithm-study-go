package main

import (
	"bufio"
	"fmt"
	"os"
)

// evenTree는 각 연결 요소의 노드 수가 짝수가 되도록 제거할 수 있는 간선의 최대 수를 반환한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (각 원소는 [2]int{u, v}, v는 u의 부모)
//
// [반환값]
//   - int: 제거 가능한 간선의 최대 수
//
// [알고리즘 힌트]
//
//	DFS로 각 서브트리의 크기를 계산한다.
//	서브트리 크기가 짝수인 노드(루트 제외)의 수가 답이다.
func evenTree(n int, edges [][2]int) int {
	// 인접 리스트 구성
	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 서브트리 크기 배열
	subtreeSize := make([]int, n+1)
	visited := make([]bool, n+1)
	count := 0

	// DFS로 서브트리 크기 계산
	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		subtreeSize[node] = 1
		for _, next := range adj[node] {
			if !visited[next] {
				dfs(next)
				subtreeSize[node] += subtreeSize[next]
			}
		}
		// 루트가 아니고 서브트리 크기가 짝수이면 간선 제거 가능
		if node != 1 && subtreeSize[node]%2 == 0 {
			count++
		}
	}

	dfs(1)
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	result := evenTree(n, edges)
	fmt.Fprintln(writer, result)
}
