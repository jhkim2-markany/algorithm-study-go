package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var adj [][]int
var discovered []int // 각 정점의 DFS 방문 순서
var isAP []bool      // 단절점 여부
var timer int

// dfs 함수는 Tarjan 알고리즘으로 단절점을 찾는다
// 반환값: cur의 서브트리에서 도달 가능한 가장 빠른 방문 순서
func dfs(cur, parent int) int {
	// 현재 정점에 방문 순서를 부여
	timer++
	discovered[cur] = timer
	minOrder := discovered[cur]
	childCount := 0

	for _, next := range adj[cur] {
		if discovered[next] == 0 {
			// 미방문 정점 → 트리 간선
			childCount++
			low := dfs(next, cur)

			// 자식 서브트리에서 도달 가능한 최소 방문 순서 갱신
			if low < minOrder {
				minOrder = low
			}

			// 루트가 아닌 정점: 자식의 low 값이 자신의 방문 순서 이상이면 단절점
			if parent != 0 && low >= discovered[cur] {
				isAP[cur] = true
			}
		} else if next != parent {
			// 이미 방문한 정점 → 역방향 간선 (부모 제외)
			if discovered[next] < minOrder {
				minOrder = discovered[next]
			}
		}
	}

	// 루트 정점: DFS 트리에서 자식이 2개 이상이면 단절점
	if parent == 0 && childCount >= 2 {
		isAP[cur] = true
	}

	return minOrder
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정점 수와 간선 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화
	adj = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력 (양방향)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Tarjan 알고리즘으로 단절점 탐색
	discovered = make([]int, n+1)
	isAP = make([]bool, n+1)
	timer = 0
	dfs(1, 0)

	// 단절점 수집 및 정렬
	result := []int{}
	for i := 1; i <= n; i++ {
		if isAP[i] {
			result = append(result, i)
		}
	}
	sort.Ints(result)

	// 결과 출력
	fmt.Fprintln(writer, len(result))
	if len(result) > 0 {
		for i, v := range result {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
