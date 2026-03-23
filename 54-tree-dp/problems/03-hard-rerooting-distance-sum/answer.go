package main

import (
	"bufio"
	"fmt"
	"os"
)

// rerootingDistanceSum은 리루팅 기법으로 모든 노드의 전체 거리 합을 계산한다.
//
// [매개변수]
//   - n: 노드의 수
//   - adj: 인접 리스트
//
// [반환값]
//   - []int: 각 노드를 루트로 했을 때의 전체 거리 합 배열 (1-indexed)
//
// [알고리즘 힌트]
//
//	리루팅(Rerooting) 기법을 사용한다.
//	1단계: 루트(1번)에서 DFS로 서브트리 크기 sz[v]와 서브트리 내 거리 합 dp[v]를 계산한다.
//	2단계: 루트에서 다시 DFS하며, 부모의 전체 거리 합으로부터 자식의 전체 거리 합을 유도한다.
//	핵심 공식: ans[child] = ans[parent] - sz[child] + (N - sz[child])
func rerootingDistanceSum(n int, adj [][]int) []int {
	sz := make([]int, n+1)
	dp := make([]int, n+1)
	ans := make([]int, n+1)

	// 1단계: Bottom-Up DFS (비재귀)
	parent := make([]int, n+1)
	order := make([]int, 0, n)
	visited := make([]bool, n+1)

	parent[1] = 0
	visited[1] = true
	queue := []int{1}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, u := range adj[v] {
			if !visited[u] {
				visited[u] = true
				parent[u] = v
				queue = append(queue, u)
			}
		}
	}

	// 역순으로 서브트리 크기와 거리 합 계산
	for i := 0; i <= n; i++ {
		sz[i] = 1
	}
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		for _, u := range adj[v] {
			if u != parent[v] {
				sz[v] += sz[u]
				dp[v] += dp[u] + sz[u]
			}
		}
	}

	// 2단계: Top-Down 리루팅
	ans[1] = dp[1]
	for i := 0; i < len(order); i++ {
		v := order[i]
		for _, u := range adj[v] {
			if u != parent[v] {
				ans[u] = ans[v] - sz[u] + (n - sz[u])
			}
		}
	}

	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	adj := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := rerootingDistanceSum(n, adj)

	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, ans[i])
	}
	fmt.Fprintln(writer)
}
